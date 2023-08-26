package auth

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/proto"
	"sync"
)

var (
	cli  *clientv3.Client
	once = &sync.Once{}
)

func AuthenticateClient(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		hlog.Infof("AuthMiddleware get req: %+v", request)
		key, err := getAESKey()
		if err != nil {
			hlog.Errorf("AuthMiddleware get aes failed: %+v", err)
			return err
		}

		// 序列化
		marshal, err := proto.Marshal(request.(proto.Message))
		if err != nil {
			hlog.Errorf("AuthMiddleware marshal failed: %+v", err)
			return err
		}

		// 加密请求
		encryptedReq, err := encrypt(marshal, key)
		if err != nil {
			hlog.Errorf("AuthMiddleware encrypt failed: %+v", err)
			return err
		}

		hlog.Infof("AuthMiddleware encrypted req: %+v", encryptedReq)
		// 调用下一个中间件
		err = next(ctx, encryptedReq, response)
		if err != nil {
			return err
		}
		hlog.Infof("AuthMiddleware get encrypted resp: %+v", response)
		message := response.(proto.Message)
		// 序列化
		marshal, err = proto.Marshal(message)
		if err != nil {
			hlog.Errorf("AuthMiddleware marshal failed: %+v", err)
			return err
		}
		return err
	}
}

func AuthenticateServer(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		hlog.Infof("AuthMiddleware get req: %+v", request)
		// 创建etcd客户端连接
		key, err := getAESKey()
		if err != nil {
			hlog.Errorf("AuthMiddleware get aes failed: %+v", err)
			return err
		}

		// 解密请求
		decryptedReq, err := decrypt(nil, key)
		hlog.Infof("AuthMiddleware decrypted req: %+v", decryptedReq)
		// 调用下一个中间件
		err = next(ctx, decryptedReq, response)
		return err
	}
}

func encrypt(data []byte, key string) ([]byte, error) {
	// 创建 AES 加密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		hlog.Errorf("AuthMiddleware encrypt failed: %+v", err)
		return nil, err
	}

	// 创建 GCM 加密模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		hlog.Errorf("AuthMiddleware encrypt failed: %+v", err)
		return nil, err
	}

	// 生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	// 加密
	ciphertext := gcm.Seal(nil, nonce, data, nil)
	encryptedData := append(nonce, ciphertext...)

	return encryptedData, nil
}

func decrypt(data []byte, key string) ([]byte, error) {
	// 创建 AES 解密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// 使用 AES 解密器创建 GCM 块模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 提取初始化向量（IV）
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// 解密密文
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func getEtcdClient(once *sync.Once) (err error) {
	if cli == nil {
		once.Do(func() {
			cli, err = clientv3.New(clientv3.Config{
				Endpoints: []string{"http://106.54.208.133:2379"}, // etcd服务地址
			})
			if err != nil {
				hlog.Fatalf("AuthMiddleware etcd client init failed: %+v", err)
			}
		})
		return nil
	}
	return nil
}

func getAESKey() (string, error) {
	err := getEtcdClient(once)
	if err != nil {
		hlog.Errorf("AuthMiddleware get etcd client failed: %+v", err)
		return "", err
	}

	// 使用客户端从etcd中获取密钥
	resp, err := cli.Get(context.Background(), "AESKEY")
	if err != nil {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return "", err
	}

	// 检查是否找到了密钥
	if len(resp.Kvs) == 0 {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return "", errors.New("AuthMiddleware get key failed")
	}

	// 获取密钥
	key := string(resp.Kvs[0].Value)
	hlog.Infof("AuthMiddleware get key: %s", key)
	return key, nil
}

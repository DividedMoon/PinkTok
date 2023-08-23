package auth

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.etcd.io/etcd/clientv3"
	"sync"
)

var (
	cli  *clientv3.Client
	once = &sync.Once{}
)

func AuthenticateMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		hlog.Infof("AuthMiddleware get req: %+v", request)
		// 创建etcd客户端连接
		err := getEtcdClient(once)
		if err != nil {
			hlog.Errorf("AuthMiddleware get etcd client failed: %+v", err)
			return err
		}

		// 使用客户端从etcd中获取密钥
		resp, err := cli.Get(context.Background(), "AESKEY")
		if err != nil {
			hlog.Errorf("AuthMiddleware get key failed: %+v", err)
			return err
		}

		// 检查是否找到了密钥
		if len(resp.Kvs) == 0 {
			hlog.Errorf("AuthMiddleware get key failed: %+v", err)
			return errors.New("AuthMiddleware get key failed")
		}

		// 获取密钥
		key := string(resp.Kvs[0].Value)
		hlog.Infof("AuthMiddleware get key: %s", key)

		// 加密请求
		encryptedReq, err := encryptReq(request, key)
		if err != nil {
			hlog.Errorf("AuthMiddleware encrypt failed: %+v", err)
			return err
		}

		hlog.Infof("AuthMiddleware encrypted req: %+v", encryptedReq)
		// 调用下一个中间件
		err = next(ctx, encryptedReq, response)
		return err
	}
}

func encryptReq(req interface{}, key string) (interface{}, error) {
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
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	// 加密数据
	if str, ok := req.(string); ok {
		ciphertext := gcm.Seal(nil, nonce, []byte(str), nil)
		// 将加密后的密文进行 Base64 编码
		encryptedData := append(nonce, ciphertext...)
		encrypted := base64.StdEncoding.EncodeToString(encryptedData)
		return encrypted, nil
	}

	return nil, errors.New("AuthMiddleware req is not string")
}

func decryptResp(resp interface{}, key string) (interface{}, error) {
	// 创建 AES 解密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 使用 AES 解密器创建 GCM 块模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 对 Base64 编码后的密文进行解码
	if _, ok := resp.(string); !ok {
		return "", errors.New("AuthMiddleware resp is not string")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(resp.(string))
	if err != nil {
		return "", err
	}

	// 提取初始化向量（IV）
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 解密密文
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func getEtcdClient(once *sync.Once) (err error) {
	once.Do(func() {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints: []string{"http://106.54.208.133:2379"}, // etcd服务地址
		})
		if err != nil {
			hlog.Errorf("AuthMiddleware etcd client init failed: %+v", err)
		}
	})
	return err
}

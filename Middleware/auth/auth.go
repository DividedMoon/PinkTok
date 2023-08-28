package auth

import (
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.etcd.io/etcd/client/v3"
	"strings"
	"sync"
)

const (
	PrivateKeyKey  = "PrivateKey"
	PublicKeyKey   = "PublicKey"
	SignReqHeader  = "signature-req"
	SignRespHeader = "signature-resp"
)

var (
	cli  *clientv3.Client
	once = &sync.Once{}
)

func AuthenticateClient(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		err := getEtcdClient(once)
		hlog.Infof("AuthMiddleware get req: %+v", request)
		key, err := getRsaPrivateKey(cli)
		if err != nil {
			hlog.Errorf("AuthMiddleware get privateKey failed: %+v", err)
			return err
		}
		requestStr := fmt.Sprintf("%s", request)
		requestStr = strings.ReplaceAll(requestStr, "\\", "")
		signature, err := sign(key, requestStr)
		if err != nil {
			hlog.Errorf("AuthMiddleware sign failed: %+v", err)
			return err
		}
		ctx = metainfo.WithValue(ctx, SignReqHeader, string(signature))
		// 调用下一个中间件
		err = next(ctx, request, response)
		if err != nil {
			return err
		}

		hlog.Infof("AuthMiddleware get resp: %+v", response)
		// 响应验签
		value, ok := metainfo.RecvBackwardValue(ctx, SignRespHeader)
		if !ok {
			return errors.New("signature-resp not found")
		}
		publicKey, err := getRsaPublicKey(cli)
		if err != nil {
			hlog.Errorf("AuthMiddleware get publicKey failed: %+v", err)
			return err
		}
		err = verify(publicKey, []byte(value), response)
		if err != nil {
			hlog.Errorf("AuthMiddleware verify failed: %+v", err)
			return err
		}
		return nil
	}
}

func AuthenticateServer(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		hlog.Infof("AuthMiddleware get req: %+v", request)
		// 创建etcd客户端连接
		err := getEtcdClient(once)
		if err != nil {
			hlog.Errorf("AuthMiddleware get etcd client failed: %+v", err)
			return err
		}
		key, err := getRsaPublicKey(cli)
		if err != nil {
			hlog.Errorf("AuthMiddleware get public failed: %+v", err)
			return err
		}

		// 验签
		value, ok := metainfo.GetValue(ctx, SignReqHeader)
		if !ok {
			return errors.New("signature-req not found")
		}
		requestStr := fmt.Sprintf("%s", request)
		err = verify(key, []byte(value), requestStr)
		if err != nil {
			hlog.Errorf("AuthMiddleware verify failed: %+v", err)
			return err
		}
		// 调用下一个中间件
		err = next(ctx, request, response)
		if err != nil {
			return err
		}
		// 响应签名
		privateKey, err := getRsaPrivateKey(cli)
		if err != nil {
			hlog.Errorf("AuthMiddleware get privateKey failed: %+v", err)
			return err
		}

		signature, err := sign(privateKey, response)
		if err != nil {
			hlog.Errorf("AuthMiddleware sign failed: %+v", err)
			return err
		}

		of := metainfo.SendBackwardValue(ctx, SignRespHeader, string(signature))
		if !of {
			return errors.New("sendBackwardValue failed")
		}
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

// 对数据进行签名
func sign(privateKey *rsa.PrivateKey, data interface{}) ([]byte, error) {
	hashedData := sha256.Sum256([]byte(fmt.Sprintf("%v", data)))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedData[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %v", err)
	}
	return signature, nil
}

// 验证数据的签名
func verify(publicKey *rsa.PublicKey, signature []byte, data interface{}) error {
	hashedData := sha256.Sum256([]byte(fmt.Sprintf("%v", data)))
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData[:], signature)
	if err != nil {
		return fmt.Errorf("failed to verify signature: %v", err)
	}
	return nil
}

func getAESKey(cli *clientv3.Client) (string, error) {
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

func getRsaPrivateKey(cli *clientv3.Client) (*rsa.PrivateKey, error) {
	// 使用客户端从etcd中获取密钥
	resp, err := cli.Get(context.Background(), PrivateKeyKey)
	if err != nil {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, err
	}

	// 检查是否找到了密钥
	if len(resp.Kvs) == 0 {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, errors.New("AuthMiddleware get key failed")
	}

	// 获取密钥
	keyValue := resp.Kvs[0].Value
	// 解码密钥
	block, _ := pem.Decode(keyValue)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, err
	}
	return privateKey, nil
}

func getRsaPublicKey(cli *clientv3.Client) (*rsa.PublicKey, error) {
	// 使用客户端从etcd中获取密钥
	resp, err := cli.Get(context.Background(), PublicKeyKey)
	if err != nil {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, err
	}

	// 检查是否找到了密钥
	if len(resp.Kvs) == 0 {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, errors.New("AuthMiddleware get key failed")
	}

	// 获取密钥
	keyValue := resp.Kvs[0].Value

	// 解码密钥
	block, _ := pem.Decode(keyValue)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		hlog.Errorf("AuthMiddleware get key failed: %+v", err)
		return nil, err
	}
	return publicKey, nil
}

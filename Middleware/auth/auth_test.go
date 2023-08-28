package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"middleware/internal"
	"os"
	"sync"
	"testing"
	"time"
)

func TestAuthenticateClient(t *testing.T) {
	testEndpoint := func(ctx context.Context, req interface{}, resp interface{}) (err error) {
		_ = getEtcdClient(once)
		key, err := getRsaPublicKey(cli)
		value := ctx.Value(SignReqHeader)
		err = verify(key, value.([]byte), req)
		assert.Nil(t, err)
		resp = &internal.UserInfoResp{
			StatusCode: 200,
			StatusMsg:  "Lalala",
		}
		privateKey, _ := getRsaPrivateKey(cli)
		signature, _ := sign(privateKey, resp)
		ctx = context.WithValue(ctx, SignRespHeader, signature)
		return nil
	}

	mw := AuthenticateClient(testEndpoint)
	ctx := context.Background()
	req := &internal.UserInfoReq{
		UserId: 342,
	}
	resp := &internal.UserInfoResp{
		StatusCode: 200,
		StatusMsg:  "Lalala",
	}
	err := mw(ctx, req, resp)
	assert.NotNil(t, err)
}

func TestEncryptReqAndDecryptResp(t *testing.T) {
	req := "Hello World!"
	key := "W0zBXMY7VL7Xo6s0"

	i, err := encrypt([]byte(req), key)
	assert.Nil(t, err)
	t.Logf("encrypted req: %s", i)
	resp, err := decrypt(i, key)
	assert.Nil(t, err)
	t.Logf("decrypted resp: %s", resp)
	assert.Assert(t, req == string(resp))
}

func TestGetEtcdClient(t *testing.T) {
	err := getEtcdClient(once)
	assert.NotNil(t, cli)
	assert.Nil(t, err)
	get, err := cli.Get(context.Background(), "AESKEY")
	t.Log(get.Kvs[0].Value)
	time.Sleep(10 * time.Minute)
	get2, err := cli.Get(context.Background(), "AESKEY")
	assert.Nil(t, err)
	t.Log(get2.Kvs[0].Value)
}

func TestSignAndVerify(t *testing.T) {
	keyData, _ := os.ReadFile("../private_key.pem")
	block, _ := pem.Decode(keyData)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	data := "Hello World!"
	bytes, _ := sign(key, data)
	s := string(bytes)
	keyData2, _ := os.ReadFile("../public_key.pem")
	block2, _ := pem.Decode(keyData2)
	key2, _ := x509.ParsePKCS1PublicKey(block2.Bytes)
	err := verify(key2, []byte(s), data)
	assert.Nil(t, err)
}

func TestGenerateRSAKeyPair(t *testing.T) {
	// 生成密钥对
	bits := 2048
	privateKeyFile := "../private_key.pem"
	publicKeyFile := "../public_key.pem"
	// 如果存在文件，则直接退出
	if _, err := os.Stat(privateKeyFile); err == nil {
		t.Log("private key file exists")
		return
	}
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)

	// 将私钥保存到文件
	privateKeyFileData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	_ = os.WriteFile(privateKeyFile, privateKeyFileData, 0600)

	// 获取公钥
	publicKey := &privateKey.PublicKey

	// 将公钥保存到文件
	publicKeyFileData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
	_ = os.WriteFile(publicKeyFile, publicKeyFileData, 0644)
}

func TestPutRSA2Etcd(t *testing.T) {
	_ = getEtcdClient(&sync.Once{})
	privateKey, _ := os.ReadFile("../private_key.pem")
	put, err := cli.Put(context.Background(), "PrivateKey", string(privateKey))
	assert.Nil(t, err)
	t.Log(put)
	publicKey, _ := os.ReadFile("../public_key.pem")
	put2, err := cli.Put(context.Background(), "PublicKey", string(publicKey))
	assert.Nil(t, err)
	t.Log(put2)
}

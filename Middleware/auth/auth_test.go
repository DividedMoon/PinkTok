package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"time"
)

func TestAuthenticateMiddleware(t *testing.T) {
	t.Log("TestAuthenticateMiddleware")
	testEndpoint := func(ctx context.Context, req interface{}, resp interface{}) (err error) {
		return nil
	}

	mw := AuthenticateMiddleware(testEndpoint)
	ctx := context.Background()
	req := "test"
	resp := "OK"
	err := mw(ctx, req, resp)
	if err != nil {
		t.Errorf("AuthenticateMiddleware failed: %+v", err)
	}
}

func TestEncryptReqAndDecryptResp(t *testing.T) {
	req := "hello"
	key := "W0zBXMY7VL7Xo6s0"
	i, err := encryptReq(req, key)
	assert.Nil(t, err)
	t.Logf("encrypted req: %+v", i)
	resp, err := decryptResp(i, key)
	assert.Nil(t, err)
	t.Logf("decrypted resp: %+v", resp)
	assert.Assert(t, req == resp)
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

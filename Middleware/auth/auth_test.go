package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"middleware/internal"
	"testing"
	"time"
)

func TestAuthenticateClient(t *testing.T) {
	testEndpoint := func(ctx context.Context, req interface{}, resp interface{}) (err error) {
		t.Log("req: ", req)
		t.Log("testEndpoint")
		t.Log("resp: ", resp)
		return nil
	}

	mw := AuthenticateClient(testEndpoint)
	ctx := context.Background()
	req := &internal.UserInfoReq{
		UserId: 342,
	}
	resp := &internal.UserInfoResp{}
	err := mw(ctx, req, resp)
	if err != nil {
		t.Errorf("AuthenticateMiddleware failed: %+v", err)
	}
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

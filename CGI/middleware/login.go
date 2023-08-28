package middleware

import (
	internalClient "cgi/internal/client"
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"user_service/biz"
)

func LoginHandler(ctx context.Context, req userLoginReq) (resp *UserLoginResp, err error) {
	// 请求userService
	loginReq := &biz.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	loginResp, err := internalClient.UserServiceClient.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	return &UserLoginResp{
		StatusCode: loginResp.StatusCode,
		StatusMsg:  loginResp.StatusMsg,
		UserId:     int(loginResp.UserId),
	}, nil
}

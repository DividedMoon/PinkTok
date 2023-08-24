package middleware

import (
	internalClient "cgi/internal/client"
	"context"
	"user_service/biz"
)

func LoginHandler(ctx context.Context, req userLoginReq) (resp *UserLoginResp, err error) {
	// 请求userService
	loginReq := &biz.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	}
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

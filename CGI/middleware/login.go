package middleware

import (
	cgi "cgi/internal/client"
	"context"
	"user_service/biz/model/client"
)

func LoginHandler(ctx context.Context, req userLoginReq) (resp *UserLoginResp, err error) {
	// 请求userService
	loginReq := &client.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	loginResp, _, err := cgi.UserServiceClient.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	return &UserLoginResp{
		StatusCode: loginResp.StatusCode,
		StatusMsg:  loginResp.StatusMsg,
		UserId:     int(loginResp.UserId),
	}, nil
}

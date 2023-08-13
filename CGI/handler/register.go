package handler

import (
	internalClient "cgi/internal/client"
	"cgi/internal/constant"
	"cgi/internal/utils"
	"cgi/middleware"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"user_service/biz"
)

type userRegisterReq struct {
	Username string `query:"username"`
	Password string `query:"password"`
}

type userRegisterResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

// RegisterHandler is the handler of the route "/douyin/user/register/".
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  userRegisterReq
		resp *utils.BaseResp
	)
	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, userRegisterResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	// 加密密码
	cryPwd, err := utils.Crypt(req.Password)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, userRegisterResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	// 请求userService
	registerReq := &biz.UserRegisterReq{
		Username: req.Username,
		Password: cryPwd,
	}
	hlog.CtxInfof(ctx, "request user_service : %+v with %+v", registerReq, c.ClientIP())
	registerResp, err := internalClient.UserServiceClient.Register(ctx, registerReq)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, userRegisterResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "response user_service: %+v with %+v", registerResp, c.ClientIP())

	token, _, err := middleware.JwtMiddleware.TokenGenerator(registerResp.UserId)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, userRegisterResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, userRegisterResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     int(registerResp.UserId),
		Token:      token,
	})
}

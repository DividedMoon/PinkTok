package handler

import (
	"cgi/internal/constant"
	"cgi/internal/utils"
	"client/client/user_service"
	"client/dto"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

type userRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	registerReq := &dto.UserRegisterReq{
		Username: req.Username,
		Password: cryPwd,
	}
	_, _, err = user_service.Register(ctx, registerReq)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, userRegisterResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	LoginHandler(ctx, c)
	token := c.GetString("token")
	v, _ := c.Get(jwt.IdentityKey)
	userId := v.(int)
	c.JSON(consts.StatusOK, userRegisterResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     int(userId),
		Token:      token,
	})
}

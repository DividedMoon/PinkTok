// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"user_service/biz/internal/constant"
	"user_service/biz/internal/service"
	"user_service/biz/model/client"
)

// Register .
// @router /internal/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var (
		err error
		req client.UserRegisterReq
	)

	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &client.UserRegisterResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  err.Error(),
			UserId:     -1,
		})
		return
	}
	hlog.CtxInfof(ctx, "request: %+v", req.Username)
	// 新建用户
	u, err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		hlog.CtxErrorf(ctx, "register user error: %+v", err)
		c.JSON(consts.StatusOK, &client.UserRegisterResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  err.Error(),
			UserId:     -1,
		})
		return
	}
	c.JSON(consts.StatusOK, &client.UserRegisterResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     u.Id,
	})
}

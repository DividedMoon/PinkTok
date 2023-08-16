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
	"relation_service/biz"
)

type relationActionReq struct {
	Token      string `query:"token"`
	ToUserId   string `query:"to_user_id"`
	ActionType string `query:"action_type"`
}

type relationActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func RelationActionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  relationActionReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, relationActionResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "get current user id failed, something wrong in authorize",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, relationActionResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  "userId cannot be parsed to int64",
		})
		return
	}
	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relationActionResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/action with req: %+v from: %+v", req, c.ClientIP())

	// 调用RelationService
	actionReq := &biz.RelationActionReq{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: 0,
	}
	internalClient.RelationServiceClient.SendRelationAction(ctx)
}

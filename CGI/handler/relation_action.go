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
	"strconv"
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
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	var userId int64
	userIdFloat, ok := v.(float64)
	if ok {
		userId = int64(userIdFloat)
	} else {
		userId, ok = v.(int64)
		if !ok {
			hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
			c.JSON(consts.StatusInternalServerError, relationActionResp{
				StatusCode: constant.ParamErrCode,
				StatusMsg:  constant.ParamErrMsg,
			})
			return
		}
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

	// 参数校验
	toUserId, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, "toUserId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, relationActionResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
	}
	actionType, err := strconv.ParseInt(req.ActionType, 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		hlog.CtxErrorf(ctx, "actionType not formatted")
		c.JSON(consts.StatusInternalServerError, relationActionResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
	}

	// 调用RelationService
	actionReq := &biz.RelationActionReq{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	}
	hlog.CtxInfof(ctx, "sendRelationAction with req: %+v", actionReq)
	actionResp, err := internalClient.RelationServiceClient.SendRelationAction(ctx, actionReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "relation/action failed with err: %+v", err)
		c.JSON(consts.StatusInternalServerError, relationActionResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/action with resp: %+v", actionResp)
	c.JSON(consts.StatusOK, relationActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  actionResp.StatusMsg,
	})
}

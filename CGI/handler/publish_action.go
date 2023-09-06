package handler

import (
	internalClient "cgi/internal/client"
	"cgi/internal/constant"
	"cgi/internal/utils"
	"cgi/middleware"
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"video_service/biz"
)

type publishActionReq struct {
	Token string `query:"token"`
	Data  []byte `query:"data"`
	Title string `query:"title"`
}

type publishActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func PublishActionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  publishActionReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, publishActionResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, publishActionResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, publishActionResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/follower/list with req: %+v from: %+v", req, c.ClientIP())

	// 调用 VideoService
	pReq := &biz.PublishReq{
		UserId: userId,
		Data:   req.Data,
		Title:  req.Title,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	pResp, err := internalClient.VideoServiceClient.PublishVideo(ctx, pReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get follower list failed, err: %+v", err)
		c.JSON(consts.StatusInternalServerError, publishActionResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get follower list success, resp: %+v", pResp)

	c.JSON(consts.StatusOK, publishActionResp{
		StatusCode: pResp.StatusCode,
		StatusMsg:  pResp.StatusMsg,
	})
}

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

type publicListReq struct {
	UserId string `query:"user_id"`
	Token  string `query:"token"`
}

type publicListResp struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	VideoList  []biz.VideoInfo `json:"video_list"`
}

func PublicListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  publicListReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, publicListResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, publicListResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, publicListResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/follower/list with req: %+v from: %+v", req, c.ClientIP())

	// 调用interact service
	plReq := &biz.GetPublishListReq{
		UserId: userId,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	plResp, err := internalClient.VideoServiceClient.GetPublishList(ctx, plReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get follower list failed, err: %+v", err)
		c.JSON(consts.StatusInternalServerError, publicListResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get follower list success, resp: %+v", plResp)
	vl := make([]biz.VideoInfo, 0)
	for _, p := range plResp.VideoList {
		if p != nil {
			vl = append(vl, biz.VideoInfo{
				Id:            p.Id,
				Author:        p.Author,
				PlayUrl:       p.PlayUrl,
				CoverUrl:      p.CoverUrl,
				FavoriteCount: p.FavoriteCount,
				CommentCount:  p.CommentCount,
				IsFavorite:    p.IsFavorite,
				Title:         p.Title,
			})
		}
	}
	c.JSON(consts.StatusOK, publicListResp{
		StatusCode: plResp.StatusCode,
		StatusMsg:  plResp.StatusMsg,
		VideoList:  vl,
	})
}

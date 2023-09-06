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

type favoriteListReq struct {
	UserId string `query:"user_id"`
	Token  string `query:"token"`
}

type favoriteListResp struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	VideoList  []biz.VideoInfo `json:"video_list"`
}

func FavoriteListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  favoriteListReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, favoriteListResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, favoriteListResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, favoriteListResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/follower/list with req: %+v from: %+v", req, c.ClientIP())

	// 调用 v service
	fvr := &biz.GetFavoriteVideoListReq{
		UserId: userId,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	fvrsp, err := internalClient.VideoServiceClient.GetFavoriteVideoList(ctx, fvr)
	if err != nil {
		hlog.CtxErrorf(ctx, "get follower list failed, err: %+v", err)
		c.JSON(consts.StatusInternalServerError, favoriteListResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get follower list success, resp: %+v", fvrsp)
	vl := make([]biz.VideoInfo, 0)
	for _, p := range fvrsp.VideoList {
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
	c.JSON(consts.StatusOK, favoriteListResp{
		StatusCode: fvrsp.StatusCode,
		StatusMsg:  fvrsp.StatusMsg,
		VideoList:  vl,
	})
}

package handler

import (
	internalClient "cgi/internal/client"
	"cgi/internal/utils"
	"cgi/middleware"
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
	"video_service/biz"
)

type feedReq struct {
	LastTime int64  `query:"last_time"`
	Token    string `query:"token"`
}

type feedResp struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	VideoList  []biz.VideoInfo `json:"video_list"`
	NextTime   int64           `json:"next_time"`
}

func GetFeedHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  feedReq
		resp *utils.BaseResp
	)
	err = c.Bind(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, feedResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get feed request, last time: %d", req.LastTime)

	var userId int64
	if req.Token != "" {
		token, err := middleware.JwtMiddleware.ParseTokenString(req.Token)
		if err != nil {
			hlog.CtxErrorf(ctx, "parse token failed, err: %v", err)
			resp = utils.BuildBaseResp(err)
			c.JSON(consts.StatusOK, feedResp{
				StatusCode: resp.StatusCode,
				StatusMsg:  resp.StatusMsg,
			})
			return
		}
		userId = token.Header["user_id"].(int64)
	}
	userId = -1
	if req.LastTime == 0 {
		req.LastTime = time.Now().Unix()
	}
	feedReqq := &biz.FeedReq{
		LatestTime: req.LastTime,
		UserId:     userId,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	feedRespp, err := internalClient.VideoServiceClient.Feed(ctx, feedReqq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get feed failed, err: %v", err)
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, feedResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get feed success, video list: %v", feedRespp.VideoList)
	var videoList []biz.VideoInfo
	for _, v := range feedRespp.VideoList {
		videoList = append(videoList, biz.VideoInfo{
			Id:            v.Id,
			Author:        v.Author,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	c.JSON(consts.StatusOK, feedResp{
		StatusCode: feedRespp.StatusCode,
		StatusMsg:  feedRespp.StatusMsg,
		VideoList:  videoList,
	})
}

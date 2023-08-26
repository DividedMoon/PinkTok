package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"video_service/biz"
	"video_service/internal/service"
	"video_service/internal/utils"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *biz.FeedReq) (resp *biz.FeedResp, err error) {
	resp, err = service.NewVideoService(ctx).GetFeed(req)

	res := utils.BuildBaseResp(err)
	hlog.CtxErrorf(ctx, "Feed Error", err.Error())
	return &biz.FeedResp{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
		VideoList:  nil,
		NextTime:   req.LatestTime,
	}, err

}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *biz.PublishReq) (resp *biz.PublishResp, err error) {
	err = service.NewVideoService(ctx).PublishAction(req)
	res := utils.BuildBaseResp(err)
	return &biz.PublishResp{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, err

}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *biz.GetPublishListReq) (resp *biz.GetPublishListResp, err error) {
	resp, err = service.NewVideoService(ctx).GetPublishList(req)
	res := utils.BuildBaseResp(err)
	hlog.CtxErrorf(ctx, "GetPublishList Error", err.Error())
	return &biz.GetPublishListResp{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, err

}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *biz.GetFavoriteVideoListReq) (resp *biz.GetFavoriteVideoListResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionReq) (resp *biz.FavoriteActionResp, err error) {
	// TODO: Your code here...
	return
}

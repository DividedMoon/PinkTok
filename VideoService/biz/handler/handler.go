package handler

import (
	"context"
	"video_service/biz"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *biz.FeedReq) (resp *biz.FeedResp, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *biz.PublishReq) (resp *biz.PublishResp, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *biz.GetPublishListReq) (resp *biz.GetPublishListResp, err error) {
	// TODO: Your code here...
	return
}

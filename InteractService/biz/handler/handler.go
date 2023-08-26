package handler

import (
	"context"
	"interact_service/biz"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// QueryFavoriteExist implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryFavoriteExist(ctx context.Context, req *biz.QueryFavoriteExistReq) (resp *biz.QueryFavoriteExistResp, err error) {
	// TODO: Your code here...
	return
}

// QueryUserFavoriteVideoIds implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryUserFavoriteVideoIds(ctx context.Context, req *biz.FavoriteVideoReq) (resp *biz.FavoriteVideoResp, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *biz.CommentActionReq) (resp *biz.CommentActionResp, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *biz.CommentListReq) (resp *biz.CommentListResp, err error) {
	// TODO: Your code here...
	return
}

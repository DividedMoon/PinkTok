package handler

import (
	"context"
	"interact_service/biz"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionReq) (resp *biz.FavoriteActionResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *biz.FavoriteListReq) (resp *biz.FavoriteListResp, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *biz.CommentActionReq) (resp *biz.CommentActionResp, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *biz.CommentListReq) (resp *biz.DouyinCommentListResp, err error) {
	// TODO: Your code here...
	return
}

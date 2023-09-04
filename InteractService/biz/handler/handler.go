package handler

import (
	"context"
	"interact_service/biz"
	"interact_service/internal/constant"
	"interact_service/internal/model"
	"interact_service/internal/service"
	utils "interact_service/internal/util"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// QueryFavoriteExist implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryFavoriteExist(ctx context.Context, req *biz.QueryFavoriteExistReq) (resp *biz.QueryFavoriteExistResp, err error) {
	userId, videoId := req.UserId, req.VideoId
	liked, err := service.QueryFavoriteExist(userId, videoId)

	bresp := utils.BuildBaseResp(err)
	return &biz.QueryFavoriteExistResp{
		IsFavorite: liked,
		StatusCode: bresp.StatusCode,
		StatusMsg:  bresp.StatusMsg,
	}, err
}

// QueryUserFavoriteVideoIds implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryUserFavoriteVideoIds(ctx context.Context, req *biz.FavoriteVideoReq) (resp *biz.FavoriteVideoResp, err error) {
	userId := req.UserId
	videoIds, err := service.QueryUserFavoriteVideoIds(userId)

	bresp := utils.BuildBaseResp(err)
	return &biz.FavoriteVideoResp{
		VideoList:  videoIds,
		StatusCode: bresp.StatusCode,
		StatusMsg:  bresp.StatusMsg,
	}, err
}

// AddFavoriteRecord implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) AddFavoriteRecord(ctx context.Context, req *biz.AddFavoriteRecordReq) (resp *biz.AddFavoriteRecordResp, err error) {
	userId, videoId, actionType := req.UserId, req.VideoId, req.ActionType
	err = service.AddFavoriteRecord(userId, videoId, actionType)
	bresp := utils.BuildBaseResp(err)
	return &biz.AddFavoriteRecordResp{
		StatusCode: bresp.StatusCode,
		StatusMsg:  bresp.StatusMsg,
	}, err
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *biz.CommentActionReq) (*biz.CommentActionResp, error) {
	userId, videoId, actionType := req.UserId, req.VideoId, req.ActionType
	var err error
	var commentResp *biz.Comment = nil
	if actionType == 1 {
		commentText := req.CommentText
		comment := &model.Comment{
			UserId:  userId,
			VideoId: videoId,
			Content: commentText,
		}
		commentResp, err = service.PublishComment(ctx, comment)
	} else if actionType == 2 {
		commentId := req.CommentId
		comment := &model.Comment{
			ID:      commentId,
			UserId:  userId,
			VideoId: videoId,
		}
		err = service.DeleteComment(comment)
	} else {
		err = constant.NewErrNo(constant.InvalidActionTypeCode, constant.InvalidActionTypeMsg)
	}
	bresp := utils.BuildBaseResp(err)
	return &biz.CommentActionResp{
		Comment:    commentResp,
		StatusCode: bresp.StatusCode,
		StatusMsg:  bresp.StatusMsg,
	}, err
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *biz.CommentListReq) (resp *biz.CommentListResp, err error) {
	userId, videoId := req.UserId, req.VideoId
	comments, err := service.GetCommentByUserAndVideo(ctx, userId, videoId)
	bresp := utils.BuildBaseResp(err)

	return &biz.CommentListResp{
		CommentList: comments,
		StatusCode:  bresp.StatusCode,
		StatusMsg:   bresp.StatusMsg,
	}, err
}

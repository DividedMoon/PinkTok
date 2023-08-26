package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"interact_service/biz"
	internalClient "interact_service/internal/client"
	"interact_service/internal/model"
	userBiz "user_service/biz"
)

func PublishComment(ctx context.Context, comment *model.Comment) (resp *biz.Comment, err error) {
	// 唯一索引保证重入安全
	err = comment.Create()
	if err != nil {
		return nil, err
	}
	userInfoReq := &userBiz.UserInfoReq{
		UserId: comment.UserId,
	}
	// TODO 添加视频commentCount
	userInfoResp, err := internalClient.UserServiceClient.UserInfo(ctx, userInfoReq)
	if err != nil {
		return nil, err
	}
	user := &biz.UserInfo{
		Id:              userInfoResp.User.Id,
		Name:            userInfoResp.User.Name,
		FollowCount:     userInfoResp.User.FollowCount,
		FollowerCount:   userInfoResp.User.FollowerCount,
		IsFollow:        userInfoResp.User.IsFollow,
		Avatar:          userInfoResp.User.Avatar,
		BackgroundImage: userInfoResp.User.BackgroundImage,
		Signature:       userInfoResp.User.Signature,
		TotalFavorited:  userInfoResp.User.TotalFavorited,
		WorkCount:       userInfoResp.User.WorkCount,
		FavoriteCount:   userInfoResp.User.FavoriteCount,
	}
	resp = &biz.Comment{
		Id:         comment.ID,
		User:       user,
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02"),
	}
	return resp, nil
}

func DeleteComment(comment *model.Comment) error {
	return comment.DeleteById()
}

func GetCommentByUserAndVideo(userId, videoId int64) ([]model.Comment, error) {
	var (
		c = &model.Comment{
			UserId:  userId,
			VideoId: videoId,
		}
	)
	list, err := c.SelectByUserIdAndVideoId()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FavoriteAction 用户点赞或者取消赞
func FavoriteAction(userId, videoId int64) error {
	liked, err := model.IsVideoLikedByUser(userId, videoId)
	if err != nil {
		hlog.Error("IsVideoLikedByUser error", err)
		return err
	}
	// TODO 调用videoService查询点赞数量
	err = model.UpdateVideoLikedStatus(userId, videoId, !liked)
	// TODO 调用videoService更新video表里的点赞数量

	if err != nil {
		hlog.Error("UpdateVideoLikedStatus error", err)
		return err
	}
	return nil
}

// QueryFavoriteExist 查询用户是否点赞
func QueryFavoriteExist(userId, videoId int64) (bool, error) {
	liked, err := model.IsVideoLikedByUser(userId, videoId)
	if err != nil {
		hlog.Error("IsVideoLikedByUser error", err)
		return false, err
	}
	return liked, nil
}

// QueryUserFavoriteVideoIds 查询用户的喜欢列表，返回喜欢的视频ID
func QueryUserFavoriteVideoIds(userId int64) ([]int64, error) {
	videoIds, err := model.SelectFavoriteVideoIdsByUserID(userId)
	if err != nil {
		hlog.Error("SelectFavoriteVideoIdsByUserID error", err)
		return nil, err
	}
	return videoIds, nil
}

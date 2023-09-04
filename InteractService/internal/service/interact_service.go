package service

import (
	"context"
	"fmt"
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

func GetCommentByUserAndVideo(ctx context.Context, userId, videoId int64) ([]*biz.Comment, error) {
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
	respList := make([]*biz.Comment, 0)
	userInfoReq := &userBiz.UserInfoReq{
		UserId: userId,
	}
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
	for _, comment := range list {

		respList = append(respList, &biz.Comment{
			Id:         comment.ID,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}

	return respList, nil
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

// AddFavoriteRecord 添加用户喜欢记录
func AddFavoriteRecord(userId, videoId, actionType int64) error {
	var err error
	if actionType == 1 { //点赞
		liked, err := model.IsVideoLikedByUser(videoId, userId)
		if liked {
			return err
		} else {
			err = model.UpdateVideoLikedStatus(userId, videoId, liked)
		}
	} else if actionType == 2 { //取消赞
		liked, _ := model.IsVideoLikedByUser(videoId, userId)
		if !liked {
			return fmt.Errorf("cancel liked but not liked")
		}
		err = model.UpdateVideoLikedStatus(userId, videoId, liked)
	} else {
		return fmt.Errorf("actionType error")
	}
	if err != nil {
		hlog.Error("InsertFavoriteVideo error", err)
		return err
	}
	return nil
}

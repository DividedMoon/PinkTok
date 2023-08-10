package service

import (
	"context"
	"gorm.io/gorm"
	relationClient "relation_service/biz/internal/client"
	"relation_service/biz/model"
	"relation_service/biz/model/client"
	"relation_service/biz/model/dto"
	userClient "user_service/biz/model/client"
)

func SubmitFollowRelationAction(ctx context.Context, req *client.RelationActionReq) (err error) {
	var (
		follow = &dto.Follow{
			UserIdA: req.UserId,
			UserIdB: req.ToUserId,
		}
		follower = &dto.Follower{
			UserIdA: req.ToUserId,
			UserIdB: req.UserId,
		}
		userAReq = &userClient.UserInfoReq{UserId: req.UserId}
		userBReq = &userClient.UserInfoReq{UserId: req.ToUserId}
	)
	// 先查询用户信息
	userAResp, _, err := relationClient.UserServiceClient.UserInfo(ctx, userAReq)
	if err != nil {
		return err
	}
	userBResp, _, err := relationClient.UserServiceClient.UserInfo(ctx, userBReq)
	if err != nil {
		return err
	}
	var ()

	// 事务更新
	err = model.DB.Transaction(func(tx *gorm.DB) error {
		// 创建关注关系
		if err = follow.Create(tx); err != nil {
			return err
		}
		// 创建粉丝关系
		if err = follower.Create(tx); err != nil {
			return err
		}
		// 更新用户信息
		if err = userAResp.Update(tx); err != nil {
			return err
		}
		if err = userBResp.Update(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func CancelFollowRelationAction(ctx context.Context, req *client.RelationActionReq) (err error) {

}

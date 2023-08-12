package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	relationClient "relation_service/biz/internal/client"
	"relation_service/biz/internal/constant"
	"relation_service/biz/model"
	_ "relation_service/biz/model"
	"relation_service/biz/model/client"
	"relation_service/biz/model/dto"
	"relation_service/biz/mw/redis"
	"sync"
	userClient "user_service/biz/model/client"
)

// SubmitFollowRelationAction 提交关注关系
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

	// 1. 先查询关注关系和粉丝关系是否存在
	var (
		existFollow   = redis.ExistFollow(req.UserId, req.ToUserId)
		existFollower = redis.ExistFollower(req.UserId, req.ToUserId)
	)

	if existFollow && existFollower {
		// 已经存在关注关系和粉丝关系
		return
	}

	// 2. 再查询用户信息
	userAResp, _, err := relationClient.UserServiceClient.UserInfo(ctx, userAReq)
	if err != nil {
		return err
	}
	if constant.SuccessCode != userAResp.StatusCode {
		return constant.UserNotExistErr
	}
	userBResp, _, err := relationClient.UserServiceClient.UserInfo(ctx, userBReq)
	if err != nil {
		return err
	}
	if constant.SuccessCode != userBResp.StatusCode {
		return constant.UserNotExistErr
	}

	// 3. 接着进行事务更新
	err = model.DB.Transaction(func(tx *gorm.DB) error {
		// 创建关注关系
		if err = follow.Create(tx); err != nil {
			return err
		}
		// 创建粉丝关系
		if err = follower.Create(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 4. 然后更新redis
	if !existFollow && existFollower {
		// 存在粉丝关系，不存在关注关系
		// 创建关注关系
		redis.AddFollow(req.UserId, req.ToUserId)
	} else if existFollow && !existFollower {
		// 存在关注关系，不存在粉丝关系
		// 创建粉丝关系
		redis.AddFollower(req.ToUserId, req.UserId)
	} else {
		// 不存在关注关系和粉丝关系
		// 创建关注关系
		redis.AddFollow(req.UserId, req.ToUserId)
		// 创建粉丝关系
		redis.AddFollower(req.UserId, req.ToUserId)
	}

	// 5. 最后更新用户关注数和粉丝数
	userAResp.User.FollowCount++
	userBResp.User.FollowerCount++
	var (
		wg    = sync.WaitGroup{}
		userA = &userClient.UserUpdateReq{
			UserId: userAResp.User.Id,
		}
		userB = &userClient.UserUpdateReq{
			UserId: userBResp.User.Id,
		}
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		updateAResp, _, err := relationClient.UserServiceClient.UserUpdate(ctx, userA)
		if err != nil {
			hlog.CtxErrorf(ctx, "update user %+v follow count failed, err: %v", userA, err)
			return
		}
		hlog.CtxInfof(ctx, "update user %+v follow count success, resp: %+v", userA, updateAResp)
	}()
	go func() {
		defer wg.Done()
		updateBResp, _, err := relationClient.UserServiceClient.UserUpdate(ctx, userB)
		if err != nil {
			hlog.CtxErrorf(ctx, "update user %+v follow count failed, err: %v", userB, err)
			return
		}
		hlog.CtxInfof(ctx, "update user %+v follow count success, resp: %+v", userB, updateBResp)
	}()
	wg.Wait()
	hlog.CtxInfof(ctx, "submit follow relation action success, req: %+v", req)
	return nil
}

// CancelFollowRelationAction 取消关注关系
func CancelFollowRelationAction(ctx context.Context, req *client.RelationActionReq) (err error) {
	return constant.ServiceErr
}

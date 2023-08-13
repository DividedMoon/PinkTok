package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"relation_service/biz"
	internalClient "relation_service/internal/client"
	"relation_service/internal/constant"
	"relation_service/internal/model"
	"relation_service/internal/mw/redis"
	"sync"
	userBiz "user_service/biz"
)

// SubmitFollowRelationAction 提交关注关系
func SubmitFollowRelationAction(ctx context.Context, req *biz.RelationActionReq) (err error) {
	hlog.CtxInfof(ctx, "SubmitFollowRelationAction req:%+v", req)
	var (
		follow = &model.Follow{
			UserIdA: req.UserId,
			UserIdB: req.ToUserId,
		}
		follower = &model.Follower{
			UserIdA: req.ToUserId,
			UserIdB: req.UserId,
		}
		userAReq = &userBiz.UserInfoReq{UserId: req.UserId}
		userBReq = &userBiz.UserInfoReq{UserId: req.ToUserId}
	)

	// 1. 先查询关注关系和粉丝关系是否存在
	var (
		existFollow   = redis.ExistFollow(req.UserId, req.ToUserId)
		existFollower = redis.ExistFollower(req.UserId, req.ToUserId)
	)
	hlog.CtxInfof(ctx, "CheckFromRedis existFollow:%v, existFollower:%v", existFollow, existFollower)

	if existFollow && existFollower {
		// 已经存在关注关系和粉丝关系
		return
	}

	// 2. 再查询用户信息
	hlog.CtxInfof(ctx, "CheckUserInfo userAReq:%+v, userBReq:%+v", userAReq, userBReq)
	userAResp, err := internalClient.UserServiceClient.UserInfo(ctx, userAReq)
	if err != nil {
		return err
	}
	if constant.SuccessCode != userAResp.StatusCode {
		return constant.UserNotExistErr
	}

	userBResp, err := internalClient.UserServiceClient.UserInfo(ctx, userBReq)
	if err != nil {
		return err
	}
	if constant.SuccessCode != userBResp.StatusCode {
		return constant.UserNotExistErr
	}

	// 3. 接着进行事务更新
	hlog.CtxInfof(ctx, "CheckFromDB follow:%+v, follower:%+v", follow, follower)
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
		hlog.CtxErrorf(ctx, "InsertDB err:%v", err)
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
		redis.AddFollower(req.ToUserId, req.UserId)
	}

	// 5. 最后更新用户关注数和粉丝数
	var (
		wg    = sync.WaitGroup{}
		userA = &userBiz.UserUpdateReq{
			UserId: userAResp.User.Id,
			Ext: []*userBiz.PairValue{
				{
					Key:   "FollowCount",
					Value: "1",
				},
			},
		}
		userB = &userBiz.UserUpdateReq{
			UserId: userBResp.User.Id,
			Ext: []*userBiz.PairValue{
				{
					Key:   "FollowerCount",
					Value: "1",
				},
			},
		}
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		hlog.CtxInfof(ctx, "update user %+v follow count", userA)
		updateAResp, err := internalClient.UserServiceClient.UserUpdate(ctx, userA)
		if err != nil {
			hlog.CtxErrorf(ctx, "update user %+v follow count failed, err: %v", userA, err)
			return
		}
		hlog.CtxInfof(ctx, "update user %+v follow count success, resp: %+v", userA, updateAResp)
	}()
	go func() {
		defer wg.Done()
		hlog.CtxInfof(ctx, "update user %+v follower count", userB)
		updateBResp, err := internalClient.UserServiceClient.UserUpdate(ctx, userB)
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
func CancelFollowRelationAction(ctx context.Context, req *biz.RelationActionReq) (err error) {
	return constant.ServiceErr
}

func GetFollowListById(ctx context.Context, userId int64) ([]*biz.UserInfo, error) {
	// 1. 查询db
	followList, err := model.SelectFollowByUserIdA(userId)
	if err != nil {
		return nil, err
	}
	// 2. 查询用户信息
	data := make([]*biz.UserInfo, len(followList))
	for i, follow := range followList {
		userInfoReq := &userBiz.UserInfoReq{UserId: follow.UserIdB}
		userInfoResp, err := internalClient.UserServiceClient.UserInfo(ctx, userInfoReq)
		if err != nil {
			return nil, err
		}
		if constant.SuccessCode != userInfoResp.StatusCode {
			return nil, constant.UserNotExistErr
		}
		u := userInfoResp.User
		data[i] = &biz.UserInfo{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        true,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
	}
	return data, nil
}

func GetFollowerListById(ctx context.Context, userId int64) ([]*biz.UserInfo, error) {
	// 1. 查询db
	followerList, err := model.SelectFollowerByUserIdA(userId)
	if err != nil {
		return nil, err
	}
	// 2. 查询用户信息
	data := make([]*biz.UserInfo, len(followerList))
	for i, follower := range followerList {
		userInfoReq := &userBiz.UserInfoReq{UserId: follower.UserIdB}
		userInfoResp, err := internalClient.UserServiceClient.UserInfo(ctx, userInfoReq)
		if err != nil {
			return nil, err
		}
		if constant.SuccessCode != userInfoResp.StatusCode {
			hlog.CtxErrorf(ctx, "get user info failed, req: %+v, resp: %+v", userInfoReq, userInfoResp)
			return nil, constant.UserNotExistErr
		}
		u := userInfoResp.User
		data[i] = &biz.UserInfo{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        u.IsFollow,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
	}
	return data, nil
}

func GetFriendListById(ctx context.Context, userId int64) ([]*biz.FriendUserInfo, error) {
	// 1. 查询db
	followList, err := model.SelectFriendByUserIdA(userId)
	if err != nil {
		return nil, err
	}
	// 2. 查询用户信息
	data := make([]*biz.FriendUserInfo, len(followList))
	for i, follow := range followList {
		userInfoReq := &userBiz.UserInfoReq{UserId: follow.UserIdB}
		userInfoResp, err := internalClient.UserServiceClient.UserInfo(ctx, userInfoReq)
		if err != nil {
			return nil, err
		}
		if constant.SuccessCode != userInfoResp.StatusCode {
			return nil, constant.UserNotExistErr
		}
		u := userInfoResp.User
		data[i] = &biz.FriendUserInfo{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        true,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
	}
	return data, nil
}

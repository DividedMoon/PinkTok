package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"relation_service/biz"
	"relation_service/internal/constant"
	"relation_service/internal/service"
	utils "relation_service/internal/util"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// SendRelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) SendRelationAction(ctx context.Context, req *biz.RelationActionReq) (resp *biz.RelationActionResp, err error) {
	if constant.ActionTypeSubmitFollow == req.ActionType {
		err = service.SubmitFollowRelationAction(ctx, req)
		if err != nil {
			res := utils.BuildBaseResp(err)
			hlog.CtxErrorf(ctx, "SubmitFollowRelationAction err:%v, res:%+v", err, res)
			return &biz.RelationActionResp{
				StatusCode: res.StatusCode,
				StatusMsg:  res.StatusMsg,
			}, err
		}
	} else if constant.ActionTypeCancelFollow == req.ActionType {
		return &biz.RelationActionResp{
			StatusCode: constant.InvalidActionTypeCode,
			StatusMsg:  constant.InvalidActionTypeMsg,
		}, nil
	}
	return &biz.RelationActionResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
	}, nil
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *biz.RelationFollowListReq) (resp *biz.RelationFollowListResp, err error) {
	ids, err := service.GetFollowListById(ctx, req.UserId)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "GetFollowListById err:%v, res:%+v", err, res)
		return &biz.RelationFollowListResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}, err
	}
	resp = &biz.RelationFollowListResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserList:   ids,
	}
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *biz.RelationFollowerListReq) (resp *biz.RelationFollowerListResp, err error) {
	users, err := service.GetFollowerListById(ctx, req.UserId)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "GetFollowerListById err:%v, res:%+v", err, res)
		return &biz.RelationFollowerListResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}, err
	}
	resp = &biz.RelationFollowerListResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserList:   users,
	}
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *biz.RelationFriendListReq) (resp *biz.RelationFriendListResp, err error) {
	friends, err := service.GetFriendListById(ctx, req.UserId)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "GetFriendListById err:%v, res:%+v", err, res)
		return &biz.RelationFriendListResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}, err
	}
	resp = &biz.RelationFriendListResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserList:   friends,
	}
	return
}

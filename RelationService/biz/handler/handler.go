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

	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *biz.RelationFollowerListReq) (resp *biz.RelationFollowerListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *biz.RelationFriendListReq) (resp *biz.RelationFriendListResp, err error) {
	// TODO: Your code here...
	return
}

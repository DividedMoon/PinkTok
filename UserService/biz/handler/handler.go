package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"user_service/biz"
	"user_service/internal/constant"
	"user_service/internal/service"
	utils "user_service/internal/util"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *biz.UserRegisterReq) (resp *biz.UserRegisterResp, err error) {
	hlog.CtxInfof(ctx, "request: %+v", req.Username)
	// 新建用户
	user, err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "register user error: %+v", err)
		resp = &biz.UserRegisterResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
			UserId:     -1,
		}
		return
	}
	resp = &biz.UserRegisterResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     user.Id,
	}
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *biz.UserLoginReq) (resp *biz.UserLoginResp, err error) {
	id, err := service.TryLogin(req.Username, req.Password)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "login error: %+v", err)
		resp = &biz.UserLoginResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
			UserId:     -1,
		}
		return
	}
	resp = &biz.UserLoginResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     id,
	}
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *biz.UserInfoReq) (resp *biz.UserInfoResp, err error) {
	u, err := service.GetUserInfo(req.UserId)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "get user info error: %+v", err)
		resp = &biz.UserInfoResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}
		return
	}
	resp = &biz.UserInfoResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		User:       u,
	}
	return
}

// UserUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserUpdate(ctx context.Context, req *biz.UserUpdateReq) (resp *biz.UserUpdateResp, err error) {
	u, err := service.GetUserInfo(req.UserId)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "get user info error: %+v", err)
		resp = &biz.UserUpdateResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}
		return
	}

	change := make(map[string]int)
	var i int64
	for _, pairValue := range req.Ext {
		i, err = strconv.ParseInt(pairValue.Value, 10, 32)
		if err != nil {
			res := utils.BuildBaseResp(err)
			hlog.CtxErrorf(ctx, "parse int error: %+v", err)
			resp = &biz.UserUpdateResp{
				StatusCode: res.StatusCode,
				StatusMsg:  res.StatusMsg,
			}
			return
		}
		change[pairValue.Key] = int(i)
	}

	time, err := service.UpdateUserInfo(u, change)
	if err != nil {
		res := utils.BuildBaseResp(err)
		hlog.CtxErrorf(ctx, "update user info error: %+v", err)
		resp = &biz.UserUpdateResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}
		return
	}
	return &biz.UserUpdateResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		ModifyTime: time,
	}, nil
}

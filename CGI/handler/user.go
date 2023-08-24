package handler

import (
	internalClient "cgi/internal/client"
	"cgi/internal/constant"
	"cgi/internal/utils"
	"cgi/middleware"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"user_service/biz"
)

type checkUserInfoResp struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	User       biz.UserInfo `json:"user"`
}

// CheckUserInfoHandler get the login user info
func CheckUserInfoHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err    error
		userId int64
		resp   *utils.BaseResp
	)
	// 解析请求参数
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, checkUserInfoResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "get current user id failed, something wrong in authorize",
		})
		return
	}
	userId = v.(int64)
	// 获取用户信息
	var (
		userInfoReq = biz.UserInfoReq{
			UserId: userId,
		}
	)

	info, err := internalClient.UserServiceClient.UserInfo(ctx, &userInfoReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get user info failed, err: %v", err)
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusInternalServerError, checkUserInfoResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	userInfo := &biz.UserInfo{
		Id:              info.User.Id,
		Name:            info.User.Name,
		FollowCount:     info.User.FollowCount,
		FollowerCount:   info.User.FollowerCount,
		IsFollow:        false,
		Avatar:          info.User.Avatar,
		BackgroundImage: info.User.BackgroundImage,
		Signature:       info.User.Signature,
		TotalFavorited:  info.User.TotalFavorited,
		WorkCount:       info.User.WorkCount,
		FavoriteCount:   info.User.FavoriteCount,
	}
	c.JSON(consts.StatusOK, checkUserInfoResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		User:       *userInfo,
	})
}

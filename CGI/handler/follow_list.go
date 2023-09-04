package handler

import (
	internalClient "cgi/internal/client"
	"cgi/internal/constant"
	"cgi/internal/utils"
	"cgi/middleware"
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"relation_service/biz"
)

type followListReq struct {
	UserId string `query:"user_id"`
	Token  string `query:"token"`
}

type followListResp struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
	UserList   []biz.UserInfo `json:"user_list"`
}

func FollowListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  followListReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, followListResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, followListResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, followListResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/follow/list with req: %+v from: %+v", req, c.ClientIP())

	// 调用relation service
	followReq := &biz.RelationFollowListReq{
		UserId: userId,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	followResp, err := internalClient.RelationServiceClient.GetFollowList(ctx, followReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get follow list failed, err: %+v", err)
		c.JSON(consts.StatusInternalServerError, followListResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get follow list success, resp: %+v", followResp)
	userList := make([]biz.UserInfo, 0)
	for _, p := range followResp.UserList {
		if p != nil {
			userList = append(userList, biz.UserInfo{
				Id:              p.Id,
				Name:            p.Name,
				FollowCount:     p.FollowCount,
				FollowerCount:   p.FollowerCount,
				IsFollow:        p.IsFollow,
				Avatar:          p.Avatar,
				BackgroundImage: p.BackgroundImage,
				Signature:       p.Signature,
				TotalFavorited:  p.TotalFavorited,
				WorkCount:       p.WorkCount,
				FavoriteCount:   p.FavoriteCount,
			})
		}
	}
	c.JSON(consts.StatusOK, followListResp{
		StatusCode: followResp.StatusCode,
		StatusMsg:  followResp.StatusMsg,
		UserList:   userList,
	})
}

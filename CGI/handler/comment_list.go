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
	"interact_service/biz"
)

type commentListReq struct {
	VideoId int64  `query:"video_id"`
	Token   string `query:"token"`
}

type commentListResp struct {
	StatusCode  int32         `json:"status_code"`
	StatusMsg   string        `json:"status_msg"`
	CommentList []biz.Comment `json:"comment_list"`
}

func CommentListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  commentListReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, commentListResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	userId, ok := v.(int64)
	if !ok {
		hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
		c.JSON(consts.StatusInternalServerError, commentListResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, commentListResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "relation/follower/list with req: %+v from: %+v", req, c.ClientIP())

	// 调用interact service
	commentReq := &biz.CommentListReq{
		UserId:  userId,
		VideoId: req.VideoId,
	}
	ctx = metainfo.WithBackwardValues(ctx)
	commentResp, err := internalClient.InteractServiceClient.CommentList(ctx, commentReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "get follower list failed, err: %+v", err)
		c.JSON(consts.StatusInternalServerError, commentListResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "get follower list success, resp: %+v", commentResp)
	commentList := make([]biz.Comment, 0)
	for _, p := range commentResp.CommentList {
		if p != nil {
			commentList = append(commentList, biz.Comment{
				Id:         p.Id,
				User:       p.User,
				Content:    p.Content,
				CreateDate: p.CreateDate,
			})
		}
	}
	c.JSON(consts.StatusOK, commentListResp{
		StatusCode:  commentResp.StatusCode,
		StatusMsg:   commentResp.StatusMsg,
		CommentList: commentList,
	})
}

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

type commentReq struct {
	Token       string `query:"token"`
	VideoId     int64  `query:"video_id"`
	ActionType  int32  `query:"action_type"`
	CommentText string `query:"comment_text"`
	CommentId   int64  `query:"comment_id"`
}

type commentResp struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    biz.Comment `json:"comment"`
}

func CommentActionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  commentReq
		resp *utils.BaseResp
	)
	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, commentResp{
			StatusCode: constant.AuthorizationFailedErrCode,
			StatusMsg:  "登录超时，请重新登录",
		})
		return
	}
	// 参数转换
	var userId int64
	userIdFloat, ok := v.(float64)
	if ok {
		userId = int64(userIdFloat)
	} else {
		userId, ok = v.(int64)
		if !ok {
			hlog.CtxErrorf(ctx, "userId cannot be parsed to int64")
			c.JSON(consts.StatusInternalServerError, commentResp{
				StatusCode: constant.ParamErrCode,
				StatusMsg:  constant.ParamErrMsg,
			})
			return
		}
	}

	// 解析请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusInternalServerError, commentResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "comment/action with req: %+v from: %+v", req, c.ClientIP())

	actionReq := &biz.CommentActionReq{
		UserId:      userId,
		ActionType:  req.ActionType,
		VideoId:     req.VideoId,
		CommentId:   req.CommentId,
		CommentText: req.CommentText,
	}
	hlog.CtxInfof(ctx, "CommentAction with req: %+v", actionReq)
	ctx = metainfo.WithBackwardValues(ctx)
	actionResp, err := internalClient.InteractServiceClient.CommentAction(ctx, actionReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "comment/action failed with err: %+v", err)
		c.JSON(consts.StatusInternalServerError, commentResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  constant.ServerErrMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "comment/action with resp: %+v", actionResp)
	c.JSON(consts.StatusOK, commentResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  actionResp.StatusMsg,
		Comment: biz.Comment{
			Id:         actionResp.Comment.Id,
			User:       actionResp.Comment.User,
			Content:    actionResp.Comment.Content,
			CreateDate: actionResp.Comment.CreateDate,
		},
	})
}

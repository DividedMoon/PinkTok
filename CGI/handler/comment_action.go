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
	"strconv"
)

type commentActionReq struct {
	ActionType  string  `query:"action_type"`  // 1-发布评论，2-删除评论
	CommentID   *string `query:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
	CommentText *string `query:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	Token       string  `query:"token"`        // 用户鉴权token
	VideoID     string  `query:"video_id"`     // 视频id
}

type commentActionResp struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    biz.Comment `json:"comment"`
}

func CommentActionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  commentActionReq
		resp *utils.BaseResp
	)

	// 判断jwt状态
	v, exists := c.Get(middleware.CurrentUserIdKey)
	if !exists {
		hlog.CtxErrorf(ctx, "get current user id failed, something wrong in authorize")
		c.JSON(consts.StatusUnauthorized, commentActionResp{
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
			c.JSON(consts.StatusInternalServerError, commentActionResp{
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
		c.JSON(consts.StatusInternalServerError, commentActionResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.CtxInfof(ctx, "comment/action with req: %+v from: %+v", req, c.ClientIP())

	// 参数校验
	actionType, err := strconv.ParseInt(req.ActionType, 10, 32)
	if err != nil || (actionType != 1 && actionType != 2) {
		c.JSON(consts.StatusInternalServerError, commentActionResp{
			StatusCode: constant.ParamErrCode,
			StatusMsg:  constant.ParamErrMsg,
		})
		return
	}
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusInternalServerError, commentActionResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	// 业务逻辑
	var actionReq *biz.CommentActionReq
	if actionType == 1 {
		actionReq = &biz.CommentActionReq{
			ActionType:  1,
			CommentText: *req.CommentText,
			UserId:      userId,
			VideoId:     videoId,
		}
	} else {
		commentId, err := strconv.ParseInt(*req.CommentID, 10, 64)
		if err != nil {
			resp = utils.BuildBaseResp(err)
			c.JSON(consts.StatusInternalServerError, commentActionResp{
				StatusCode: resp.StatusCode,
				StatusMsg:  resp.StatusMsg,
			})
			return
		}
		actionReq = &biz.CommentActionReq{
			ActionType: 2,
			CommentId:  commentId,
			UserId:     userId,
			VideoId:    videoId,
		}
	}
	hlog.CtxInfof(ctx, "comment/action with actionReq: %+v", actionReq)
	ctx = metainfo.WithBackwardValues(ctx)
	actionResp, err := internalClient.InteractServiceClient.CommentAction(ctx, actionReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "comment/action failed with err: %+v", err)
		resp = utils.BuildBaseResp(err)
		c.JSON(consts.StatusInternalServerError, commentActionResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}
	hlog.CtxInfof(ctx, "comment/action with resp: %+v", actionResp)
	c.JSON(consts.StatusOK, relationActionResp{
		StatusCode: actionResp.StatusCode,
		StatusMsg:  actionResp.StatusMsg,
	})

}

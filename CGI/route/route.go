package route

import (
	"cgi/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterGroupRoute(h *server.Hertz) {
	tourist := h.Group("/douyin")
	tourist.GET("/feed/", handler.GetFeedHandler)
	tourist.POST("/user/register/", handler.RegisterHandler)
	tourist.POST("/user/login/", handler.LoginHandler)

	// 以下接口需要登录
	user := h.Group("/douyin")
	user.Use()
	user.GET("/user/", handler.CheckUserInfoHandler)
	user.POST("/publish/action/", handler.PublishActionHandler)
	user.GET("/publish/list/", handler.PublicListHandler)
	// 互动接口
	user.POST("/favorite/action/", handler.FavoriteActionHandler)
	user.GET("/favorite/list/", handler.FavoriteListHandler)
	user.POST("/comment/action/", handler.CommentActionHandler)
	user.GET("/comment/list/", handler.CommentListHandler)
	// 社交接口
	user.POST("/follow/action/", handler.RelationActionHandler)
	user.GET("/relation/follow/list/", handler.FollowListHandler)
	user.GET("/relation/follower/list/", handler.FollowerListHandler)
	user.GET("/relation/friend/list/", handler.FriendListHandler)
}

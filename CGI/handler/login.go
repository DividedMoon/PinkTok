package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"time"
)

type LoginReq struct {
	Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
}

func LoginHandler(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, utils.H{
		"code":    code,
		"token":   token,
		"expire":  expire.Format(time.RFC3339),
		"message": "success",
	})
}

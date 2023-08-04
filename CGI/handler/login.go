package handler

import (
	"cgi/internal/constant"
	"client/dto"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	_ "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
)

func LoginHandler(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get(jwt.IdentityKey)
	userId := v.(int64)
	token := c.GetString("token")
	c.JSON(http.StatusOK, dto.UserLoginResp{
		StatusCode: constant.SuccessCode,
		StatusMsg:  constant.SuccessMsg,
		UserId:     userId,
		Token:      token,
	})
}

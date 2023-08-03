package middleware

import (
	"cgi/handler"
	"client/dto"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "byllhj"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Duration(2) * time.Minute,
		MaxRefresh:    time.Minute,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		LoginResponse: handler.LoginHandler,
		Authenticator: authenticate,
		IdentityKey:   IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &dto.UserInfo{
				Name: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*dto.UserInfo); ok {
				return jwt.MapClaims{
					IdentityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

func authenticate(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var loginStruct handler.LoginReq
	if err := c.BindAndValidate(&loginStruct); err != nil {
		return nil, err
	}

	return "users[0]", nil
}

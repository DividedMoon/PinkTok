package middleware

import (
	"cgi/internal/constant"
	utils2 "cgi/internal/utils"
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
	identityKey   = "user_id"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "PinkTok",
		Key:     []byte("secret key"),
		Timeout: 2 * time.Minute,
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, time time.Time) {
			hlog.CtxInfof(ctx, "Get login response = %+v, token is issued by: %+v", token, c.ClientIP())
			c.Set("token", token)
		},
		Authenticator: authenticator,
		Authorizator:  authorizator,
		IdentityKey:   identityKey,
		// 设置 token 中的 payload
		PayloadFunc: payloadFunc,
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "Get http status message = %+v", e.Error())
			return utils2.BuildBaseResp(e).StatusMsg
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    constant.AuthorizationFailedErrCode,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(int64); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

// authenticator verifies password at login
func authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var loginStruct dto.UserLoginReq
	if err := c.BindAndValidate(&loginStruct); err != nil {
		return nil, err
	}
	//TODO 查User
	c.Set(identityKey, 2)
	return 2, nil
}

// authorizator verifies the token at each request
func authorizator(data interface{}, ctx context.Context, c *app.RequestContext) bool {
	if v, ok := data.(int64); ok {
		currentUserId := v
		c.Set("current_user_id", currentUserId)
		hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
		return true
	}
	return false
}

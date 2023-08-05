package middleware

import (
	"cgi/internal/constant"
	utils2 "cgi/internal/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

type userLoginReq struct {
	Username string `query:"username"`
	Password string `query:"password"`
}

type userLoginResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

const (
	UserIdKey = "user_id"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = UserIdKey
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "PinkTok",
		Key:     []byte("secret key"),
		Timeout: 2 * time.Minute,
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, time time.Time) {
			hlog.CtxInfof(ctx, "Get login response = %+v, token is issued by: %+v", token, c.ClientIP())
			v, _ := c.Get(UserIdKey)
			userId := v.(int)
			c.JSON(http.StatusOK, userLoginResp{
				StatusCode: constant.SuccessCode,
				StatusMsg:  constant.SuccessMsg,
				UserId:     userId,
				Token:      token,
			})
		},
		Authenticator: authenticator,
		Authorizator:  authorizator,
		IdentityKey:   IdentityKey,
		// 设置 token 中的 payload
		PayloadFunc: payloadFunc,
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxInfof(ctx, "Get http status message = %+v", e.Error())
			return utils2.BuildBaseResp(e).StatusMsg
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": constant.AuthorizationFailedErrCode,
				"status_msg":  message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(int); ok {
		return jwt.MapClaims{
			IdentityKey: v,
		}
	}
	return jwt.MapClaims{}
}

// authenticator verifies password at login
func authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	hlog.CtxInfof(ctx, "Get authenticator clientIP: "+c.ClientIP())
	var loginStruct userLoginReq
	if err := c.BindAndValidate(&loginStruct); err != nil {
		return nil, err
	}
	if loginStruct.Password != loginStruct.Username {
		return nil, constant.UserIsNotExistErr
	}
	//TODO 查User
	c.Set(IdentityKey, 2)
	return 2, nil
}

// authorizator verifies the token at each request
func authorizator(data interface{}, ctx context.Context, c *app.RequestContext) bool {
	hlog.CtxInfof(ctx, "Get authorizator clientIP: "+c.ClientIP())
	if v, ok := data.(int64); ok {
		currentUserId := v
		c.Set("current_user_id", currentUserId)
		hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
		return true
	}
	return false
}

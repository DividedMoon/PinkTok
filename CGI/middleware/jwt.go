package middleware

import (
	"cgi/internal/constant"
	utils2 "cgi/internal/utils"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"strconv"
	"time"
)

type userLoginReq struct {
	Username string `query:"username"`
	Password string `query:"password"`
}

type UserLoginResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

const (
	UserIdKey        = "user_id"
	CurrentUserIdKey = "current_user_id"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = UserIdKey
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "PinkTok",
		Key:         []byte("secret key"),
		Timeout:     24 * time.Hour,
		TokenLookup: "query:token",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, time time.Time) {
			hlog.CtxInfof(ctx, "Get login response = %+v, token is issued by: %+v", token, c.ClientIP())
			v, _ := c.Get(UserIdKey)
			userId := v.(int)
			c.JSON(http.StatusOK, UserLoginResp{
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
			resp := utils2.BuildBaseResp(e)
			respBytes, _ := json.Marshal(resp)
			return string(respBytes)
		},
		Unauthorized: unauthorized,
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
	if v, ok := data.(int64); ok {
		return jwt.MapClaims{
			IdentityKey: v,
		}
	}
	if v, ok := data.(int32); ok {
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
	loginResp, err := LoginHandler(ctx, loginStruct)
	if err != nil {
		hlog.CtxErrorf(ctx, "Login in authenticator error: %+v", err)
		return nil, err
	}
	if loginResp.StatusCode != constant.SuccessCode {
		hlog.CtxInfof(ctx, "Get authenticator failed: %+v", loginResp.StatusMsg)
		return nil, constant.ErrNo{
			ErrCode: loginResp.StatusCode,
			ErrMsg:  loginResp.StatusMsg,
		}
	}
	c.Set(IdentityKey, loginResp.UserId)
	return loginResp.UserId, nil
}

// authorizator verifies the token at each request
func authorizator(data interface{}, ctx context.Context, c *app.RequestContext) bool {
	hlog.CtxInfof(ctx, "Get authorizator clientIP: "+c.ClientIP())
	userIdString, ok := c.GetQuery(IdentityKey)
	if !ok {
		hlog.CtxErrorf(ctx, "Can not get user_id in query")
		return false
	}
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, "Parse userId error: %+v", err)
		return false
	}
	if v, ok := data.(float64); ok {
		currentUserId := int64(v)
		if userId == currentUserId {
			c.Set(CurrentUserIdKey, currentUserId)
			hlog.CtxInfof(ctx, "Token is verified by clientIP %s userId %s", c.ClientIP(), userId)
			return true

		}
	}
	return false
}

func unauthorized(ctx context.Context, c *app.RequestContext, code int, message string) {
	hlog.CtxInfof(ctx, "Get unauthorized clientIP: %s and message: %s", c.ClientIP(), message)
	resp := &utils2.BaseResp{}
	err := json.Unmarshal([]byte(message), resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "Unmarshal error when unauthorized: %+v", err)
		c.JSON(code, utils2.BaseResp{
			StatusCode: constant.ServiceErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(code, utils2.BaseResp{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}

// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	client "PinkTok/VideoService/biz/router/client"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	client.Register(r)
}

package main

import (
	"cgi/middleware"
	"cgi/route"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
	// 初始化jwt
	middleware.InitJwt()
	// 注册路由
	route.RegisterGroupRoute(h)
	h.Spin()
}

package main

import (
	"cgi/internal/client"
	"cgi/internal/config"
	"cgi/middleware"
	"cgi/route"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func main() {

	h := server.Default(server.WithHostPorts("0.0.0.0:11010"))
	if err := config.InitConfigs(); err != nil {
		hlog.Fatal(err.Error())
		return
	}

	// 初始化客户端
	client.InitClient()
	// 初始化jwt
	middleware.InitJwt()
	// 注册路由
	route.RegisterGroupRoute(h)
	h.Spin()
}

package main

import (
	"cgi/internal"
	"cgi/middleware"
	"cgi/route"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
	if err := internal.InitConfigs(); err != nil {
		log.Fatal(err.Error())
		return
	}
	// 初始化jwt
	middleware.InitJwt()
	// 注册路由
	route.RegisterGroupRoute(h)
	h.Spin()
}

package main

import (
	"cgi/internal/client"
	"cgi/internal/config"
	"cgi/middleware"
	"cgi/route"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	client.InitClient()

	h := server.Default(server.WithHostPorts("0.0.0.0:9000"))
	if err := config.InitConfigs(); err != nil {
		log.Fatal(err.Error())
		return
	}
	// 初始化jwt
	middleware.InitJwt()
	// 注册路由
	route.RegisterGroupRoute(h)
	h.Spin()
}

// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"user_service/biz/model"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8889"))

	register(h)
	model.InitDB()

	h.Spin()
}

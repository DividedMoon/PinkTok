package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/server"
	"interact_service/biz/handler"
	biz "interact_service/biz/interactservice"
	"interact_service/internal/model"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:11013")
	if err != nil {
		hlog.Error(err)
	}
	model.InitDB()
	svr := biz.NewServer(new(handler.InteractServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

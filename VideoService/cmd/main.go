package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/server"
	"net"
	"os"
	"video_service/biz/handler"
	biz "video_service/biz/videoservice"
	"video_service/internal/client"
	"video_service/internal/config"
	"video_service/internal/dal/db"
	"video_service/internal/dal/redis"
	"video_service/internal/middleware/minio"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8890")
	if err != nil {
		hlog.Error(err.Error())
	}
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	client.InitClient()

	svr := biz.NewServer(new(handler.VideoServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()

	if err != nil {
		hlog.Error(err.Error())
	}

}

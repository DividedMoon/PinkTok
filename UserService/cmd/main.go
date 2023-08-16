package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user_service/biz/handler"
	biz "user_service/biz/userservice"
	"user_service/internal/model"
)

func main() {
	// 创建一个通道来接收中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:11011")
	if err != nil {
		hlog.Errorf("resolve tcp addr failed, err:%v", err)
	}
	model.InitDB()

	servername := "user_service"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(servername),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithInsecure())
	defer p.Shutdown(context.Background())
	svr := biz.NewServer(new(handler.UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: servername,
		}),
	)

	go func() {
		err = svr.Run()
		if err != nil {
			hlog.Errorf(err.Error())
		}
	}()

	<-interrupt
	hlog.Info("接收到中断信号,关闭服务器中")
	err = svr.Stop()
	if err != nil {
		hlog.Errorf(err.Error())
	}
	hlog.Info("服务器已关闭")
}

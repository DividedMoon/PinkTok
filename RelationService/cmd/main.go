package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"net"
	"relation_service/biz/handler"
	biz "relation_service/biz/relationservice"
	internalClient "relation_service/internal/client"
	"relation_service/internal/model"
	"relation_service/internal/mw/redis"
	"sync"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:11012")
	if err != nil {
		hlog.Errorf("resolve tcp addr failed, err:%v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		redis.InitRedis()
	}()
	go func() {
		defer wg.Done()
		model.InitDB()
	}()
	go func() {
		defer wg.Done()
		internalClient.InitClient()
	}()
	wg.Wait()
	servername := "relation_service"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(servername),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithInsecure())
	defer p.Shutdown(context.Background())
	svr := biz.NewServer(new(handler.RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: servername,
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

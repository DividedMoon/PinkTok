package client

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"interact_service/biz/interactservice"
	"time"
	"user_service/biz/userservice"
)

var (
	UserServiceClient     userservice.Client
	InteractServiceClient interactservice.Client
	err                   error
)

func InitClient() {
	pUserService := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("user_service"),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithEnableTracing(true),
		provider.WithEnableMetrics(false),
		provider.WithInsecure(),
	)
	defer pUserService.Shutdown(context.Background())
	UserServiceClient, err =
		userservice.NewClient("user_service",
			client.WithHostPorts("127.0.0.1:11011"),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "user_service",
			}),
			client.WithRPCTimeout(5*time.Duration(time.Second)),
		)
	if err != nil {
		hlog.Errorf("UserServiceClient init failed: %+v", err)
	}

	pInteractService := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("interact_service"),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithEnableTracing(true),
		provider.WithEnableMetrics(false),
		provider.WithInsecure(),
	)
	defer pInteractService.Shutdown(context.Background())

	InteractServiceClient, err =
		interactservice.NewClient("interact_service",
			client.WithHostPorts("127.0.0.1:11013"),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "interact_service",
			}),
			client.WithRPCTimeout(5*time.Duration(time.Second)),
		)
	if err != nil {
		hlog.Errorf("InteractServiceClient init failed: %+v", err)
	}
}

package client

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"user_service/biz/userservice"
)

var (
	UserServiceClient userservice.Client
	err               error
)

func InitClient() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("user_service_client"),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithEnableTracing(true),
		provider.WithEnableMetrics(false),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	UserServiceClient, err =
		userservice.NewClient("user_service_client",
			client.WithHostPorts("127.0.0.1:11011"),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "user_service_client",
			}))
	if err != nil {
		hlog.Errorf("UserServiceClient init failed: %+v", err)
	}
}

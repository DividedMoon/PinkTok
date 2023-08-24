package client

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"relation_service/biz/relationservice"
	"user_service/biz/userservice"
)

var (
	UserServiceClient     userservice.Client
	RelationServiceClient relationservice.Client
	err                   error
)

func InitClient() {
	p1 := setupTracing("user_service_client")
	defer p1.Shutdown(context.Background())
	p2 := setupTracing("relation_service_client")
	defer p2.Shutdown(context.Background())
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
	RelationServiceClient, err =
		relationservice.NewClient("relation_service_client",
			client.WithHostPorts("127.0.0.1:11012"),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "relation_service_client",
			}))
	if err != nil {
		hlog.Errorf("RelationServiceClient init failed: %+v", err)
	}
}

func setupTracing(serviceName string) provider.OtelProvider {
	return provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithEnableTracing(true),
		provider.WithEnableMetrics(false),
		provider.WithInsecure(),
	)
}

package client

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"interact_service/biz/interactservice"
	"middleware/auth"
	"middleware/msgno"
	"relation_service/biz/relationservice"
	"time"
	"user_service/biz/userservice"
)

const (
	Remote = "127.0.0.1"
)

var (
	UserServiceClient     userservice.Client
	RelationServiceClient relationservice.Client
	InteractServiceClient interactservice.Client
	err                   error
)

func InitClient() {
	p1 := setupTracing("user_service_client")
	defer p1.Shutdown(context.Background())
	p2 := setupTracing("relation_service_client")
	defer p2.Shutdown(context.Background())
	p3 := setupTracing("interact_service_client")
	defer p3.Shutdown(context.Background())
	UserServiceClient =
		userservice.MustNewClient("user_service_client",
			client.WithHostPorts(fmt.Sprintf("%s%s", Remote, ":11011")),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithMiddleware(msgno.MsgNoMiddleware),
			client.WithMiddleware(auth.AuthenticateClient),
			client.WithConnectTimeout(time.Second*2),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "user_service_client",
			}))
	if err != nil {
		hlog.Errorf("UserServiceClient init failed: %+v", err)
	}
	RelationServiceClient =
		relationservice.MustNewClient("relation_service_client",
			client.WithHostPorts(fmt.Sprintf("%s%s", Remote, ":11012")),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithMiddleware(msgno.MsgNoMiddleware),
			client.WithMiddleware(auth.AuthenticateClient),
			client.WithConnectTimeout(time.Second*2),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "relation_service_client",
			}))
	if err != nil {
		hlog.Errorf("RelationServiceClient init failed: %+v", err)
	}
	InteractServiceClient =
		interactservice.MustNewClient("interact_service_client",
			client.WithHostPorts(fmt.Sprintf("%s%s", Remote, ":11013")),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithMiddleware(msgno.MsgNoMiddleware),
			client.WithMiddleware(auth.AuthenticateClient),
			client.WithConnectTimeout(time.Second*2),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "interact_service_client",
			}))
	if err != nil {
		hlog.Errorf("InteractServiceClient init failed: %+v", err)
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

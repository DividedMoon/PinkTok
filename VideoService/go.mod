module video_service

go 1.20

require (
	github.com/cloudwego/fastpb v0.0.4
	github.com/cloudwego/hertz v0.6.7
	github.com/cloudwego/kitex v0.6.2
	github.com/go-redis/redis/v7 v7.4.1
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8
	github.com/hertz-contrib/logger/zap v1.0.0
	github.com/kitex-contrib/obs-opentelemetry v0.2.3
	github.com/minio/minio-go/v7 v7.0.61
	github.com/u2takey/ffmpeg-go v0.5.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.3
	user_service v0.0.0
	interact_service v0.0.0
)

require (
	github.com/apache/thrift v0.13.0 // indirect
	github.com/aws/aws-sdk-go v1.38.20 // indirect
	github.com/bytedance/gopkg v0.0.0-20230531144706-a12972768317 // indirect
	github.com/bytedance/sonic v1.8.8 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/chenzhuoyu/iasm v0.0.0-20230222070914-0b1b64b0e762 // indirect
	github.com/choleraehyq/pid v0.0.16 // indirect
	github.com/cloudwego/configmanager v0.2.0 // indirect
	github.com/cloudwego/frugal v0.1.6 // indirect
	github.com/cloudwego/netpoll v0.4.0 // indirect
	github.com/cloudwego/thriftgo v0.2.11 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20230509042627-b1315fad0c5a // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/tidwall/gjson v1.13.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/u2takey/go-utils v0.3.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.42.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.17.0 // indirect
	go.opentelemetry.io/contrib/propagators/ot v1.17.0 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/grpc v1.55.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace user_service => ../UserService
replace interact_service => ../InteractService
replace middleware => ../Middleware


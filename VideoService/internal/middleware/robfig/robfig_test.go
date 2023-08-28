package robfig

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"testing"
	"time"
	"video_service/internal/client"
	"video_service/internal/config"
	"video_service/internal/dal/db"
	"video_service/internal/dal/redis"
	"video_service/internal/middleware/minio"
)

func TestRobfig(t *testing.T) {
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	client.InitClient()
	hlog.Info("init")
	Init()
	time.Sleep(time.Second * 70)
}

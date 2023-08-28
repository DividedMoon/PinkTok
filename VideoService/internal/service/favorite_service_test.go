package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"testing"
	"time"
	"video_service/internal/client"
	"video_service/internal/config"
	"video_service/internal/dal/db"
	"video_service/internal/dal/redis"
	"video_service/internal/middleware/minio"
	"video_service/internal/middleware/robfig"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func setup() {
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	client.InitClient()
	robfig.Init()
}

func teardown() {
	redis.CloseRedis()
}

// TestVideoServiceUnFavoriteAction 取消赞
func TestFavoriteServiceUnFavoriteAction(t *testing.T) {

	service := NewVideoService(context.Background())
	err := service.FavoriteAction(10, 10, 0)
	if err != nil {
		t.Error(err)
	}
}

// TestVideoServiceFavoriteAction 赞
func TestFavoriteServiceFavoriteAction(t *testing.T) {

	service := NewVideoService(context.Background())
	err := service.FavoriteAction(10, 10, 1)
	if err != nil {
		t.Error(err)
	}
}

// TestVideoServiceFavoriteActionUsingCache Done
func TestFavoriteServiceFavoriteActionUsingCache(t *testing.T) {
	/*
		流程是：
		1.首先打印video的初始值，查看video的favorite_count是多少
		2.大于1000 然后会将其加一然后结果存到redis中并不会落盘 这时候第二次查询数据库结果还是相同的
		3.
	*/
	service := NewVideoService(context.Background())
	v, err := db.GetVideoByID(10)
	if err != nil {
		t.Error(err)
	}
	hlog.Info("1.首先打印video的初始值，查看video的favorite_count是多少")
	hlog.Info("v", v)

	hlog.Infof("加一操作，第一次初始化并进入redis")
	err = service.FavoriteAction(10, 10, 1)
	if err != nil {
		t.Error(err)
	}

	hlog.Info("加一操作，第二次进入redis")
	err = service.FavoriteAction(10, 11, 1)

	hlog.Infof("手动查询一次 查看数据库中的值是多少 这时候应该差1 因为第一次的时候既更新了redis也更新了数据库")
	video, err := db.GetVideoByID(10)
	if err != nil {
		t.Error(err)
	}
	hlog.Info("video", video)

	hlog.Infof("再次更新两次")
	time.Sleep(40 * time.Second)
	_ = service.FavoriteAction(10, 12, 1)
	_ = service.FavoriteAction(10, 13, 1)

	hlog.Info("这时候值应该为第一次更新加2之后的值因为上一次的更新已经落盘了")
	video, err = db.GetVideoByID(10)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(60 * time.Second)
	hlog.Info("video", video)
}

// TestVideoService_GetFavoriteVideoList Done
func TestVideoService_GetFavoriteVideoList(t *testing.T) {
	service := NewVideoService(context.Background())
	videos, err := service.GetFavoriteVideoList(1)
	if err != nil {
		t.Error(err)
	}
	hlog.Info("videos", videos)
}

package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io/ioutil"
	"os"
	"testing"
	"video_service/biz"
	"video_service/internal/config"
	"video_service/internal/dal/db"
	"video_service/internal/dal/redis"
	"video_service/internal/middleware/ffmpeg"
	"video_service/internal/middleware/minio"
)

// Done! 需要小视频 1M左右即可
func TestPublishAction(t *testing.T) {

	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()

	videoPath := "/tmp/test-videos/pexels-danilo-arenas-16263007 (240p).mp4" // 替换为实际视频文件的路径
	videoFile, err := os.Open(videoPath)
	if err != nil {
		hlog.Error("Error opening video file", err)
		return
	}
	defer videoFile.Close()

	videoBytes, err := ioutil.ReadAll(videoFile)
	if err != nil {
		hlog.Error("Error reading video file", err)
		return
	}

	testReq := &biz.PublishReq{
		UserId: 1,
		Title:  "test",
		Data:   videoBytes,
	}

	err = NewPublishService(context.TODO()).PublishAction(testReq)

}

// Done!
func TestFFMPEG(t *testing.T) {
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	buf, err := ffmpeg.GetSnapshot("http://106.54.208.133:18000/videobucket/qq.mp4")
	if err != nil {
		hlog.Error(err.Error())
	}
	hlog.Infof("buf: %v", buf)
}

func TestMinioConnectGet(t *testing.T) {
	os.Setenv("http_proxy", "")
	minio.Init()
	url, err := minio.GetObjURL(context.TODO(), "videobucket", "qq.mp4")
	if err != nil {
		hlog.Error(err.Error())
	}
	hlog.Infof("url: %s", url.String())
}

// Done!
func TestGetPublishList(t *testing.T) {
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	req := &biz.GetPublishListReq{
		UserId: 1,
	}
	resp, err := NewPublishService(context.Background()).GetPublishList(req)
	if err != nil {
		hlog.Error(err.Error())
	}
	hlog.Infof("resp: %v", resp)
}

package robfig

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/henrylee2cn/goutil/calendar/cron"
	"time"
	"video_service/internal/dal/db"
	rd "video_service/internal/dal/redis"
)

var jobErrChan chan error

func Init() {
	c := cron.New()
	jobErrChan = make(chan error, 5)
	// 添加定时任务
	hlog.Infof("cron start at %s", time.Now().Format("2006-01-02 15:04:05"))
	err := c.AddFunc("@every 30s", FlushChangedVideo2DB)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}

	// 启动定时器
	hlog.Infof("cron start at %s", time.Now().Format("2006-01-02 15:04:05"))
	c.Start()

	// 监听jobErrChan 一旦有错误发生，打印错误日志 并停止定时任务
	go func() {
		for {
			select {
			case err := <-jobErrChan:
				hlog.Error(err)
				c.Stop()
			}
		}
	}()
}

func FlushChangedVideo2DB() {
	hlog.Infof("FlushChangedVideo2DB start at %s", time.Now().Format("2006-01-02 15:04:05"))
	videoDBs, err := GetChangedVideos()
	if err != nil {
		hlog.Error("GetChangedVideos error", err.Error())
		return
	}
	for i := 0; i < len(videoDBs); i++ {
		err = db.UpdateVideo(videoDBs[i])
		if err != nil {

			hlog.Error("UpdateVideo error", err.Error())
			jobErrChan <- err
			return
		}
	}
	hlog.Infof("FlushChangedVideo2DB end at %s", time.Now().Format("2006-01-02 15:04:05"))
}

func GetChangedVideos() ([]*db.VideoDBInfo, error) {
	//1.获取所有发生变化的视频id
	videoIds, err := rd.GetChangedVideoIds()

	if err != nil {
		hlog.Error("GetChangedVideoIds error", err.Error())
		return nil, err
	}
	//2.逐个从redis里提取视频信息
	videoDBInfos := make([]*db.VideoDBInfo, 0, len(videoIds))
	for _, id := range videoIds {

		videoDBInfo, err := rd.GetVideoHash(int64(id))

		if err != nil {
			hlog.Error("GetVideoHash error", err.Error())
			return nil, err
		}
		videoDBInfos = append(videoDBInfos, videoDBInfo)
	}
	//3.返回视频信息

	return videoDBInfos, nil

}

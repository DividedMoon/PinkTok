package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"sync"
	"testing"
	"time"
	"video_service/biz"
	"video_service/internal/config"
	"video_service/internal/dal/db"
	"video_service/internal/dal/redis"
	"video_service/internal/middleware/minio"
)

func TestErrProcess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errChan := make(chan error, 1)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)

	go worker(ctx, "Worker 1", &wg)
	go worker(ctx, "Worker 2", &wg)
	go worker(ctx, "Worker 3", &wg)
	go func() {
		for {
			select {
			case err := <-errChan:
				fmt.Printf("Error: %s\n", err.Error())
				cancel()
			case <-done:
				println("Error detect Done")
				return
			}
		}
	}()
	// 模拟某个条件满足后取消所有协程
	time.Sleep(2 * time.Second)
	errChan <- fmt.Errorf("Something wrong")

	wg.Wait()
	done <- struct{}{}
	fmt.Println("All workers done")
	time.Sleep(1 * time.Second)
}
func worker(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s cancelled\n", name)
			return
		default:
			fmt.Printf("%s is working\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Done!
func TestGetFeed(t *testing.T) {
	os.Setenv("http_proxy", "")
	db.Init()
	redis.InitRedis()
	_ = config.InitConfigs()
	minio.Init()
	hlog.Infof("Now: %v", time.Now())
	req := &biz.FeedReq{
		UserId:     1,
		LatestTime: time.Now().Add(-24 * time.Hour).Unix(),
	}
	hlog.Infof("GetFeed req: %+v", req)
	resp, err := NewFeedService(context.Background()).GetFeed(req)
	if err != nil {
		t.Errorf("GetFeed error: %v", err)
	}
	hlog.Infof("GetFeed resp: %+v", resp)
}

func TestTime(t *testing.T) {
	fmt.Printf("Now: %v\n", time.Unix(1692430619, 0))
}

package robfig

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/henrylee2cn/goutil/calendar/cron"
	"time"
)

var jobErrChan chan error

func Init() {
	c := cron.New()
	jobErrChan = make(chan error, 5)
	// 添加定时任务
	err := c.AddFunc("@every 30m", ProcessTemplate)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}

	// 启动定时器
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

func ProcessTemplate() {
	// do something
	time.Sleep(30 * time.Second)
	jobErrChan <- fmt.Errorf("test error")

}

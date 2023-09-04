package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
	//userModel "user_service/biz/model/client"
	client "video_service/biz"
	//userClient "video_service/biz/internal/client"
	"video_service/internal/constants"
	"video_service/internal/dal/db"
)

// NewFeedService create feed service

func (s *VideoService) GetFeed(req *client.FeedReq) (*client.FeedResp, error) {
	resp := &client.FeedResp{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		hlog.Infof("LatestTime is 0")
		lastTime = time.Now()
	} else {
		hlog.Infof("LatestTime is not 0")
		lastTime = time.Unix(req.LatestTime, 0)
	}
	fmt.Printf("LastTime: %v\n", lastTime)
	currentUserId := req.GetUserId()

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		hlog.Error(s.ctx, "GetVideosByLastTime func error: %v", err)
		return nil, constants.NewErrNo(constants.DBErrCode, constants.DBErrMsg)
	}
	// 拷贝视频
	videos := make([]*client.VideoInfo, 0, constants.VideoFeedCount)
	err = s.CopyVideos(&videos, &dbVideos, currentUserId)
	if err != nil {
		hlog.Error(s.ctx, "CopyVideos func error: %v", err)
		return nil, constants.NewErrNo(constants.VideoCopyErrCode, constants.VideoCopyErrMsg)
	}

	resp.VideoList = videos

	// TODO FeedService里是以最后一个视频的发布时间作为下次请求的LatestTime，这样会导致有些视频会被重复请求，需要优化
	// TODO 优化方案：将视频的发布时间作为唯一标识，这样就不会出现重复请求的情况????????
	// TODO 在Response里是以当前时间作为LastTime，这样会导致有些视频会被漏掉，需要优化
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

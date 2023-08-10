package service

import (

	"context"
	"feed_service/biz/constants"
	"feed_service/biz/dal/db"
	dto "feed_service/biz/model/client"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type FeedService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewFeedService create feed service
func NewFeedService(ctx context.Context, c *app.RequestContext) *FeedService {
	return &FeedService{ctx: ctx, c: c}
}

func (s *FeedService) GetFeed(req *dto.FeedReq) (*dto.FeedResp, error) {
	resp := &dto.FeedResp{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime/1000, 0)
	}
	fmt.Printf("LastTime: %v\n", lastTime)
	current_user_id := req.GetUserId()

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return resp, err
	}
	// 拷贝视频
	videos := make([]*dto.VideoInfo, 0, constants.VideoFeedCount)
	err = s.CopyVideos(videos, &dbVideos, current_user_id.(int64))
	if err != nil {
		return resp, nil
	}
	resp.VideoList = videos
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

func (s *FeedService) CopyVideos(result []*dto.VideoInfo, data *[]*db.VideoDBInfo, userId int64) error {
	for _, item := range *data {
		video := s.createVideo(item, userId)
		*result = append(*result, video)
	}
	Video


	return nil
}

func (s *FeedService) createVideo(data *db.VideoDBInfo, userId int64) *dto.VideoInfo {
	video := &dto.VideoInfo{
		Id: data.ID,
		PlayUrl: data.PlayURL,
		CoverUrl:
	}
}

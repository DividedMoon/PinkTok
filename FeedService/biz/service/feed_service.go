package service

import (
	"context"
	"feed_service/biz/constants"
	"feed_service/biz/dal/db"
	client "feed_service/biz/model/client"
	"feed_service/biz/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"sync"
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

func (s *FeedService) GetFeed(req *client.FeedReq) (*client.FeedResp, error) {
	resp := &client.FeedResp{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime/1000, 0)
	}
	fmt.Printf("LastTime: %v\n", lastTime)
	currentUserId := req.GetUserId()

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return resp, err
	}
	// 拷贝视频
	videos := make([]*client.VideoInfo, 0, constants.VideoFeedCount)
	err = s.CopyVideos(&videos, &dbVideos, currentUserId)
	if err != nil {
		return resp, nil
	}
	resp.VideoList = videos
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

func (s *FeedService) CopyVideos(result *[]*client.VideoInfo, data *[]*db.VideoDBInfo, userId int64) error {
	ch := make(chan error)
	for _, item := range *data {
		video, err := s.createVideo(item, userId)
		if err != nil{
			ch <- err
		}
		*result = append(*result, video)
	}
	close(ch)

	select {
	case err:=<-ch:
		return err
	default:
		return nil
	}
}

func (s *FeedService) createVideo(data *db.VideoDBInfo, userId int64) (*client.VideoInfo, error) {
	video := &client.VideoInfo{
		Id: data.ID,
		// 通过前后值将DB中的路径转换为完整的可被访问的路径
		PlayUrl:  utils.URLconvert(s.ctx, s.c, data.PlayURL),
		CoverUrl: utils.URLconvert(s.ctx, s.c, data.CoverURL),
		Title:    data.Title,
	}

	var wg sync.WaitGroup
	wg.Add(4)
	errChannel := make(chan error)

	// 调用UserService获取本条视频的作者信息
	go func() {
		author, err := //TODO call userService to find UserInfo

		if err != nil{
			errChannel <- fmt.Errorf("GetUserInfo func error:" + err.Error())
		}
		video.Author := &UserInfo{

		}

		wg.Done()

	}()

	// 调用VideoService获取视频点赞数量
	go func() {
		err := *new(error)
		video.FavoriteCount, err = // TODO 调用VideoService获取视频点赞数量
		if err != nil {
			errChannel <- fmt.Errorf("GetFavoriteCount func error:" + err.Error())
		}
		wg.Done()
	}()

	// 调用VideoService返回评论数量
	go func() {
		err := *new(error)
		video.CommentCount, err = // TODO 调用VideoService返回评论数量
		if err != nil {
			errChannel <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		wg.Done()
	}()

	// Get favorite exist
	go func() {
		err := *new(error)
		video.IsFavorite, err = db.QueryFavoriteExist(userId, data.ID)
		if err != nil {
			errChannel <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
	close(errChannel)

	select {
	case err :=  <- errChannel:
		return nil, err
	default:
		return video, nil
	}
}

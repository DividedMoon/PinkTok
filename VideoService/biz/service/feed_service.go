package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	"sync"
	"time"
	"video_service/biz/dal/db"
	rd "video_service/biz/dal/redis"
	"video_service/biz/internal/constants"
	"video_service/biz/internal/utils"
	client "video_service/biz/model/client"
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
	// TODO 这个写法是错误的 记得修正
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

		video.FavoriteCount, err = getVideoFavoriteCount(data.ID) // TODO 获取视频点赞数量
		if err != nil {
			errChannel <- fmt.Errorf("GetFavoriteCount func error:" + err.Error())
		}
		wg.Done()
	}()

	// 调用VideoService返回评论数量
	go func() {
		err := *new(error)
		video.CommentCount, err = getVideoCommentCount(data.ID) // TODO 调用VideoService返回评论数量
		if err != nil {
			errChannel <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		wg.Done()
	}()

	// Get favorite exist
	go func() {
		err := *new(error)
		video.IsFavorite, err = queryFavoriteExist(userId, data.ID)
		if err != nil {
			errChannel <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		wg.Done()
	}()

	// TODO 专门开启一个协程来处理错误 一旦出现错误其余协程也要停止 考虑一下close这个是怎么判断的
	go func() {
		select {
		case err :=  <- errChannel:
			return nil, err
		default:
			return video, nil
		}
	}()

	wg.Wait()
	close(errChannel)


}

func getVideoFavoriteCount(videoId int64) (int64, error) {
	count, err := rd.GetVideoFavoriteCount(videoId)
	if err == redis.Nil{
		//从数据库里查询然后更新缓存
		err = buildVideoFavoriteCache(videoId)
		if err != nil{
			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
			return 0, err
		}

		return rd.GetVideoFavoriteCount(videoId)

	}else if err != nil {
		hlog.Error("GetVideoFavoriteCount func error:" + err.Error())
		return 0, err
	}
	return count, nil
}

func getVideoCommentCount(videoId int64)(int64, error){
	count, err := rd.GetVideoCommentCount(videoId)
	if err == redis.Nil{
		//从数据库里查询然后更新缓存
		count, err = db.GetVideoCommentCount(videoId)
		if err != nil {
			hlog.Error("GetVideoCommentCount func error:" + err.Error())
			return 0, err
		}

		//更新缓存
		err = rd.SetVideoCommentCount(videoId, count)
		if err!= nil {
			hlog.Error("SetVideoCommentCount func error:" + err.Error())
			return 0, err
		}
		return count, nil
	}else if err != nil {
		hlog.Error("GetVideoCommentCount func error:" + err.Error())
		return 0, err
	}
	return count, nil
}

func queryFavoriteExist(userId int64, videoId int64) (bool, error) {
	e, err := rd.QueryFavoriteExist(userId, videoId)
	if err == redis.Nil{
		//从数据库里查询然后更新缓存
		err = buildVideoFavoriteCache(videoId)
		if err != nil{
			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
			return false, err
		}
		return rd.QueryFavoriteExist(userId, videoId)
	}else if err != nil {
		hlog.Error("QueryFavoriteExist func error:" + err.Error())
		return false, err
	}

	return e, nil
}

func buildVideoFavoriteCache(videoId int64) error{
	userIds, err := db.GetVideoFavoriteUserIds(videoId)
	if err != nil{
		hlog.Error("GetVideoFavoriteUserIds func error:" + err.Error())
		return err
	}
	//TODO 如果这条数据在数据库里没有 会返回什么错误?
	err = rd.SetVideoFavoriteSet(videoId, userIds)
	if err != nil{
		hlog.Error("SetVideoFavoriteSet func error:" + err.Error())
		return err
	}
	return nil

}
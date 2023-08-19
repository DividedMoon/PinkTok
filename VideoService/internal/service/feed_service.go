package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	"sync"
	"time"
	"video_service/internal/utils"

	//userModel "user_service/biz/model/client"
	client "video_service/biz"
	//userClient "video_service/biz/internal/client"
	"video_service/internal/constants"
	"video_service/internal/dal/db"
	rd "video_service/internal/dal/redis"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService create feed service
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) GetFeed(req *client.FeedReq) (*client.FeedResp, error) {
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

func (s *FeedService) CopyVideos(result *[]*client.VideoInfo, data *[]*db.VideoDBInfo, userId int64) error {

	for _, item := range *data {
		video, err := s.createVideo(item, userId)
		if err != nil {
			return err
		}
		*result = append(*result, video)
	}
	return nil
}

func (s *FeedService) createVideo(data *db.VideoDBInfo, userId int64) (*client.VideoInfo, error) {
	hlog.Infof("createVideo func data: %+v", data)

	video := &client.VideoInfo{
		Id: data.ID,
		// 通过前后值将DB中的路径转换为完整的可被访问的路径
		PlayUrl:  utils.URLconvert(s.ctx, data.PlayURL),
		CoverUrl: utils.URLconvert(s.ctx, data.CoverURL),
		Title:    data.Title,
	}
	hlog.Infof("createVideo func video: %+v", video)

	var wg sync.WaitGroup
	wg.Add(4)
	//ctx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)

	// 调用UserService获取本条视频的作者信息
	go func() {
		//TODO 这里传的是假数据 到时候需要调用UserService
		//userInfoReq := &userModel.UserInfoReq{UserId: data.AuthorID}
		//resp, _, err := userClient.UserServiceClient.UserInfo(s.ctx, userInfoReq)
		//TODO 临时的
		err := *new(error)
		if err != nil {
			errChan <- fmt.Errorf("GetUserInfo func error:" + err.Error())
		}
		hlog.Infof("start to get user info")
		video.Author = &client.UserInfo{
			Id:              1,
			Name:            "testUser",
			FollowCount:     2,
			FollowerCount:   3,
			IsFollow:        true,
			Avatar:          "resp.User.Avatar",
			BackgroundImage: "resp.User.BackgroundImage",
			Signature:       "resp.User.Signature",
			TotalFavorited:  5,
			WorkCount:       6,
			FavoriteCount:   7,
		}
		hlog.Infof("get user info success : %+v", video.Author)

		wg.Done()

	}()

	// 调用VideoService获取视频点赞数量
	go func() {
		err := *new(error)
		hlog.Infof("start to get favorite count")

		video.FavoriteCount, err = getVideoFavoriteCount(data.ID) // TODO 获取视频点赞数量
		if err != nil {
			errChan <- fmt.Errorf("GetFavoriteCount func error:" + err.Error())
		}
		hlog.Infof("get favorite count success : %+v", video.FavoriteCount)
		wg.Done()
	}()

	// 调用VideoService返回评论数量
	go func() {
		err := *new(error)
		hlog.Infof("start to get comment count")
		video.CommentCount, err = getVideoCommentCount(data.ID)
		if err != nil {
			errChan <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		hlog.Infof("get comment count success : %+v", video.CommentCount)
		wg.Done()
	}()

	// Get favorite exist
	go func() {
		err := *new(error)
		hlog.Infof("start to get favorite exist")
		video.IsFavorite, err = queryFavoriteExist(userId, data.ID)
		if err != nil {
			errChan <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
		}
		hlog.Infof("get favorite exist success : %+v", video.IsFavorite)
		wg.Done()
	}()

	hlog.Infof("start to wait")
	wg.Wait()
	hlog.Infof("wait success")
	//处理errChan中的错误
	hasError := false
	close(errChan)

	for err := range errChan {
		if err != nil {
			hasError = true
			hlog.Error(err.Error())
		}
	}

	if hasError {
		return nil, fmt.Errorf("some Errors occur in feedService goroutines")
	} else {
		hlog.Infof("createVideo func finished with no error")
		return video, nil
	}
}

func getVideoFavoriteCount(videoId int64) (int64, error) {
	count, err := rd.GetVideoFavoriteCount(videoId)
	if err == redis.Nil {
		//从数据库里查询然后更新缓存
		err = buildVideoFavoriteCache(videoId)
		if err != nil {
			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
			return 0, err
		}

		return rd.GetVideoFavoriteCount(videoId)

	} else if err != nil {
		hlog.Error("GetVideoFavoriteCount func error:" + err.Error())
		return 0, err
	}
	return count, nil
}

func getVideoCommentCount(videoId int64) (int64, error) {
	count, err := rd.GetVideoCommentCount(videoId)
	if err == redis.Nil {
		//从数据库里查询然后更新缓存
		count, err = db.GetVideoCommentCount(videoId)
		if err != nil {
			hlog.Error("GetVideoCommentCount func error:" + err.Error())
			return 0, err
		}

		//更新缓存
		err = rd.SetVideoCommentCount(videoId, count)
		if err != nil {
			hlog.Error("SetVideoCommentCount func error:" + err.Error())
			return 0, err
		}
		return count, nil
	} else if err != nil {
		hlog.Error("GetVideoCommentCount func error:" + err.Error())
		return 0, err
	}
	return count, nil
}

func queryFavoriteExist(userId int64, videoId int64) (bool, error) {
	e, err := rd.QueryFavoriteExist(userId, videoId)
	if err == redis.Nil {
		//从数据库里查询然后更新缓存
		err = buildVideoFavoriteCache(videoId)
		if err != nil {
			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
			return false, err
		}
		return rd.QueryFavoriteExist(userId, videoId)
	} else if err != nil {
		hlog.Error("QueryFavoriteExist func error:" + err.Error())
		return false, err
	}

	return e, nil
}

func buildVideoFavoriteCache(videoId int64) error {
	userIds, err := db.GetVideoFavoriteUserIds(videoId)
	if err != nil {
		hlog.Error("GetVideoFavoriteUserIds func error:" + err.Error())
		return err
	}
	//TODO 如果这条数据在数据库里没有 会返回什么错误?
	err = rd.SetVideoFavoriteSet(videoId, userIds)
	if err != nil {
		hlog.Error("SetVideoFavoriteSet func error:" + err.Error())
		return err
	}
	return nil

}

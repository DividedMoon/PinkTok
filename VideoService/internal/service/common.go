package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	interactBiz "interact_service/biz"
	"sync"
	userBiz "user_service/biz"
	client "video_service/biz"
	internalClient "video_service/internal/client"
	"video_service/internal/dal/db"
	rd "video_service/internal/dal/redis"
	"video_service/internal/utils"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) CopyVideos(result *[]*client.VideoInfo, data *[]*db.VideoDBInfo, userId int64) error {

	for _, item := range *data {
		video, err := s.createVideo(item, userId)
		if err != nil {
			return err
		}
		*result = append(*result, video)
	}
	return nil
}

func (s *VideoService) createVideo(data *db.VideoDBInfo, userId int64) (*client.VideoInfo, error) {
	hlog.Infof("createVideo func data: %+v", data)

	video := &client.VideoInfo{
		Id: data.ID,
		// 通过前后值将DB中的路径转换为完整的可被访问的路径
		PlayUrl:       utils.URLconvert(s.ctx, data.PlayURL),
		CoverUrl:      utils.URLconvert(s.ctx, data.CoverURL),
		Title:         data.Title,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
	}
	hlog.Infof("createVideo func video: %+v", video)

	var wg sync.WaitGroup
	wg.Add(2)

	errChan := make(chan error)

	// 调用UserService获取本条视频的作者信息
	go func() {
		err := *new(error)
		if err != nil {
			errChan <- fmt.Errorf("GetUserInfo func error:" + err.Error())
		}
		hlog.Infof("start to get user info")
		userInfoReq := &userBiz.UserInfoReq{
			UserId: data.AuthorID,
		}
		userInfoResp, err := internalClient.UserServiceClient.UserInfo(s.ctx, userInfoReq)
		if err != nil {
			hlog.Error("rpc call user service error: ", err)
			errChan <- fmt.Errorf("GetUserInfo func error:" + err.Error())
			wg.Done()
			return
		}
		video.Author = &client.UserInfo{
			Id:              userInfoResp.User.Id,
			Name:            userInfoResp.User.Name,
			FollowCount:     userInfoResp.User.FollowCount,
			FollowerCount:   userInfoResp.User.FollowerCount,
			IsFollow:        userInfoResp.User.IsFollow,
			Avatar:          userInfoResp.User.Avatar,
			BackgroundImage: userInfoResp.User.BackgroundImage,
			Signature:       userInfoResp.User.Signature,
			TotalFavorited:  userInfoResp.User.TotalFavorited,
			WorkCount:       userInfoResp.User.WorkCount,
			FavoriteCount:   userInfoResp.User.FavoriteCount,
		}
		hlog.Infof("get user info success : %+v", video.Author)

		wg.Done()

	}()

	// Get favorite exist
	go func() {
		err := *new(error)
		hlog.Infof("start to get favorite exist")
		queryFavoriteExistReq := &interactBiz.QueryFavoriteExistReq{
			UserId:  userId,
			VideoId: data.ID,
		}
		queryFavoriteExistResp, err := internalClient.InteractServiceClient.QueryFavoriteExist(s.ctx, queryFavoriteExistReq)
		if err != nil || queryFavoriteExistResp.StatusCode != 0 {
			hlog.Error("rpc call interact service error: ", err)
			errChan <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
			wg.Done()
			return
		}
		video.IsFavorite = queryFavoriteExistResp.IsFavorite

		if err != nil {
			errChan <- fmt.Errorf("GetCommentCountByVideoID func error:" + err.Error())
			wg.Done()
			return
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

func getVideoInfo(videoId int64) (*db.VideoDBInfo, error) {
	video, err := rd.GetVideoHash(videoId)
	if err == redis.Nil {
		video, err = db.GetVideoByID(videoId)
		if err != nil {
			hlog.Error("GetVideoByID func error:" + err.Error())
			return nil, err
		}
		if video.CommentCount > 100 && video.FavoriteCount > 1000 {
			err = buildVideoInfoCache(video)
		}
	}
	return video, nil
}

//func getVideoFavoriteCount(videoId int64) (int64, error) {
//	count, err := rd.GetVideoFavoriteCount(videoId)
//	if err == redis.Nil {
//		//从数据库里查询然后更新缓存
//		err = buildVideoFavoriteCache(videoId)
//		if err != nil {
//			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
//			return 0, err
//		}
//
//		return rd.GetVideoFavoriteCount(videoId)
//
//	} else if err != nil {
//		hlog.Error("GetVideoFavoriteCount func error:" + err.Error())
//		return 0, err
//	}
//	return count, nil
//}

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

//func queryFavoriteExist(userId int64, videoId int64) (bool, error) {
//	e, err := rd.QueryFavoriteExist(userId, videoId)
//	if err == redis.Nil {
//		//从数据库里查询然后更新缓存
//		err = buildVideoFavoriteCache(videoId)
//		if err != nil {
//			hlog.Error("buildVideoFavoriteCache func error:" + err.Error())
//			return false, err
//		}
//		return rd.QueryFavoriteExist(userId, videoId)
//	} else if err != nil {
//		hlog.Error("QueryFavoriteExist func error:" + err.Error())
//		return false, err
//	}
//
//	return e, nil
//}

//func buildVideoFavoriteCache(videoId int64) error {
//	userIds, err := db.GetVideoFavoriteUserIds(videoId)
//	if err != nil {
//		hlog.Error("GetVideoFavoriteUserIds func error:" + err.Error())
//		return err
//	}

//	err = rd.SetVideoFavoriteSet(videoId, userIds)
//	if err != nil {
//		hlog.Error("SetVideoFavoriteSet func error:" + err.Error())
//		return err
//	}
//	return nil
//}

func buildVideoInfoCache(video *db.VideoDBInfo) error {
	err := rd.SetVideoHash(video)
	if err != nil {
		hlog.Error("SetVideoHash func error:" + err.Error())
		return err
	}
	return nil
}

package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	client "video_service/biz"
	"video_service/internal/dal/db"
	rd "video_service/internal/dal/redis"
)

func (s *VideoService) FavoriteAction(videoId, userId int64, actionType int32) error {
	//1为点赞，0为取消点赞
	// TODO 1. 调用interactService去修改点赞表 注意要等待返回值
	err := error(nil)

	// 2. 修改video表中的点赞数，有两种，对于热门视频直接修改redis中的值，对于冷门视频可以直接修改数据库
	if actionType == 1 { // 如果操作类型为1则为点赞
		err = AddVideoFavoriteCount(videoId, 1)
	} else if actionType == 0 { // 如果操作类型为0 则为取消赞操作
		err = AddVideoFavoriteCount(videoId, -1)
	} else { // 如果不是这两种类型 则返回错误
		hlog.Error("FavoriteAction error", "actionType is not 1 or -1")
		return fmt.Errorf("actionType is not 1 or -1")
	}

	if err != nil {
		hlog.Error("AddVideoFavoriteCount error", err.Error())
		return err
	}
	// 3. 返回结果
	return nil
}

func (s *VideoService) GetFavoriteVideoList(userId int64) ([]*client.VideoInfo, error) {

	// TODO 1. 调用 interactService 获取用户点赞的视频列表 返回值是videoIds 视频ID列表
	err := error(nil)
	favoriteVideoIds := []int64{1, 2, 3, 4, 5}
	// 2. 调用copyVideo方法获取视频信息并返回
	videoDBInfos, err := db.GetVideoDBInfoByIDs(favoriteVideoIds)
	if err != nil {
		hlog.Error("GetVideoDBInfoByIDs error", err.Error())
		return nil, err
	}
	videos := make([]*client.VideoInfo, 0, len(favoriteVideoIds))
	err = s.CopyVideos(&videos, &videoDBInfos, userId)
	if err != nil {
		hlog.Error("CopyVideos error", err.Error())
		return nil, err
	}
	return videos, nil
}

func AddVideoFavoriteCount(videoId, increment int64) error {
	// 1.先查redis里有没有 如果有的话直接加一或者减一返回
	video, err := rd.GetVideoHash(videoId)
	if err != nil && err != redis.Nil {
		hlog.Error("GetVideoFavoriteCount error", err.Error())
		return err
	}
	if video != nil { // video不为空 直接更新redis后返回
		err = rd.SetVideoField(videoId, "FavoriteCount", video.FavoriteCount+increment)
		if err != nil {
			hlog.Error("SetVideoField error", err.Error())
			return err
		}
		return nil
	}
	// 2.如果没有的话，查数据库
	video, err = db.GetVideoByID(videoId)
	if err != nil {
		hlog.Error("GetVideoFavoriteCount error", err.Error())
		return err
	}
	// 2.1如果数据库中的点赞数大于1000 则认为是热门视频，将点赞数 + 1，写入redis
	if video.FavoriteCount >= 1000 {
		video.FavoriteCount += increment
		err = buildVideoInfoCache(video)
		if err != nil {
			hlog.Error("buildVideoInfoCache error", err.Error())
			return err
		}
	}
	// 2.2更新数据库
	err = db.UpdateVideoFavoriteCount(videoId, video.FavoriteCount+increment)
	if err != nil {
		hlog.Error("UpdateVideoFavoriteCount error", err.Error())
		return err
	}
	return nil
}

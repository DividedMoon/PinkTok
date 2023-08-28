package redis

import (
	"strconv"
	"time"
	"video_service/internal/dal/db"
	"video_service/internal/utils"
)

// GetVideoFavoriteCount 获取视频点赞数
func GetVideoFavoriteCount(videoId int64) (int64, error) {
	return countSetSize(rdFavorite, utils.GetVideoFavoriteKey(videoId))
}

// GetVideoCommentCount 获取视频评论数量
func GetVideoCommentCount(videoId int64) (int64, error) {
	return getStringValue(rdComment, utils.GetVideoCommentKey(videoId))
}

// QueryFavoriteExist 查询用户是否对视频点赞
func QueryFavoriteExist(userId int64, videoId int64) (bool, error) {
	return checkSetExist(rdFavorite, utils.GetVideoFavoriteKey(videoId), userId)
}

func SetVideoFavoriteSet(videoId int64, userIds []int64) error {
	return initSet(rdFavorite, utils.GetVideoFavoriteKey(videoId), userIds)
}

func SetVideoCommentCount(videoId int64, count int64) error {
	return initString(rdComment, utils.GetVideoCommentKey(videoId), count)
}

func SetVideoHash(video *db.VideoDBInfo) error {
	videoMap := make(map[string]interface{})
	videoMap["ID"] = video.ID
	videoMap["AuthorID"] = video.AuthorID
	videoMap["PlayURL"] = video.PlayURL
	videoMap["CoverURL"] = video.CoverURL
	videoMap["PublishTime"] = video.PublishTime
	videoMap["Title"] = video.Title
	videoMap["FavoriteCount"] = video.FavoriteCount
	videoMap["CommentCount"] = video.CommentCount
	err := initHash(rdVideo, utils.GetVideoKey(video.ID), videoMap)
	if err != nil {
		return err
	}
	return nil
}

func GetVideoHash(videoId int64) (*db.VideoDBInfo, error) {
	videoHash, err := getHash(rdVideo, utils.GetVideoKey(videoId))
	//hlog.Info("videoHash", videoHash)
	//hlog.Infof("err:%v", err)
	if err != nil {
		return nil, err
	}
	// 如果视频不存在Redis中，返回nil
	if len(videoHash) == 0 {
		return nil, nil
	}
	id, _ := strconv.ParseInt(videoHash["ID"], 10, 64)
	authorId, _ := strconv.ParseInt(videoHash["AuthorID"], 10, 64)
	favoriteCount, _ := strconv.ParseInt(videoHash["FavoriteCount"], 10, 64)
	commentCount, _ := strconv.ParseInt(videoHash["CommentCount"], 10, 64)
	layout := "2006-01-02T15:04:05-07:00"
	publishTime, err := time.Parse(layout, videoHash["PublishTime"])
	if err != nil {
		return nil, err
	}

	video := db.VideoDBInfo{
		ID:            id,
		AuthorID:      authorId,
		PlayURL:       videoHash["PlayURL"],
		CoverURL:      videoHash["CoverURL"],
		PublishTime:   publishTime,
		Title:         videoHash["Title"],
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
	}
	return &video, nil

}

func SetVideoField(videoId int64, field string, value interface{}) error {
	return setHashField(rdVideo, utils.GetVideoKey(videoId), field, value)
}

func GetVideoField(videoId int64, field string) (string, error) {
	return getHashField(rdVideo, utils.GetVideoKey(videoId), field)
}

package redis

import (
	"video_service/internal/utils"
)

//GetVideoFavoriteCount 获取视频点赞数
func GetVideoFavoriteCount(videoId int64) (int64, error) {
	return countSetSize(rdFavorite, utils.GetVideoFavoriteKey(videoId))
}

//GetVideoCommentCount 获取视频评论数量
func GetVideoCommentCount(videoId int64) (int64, error) {
	return getStringValue(rdComment, utils.GetVideoCommentKey(videoId))
}

//QueryFavoriteExist 查询用户是否对视频点赞
func QueryFavoriteExist(userId int64, videoId int64) (bool, error) {
	return checkSetExist(rdFavorite, utils.GetVideoFavoriteKey(videoId), userId)
}

func SetVideoFavoriteSet(videoId int64, userIds []int64) error {
	return initSet(rdFavorite, utils.GetVideoFavoriteKey(videoId), userIds)
}

func SetVideoCommentCount(videoId int64, count int64) error {
	return initString(rdComment, utils.GetVideoCommentKey(videoId), count)
}

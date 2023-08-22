package utils

import "strconv"

const split = ":"
const videoFavoriteSuffix = "favorite" + split + "video"
const commentVideoSuffix = "comment" + split + "video"
const videoSuffix = "video"
const userSuffix = "user"

func GetVideoFavoriteKey(videoId int64) string {
	return videoFavoriteSuffix + split + strconv.FormatInt(videoId, 10)
}

func GetVideoCommentKey(videoId int64) string {
	return commentVideoSuffix + split + strconv.FormatInt(videoId, 10)
}

func GetVideoKey(videoId int64) string {
	return videoSuffix + split + strconv.FormatInt(videoId, 10)
}

func GetUserKey(userId int64) string {
	return userSuffix + split + strconv.FormatInt(userId, 10)
}

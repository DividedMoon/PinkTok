package utils

import "strconv"

const split = ":"
const videoFavoriteSuffix = "favorite" + split + "video"
const commentVideoSuffix = "comment" + split + "video"

func GetVideoFavoriteKey(videoId int64) string {
	return videoFavoriteSuffix + split + strconv.FormatInt(videoId, 10)
}

func GetVideoCommentKey(videoId int64) string {
	return commentVideoSuffix + split + strconv.FormatInt(videoId, 10)
}

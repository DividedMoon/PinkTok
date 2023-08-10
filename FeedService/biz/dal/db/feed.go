package db

import (
	"feed_service/biz/constants"
	"time"
)

//TODO 验证数据库中的表是否与之对应
type VideoDBInfo struct {
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func GetVideosByLastTime(lastTime time.Time) ([]*VideoDBInfo, error) {
	videos := make([]*VideoDBInfo, constants.VideoFeedCount)
	err := DB.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

package db

import (
	"time"
	"video_service/biz/internal/constants"
)

type VideoDBInfo struct {
	ID          int64     `gorm:"column:id"`
	AuthorID    int64     `gorm:"column:author_id"`
	PlayURL     string    `gorm:"column:play_url"`
	CoverURL    string    `gorm:"column:cover_url"`
	PublishTime time.Time `gorm:"column:created_time"`
	Title       string    `gorm:"column:title"`

	//以下作为其他推荐算法的保留属性
	FavoriteCount int64 `gorm:"column:favorite_count"`
	CommentCount  int64 `gorm:"column:comment_count"`
}

func (VideoDBInfo) TableName() string {
	return constants.VideosTableName
}

func GetVideosByLastTime(lastTime time.Time) ([]*VideoDBInfo, error) {
	videos := make([]*VideoDBInfo, constants.VideoFeedCount)
	err := DB.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

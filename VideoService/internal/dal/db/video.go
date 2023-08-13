package db

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
	"video_service/internal/constants"
)

type Base struct {
	UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime;column:updated_time"`
	CreatedAt *time.Time `json:"-" gorm:"autoCreateTime;column:created_time"`
	Deleted   int        `json:"-" gorm:"column:deleted"`
}
type VideoDBInfo struct {
	Base
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
type FavoriteDBInfo struct {
	Base
	ID      int64 `gorm:"column:id"`
	VideoId int64 `gorm:"column:video_id"`
	UserId  int64 `gorm:"column:user_id"`
}

func (VideoDBInfo) TableName() string {
	return constants.VideosTableName
}

func (FavoriteDBInfo) TableName() string {
	return constants.FavoriteTableName
}

func GetVideosByLastTime(lastTime time.Time) ([]*VideoDBInfo, error) {
	videos := make([]*VideoDBInfo, constants.VideoFeedCount)
	err := DB.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
	if err != nil {
		hlog.Error("GetVideosByLastTime", "err", err)
		return videos, err
	}
	return videos, nil
}

func GetVideoFavoriteUserIds(videoId int64) ([]int64, error) {
	var favorites []FavoriteDBInfo
	if err := DB.Where("video_id = ?", videoId).Find(&favorites).Error; err != nil {
		hlog.Error("GetVideoFavoriteUserIds", "err", err)
		return nil, err
	}
	var userIds []int64
	for _, item := range favorites {
		userIds = append(userIds, item.UserId)
	}
	return userIds, nil
}

func GetVideoCommentCount(videoId int64) (int64, error) {
	var count int64
	err := DB.Model(&VideoDBInfo{}).Where("id = ?", videoId).Select("comment_count").Scan(&count).Error
	if err != nil {
		hlog.Error("GetCommentCount", "err", err)
		return -1, err
	}
	return count, nil
}

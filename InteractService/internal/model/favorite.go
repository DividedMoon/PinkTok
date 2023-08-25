package model

import (
	"errors"
	"gorm.io/gorm"
)

type Favorite struct {
	Base
	ID      int64 `json:"id" gorm:"column:id"`
	VideoID int64 `json:"video_id" gorm:"column:video_id"`
	UserID  int64 `json:"user_id" gorm:"column:user_id"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}

func IsVideoLikedByUser(videoId, userId int64) (bool, error) {
	var count int64
	err := DB.Where("video_id = ? and user_id = ? and deleted = 0", videoId, userId).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return count > 0, nil
}

func UpdateVideoLikedStatus(userID, videoID int64, liked bool) error {
	// 查询是否已存在记录
	var favorite Favorite
	err := DB.Model(&Favorite{}).
		Where("user_id = ? AND video_id = ?", userID, videoID).
		First(&favorite).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 更新或插入记录
	if liked {
		if err == gorm.ErrRecordNotFound {
			favorite = Favorite{
				VideoID: videoID,
				UserID:  userID,
				Base: Base{
					Deleted: 0,
				},
			}
		} else {
			favorite.Deleted = 1
		}
	} else {
		if err == nil {
			favorite.Deleted = 1
		}
	}

	err = DB.Save(&favorite).Error
	return err
}

func SelectFavoriteVideoIdsByUserID(userId int64) (videoIds []int64, err error) {
	var favorites []Favorite
	err = DB.Where("user_id = ? and deleted = 0", userId).
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoID)
	}
	return videoIds, nil
}

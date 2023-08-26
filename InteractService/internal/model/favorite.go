package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"interact_service/internal/constant"
)

type Favorite struct {
	Base
	ID      int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	VideoID int64 `json:"video_id" gorm:"column:video_id"`
	UserID  int64 `json:"user_id" gorm:"column:user_id"`
}

func (f *Favorite) TableName() string {
	return constant.FavoriteTableName
}

func IsVideoLikedByUser(videoId, userId int64) (bool, error) {
	var count int64
	err := DB.Model(&Favorite{}).Where("video_id = ? and user_id = ? and deleted = 0", videoId, userId).
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
	if !liked {
		if err == gorm.ErrRecordNotFound { //如果原始状态为不喜欢，且没有记录，则插入一条喜欢记录
			favorite = Favorite{
				VideoID: videoID,
				UserID:  userID,
				Base: Base{
					Deleted: 0,
				},
			}
		} else { //如果原始状态为不喜欢且存在记录，则将记录标记为未删除
			favorite.Deleted = 0
		}
	} else { // 如果原始状态喜欢且仍然存在错误 则说明出错
		if err == nil {
			return fmt.Errorf("user already liked video but dont have record")
		} else { //如果原始状态为喜欢且存在记录，则将该记录标记为删除
			favorite.Deleted = 1
		}
	}

	DB.Error = nil
	dbSaveErr := DB.Model(&Favorite{}).Save(&favorite).Error

	return dbSaveErr
}

func SelectFavoriteVideoIdsByUserID(userId int64) (videoIds []int64, err error) {
	var favorites []Favorite
	err = DB.Model(&Favorite{}).Where("user_id = ? and deleted = 0", userId).
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoID)
	}
	return videoIds, nil
}

//func getTable(userId int64) *gorm.DB {
//	shardingIndex := userId % constant.FavoriteSharding
//	tableName := constant.FavoriteTableName + "_" + strconv.FormatInt(shardingIndex, 10)
//
//	return DB.Table(tableName)
//}

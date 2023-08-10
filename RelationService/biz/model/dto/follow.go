package dto

import (
	"gorm.io/gorm"
	"relation_service/biz/model"
	"time"
)

type Base struct {
	UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime;column:updated_time"`
	CreatedAt *time.Time `json:"-" gorm:"autoCreateTime;column:created_time"`
	Deleted   int        `json:"-" gorm:"column:deleted"`
}

type Follow struct {
	Base
	ID      int64 `json:"id" gorm:"column:id"`
	UserIdA int64 `json:"user_id_a" gorm:"column:user_id_a"`
	UserIdB int64 `json:"user_id_b" gorm:"column:user_id_b"`
}

func (*Follow) TableName() string {
	return "follow"
}

func (f *Follow) Create(tx *gorm.DB) error {
	return tx.Create(&f).Error
}

func SelectFollowByUserIdAAndUserIdB(userIdA, userIdB int64) (f Follow, err error) {
	err = model.DB.
		Where("user_id_a = ? AND user_id_b = ? AND deleted = 0", userIdA, userIdB).
		First(&f).
		Error
	return
}

func SelectFollowByUserIdA(userIdA int64) (fs []Follow, err error) {
	err = model.DB.
		Where("user_id_a = ? AND deleted = 0", userIdA).
		Find(&fs).
		Error
	return
}

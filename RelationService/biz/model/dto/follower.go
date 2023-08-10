package dto

import (
	"gorm.io/gorm"
	"relation_service/biz/model"
)

type Follower struct {
	Base
	ID      int64 `json:"id" gorm:"column:id"`
	UserIdA int64 `json:"user_id_a" gorm:"column:user_id_a"`
	UserIdB int64 `json:"user_id_b" gorm:"column:user_id_b"`
}

func (*Follower) TableName() string {
	return "follower"
}

func (f *Follower) Create(tx *gorm.DB) error {
	return tx.Create(&f).Error
}

func SelectFollowerByUserIdAAndUserIdB(userIdA, userIdB int64) (f Follower, err error) {
	err = model.DB.
		Where("user_id_a = ? AND user_id_b = ? AND deleted = 0", userIdA, userIdB).
		First(&f).
		Error
	return
}

func SelectFollowerByUserIdA(userIdA int64) (fs []Follower, err error) {
	err = model.DB.
		Where("user_id_a = ? AND deleted = 0", userIdA).
		Find(&fs).
		Error
	return
}

func SelectFriendByUserIdA(userIdA int64) (fs []Follower, err error) {
	err = model.DB.
		Where("user_id_a = ? AND deleted = 0", userIdA).
		Find(&fs).
		Error
	return
}

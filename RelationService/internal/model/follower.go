package model

import (
	"database/sql"
	"gorm.io/gorm"
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

func (f *Follower) HardDelete() {
	DB.Unscoped().Delete(&f)
}

func SelectFollowerByUserIdAAndUserIdB(userIdA, userIdB int64) (f Follower, err error) {
	err = DB.
		Where("user_id_a = ? AND user_id_b = ? AND deleted = 0", userIdA, userIdB).
		First(&f).
		Error
	return
}

func SelectFollowerByUserIdA(userIdA int64) (fs []Follower, err error) {
	err = DB.
		Where("user_id_a = ? AND deleted = 0", userIdA).
		Find(&fs).
		Error
	return
}

func SelectFriendByUserIdA(userIdA int64) (fs []Follower, err error) {
	err = DB.
		Table("follower").
		Joins("JOIN follow ON follower.user_id_b = follow.user_id_b").
		Where("follow.user_id_a = ? AND follower.user_id_a = ? AND follower.deleted = 0", userIdA, userIdA).
		Find(&fs).
		Error
	return
}

func SubmitFollow(follow *Follow, follower *Follower) error {
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(follow).Error; err != nil {
			return err
		}
		if err := tx.Save(follower).Error; err != nil {
			return err
		}
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return err
	}
	return nil
}

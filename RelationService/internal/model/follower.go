package model

import (
	"database/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"relation_service/internal/constant"
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

func SubmitFollowAndFollower(follow *Follow, follower *Follower) error {
	tx := DB.Begin(&sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	affected := tx.Model(follow).
		Update("deleted", 0).
		Where("user_id_a = ? AND user_id_b = ?", follow.UserIdA, follow.UserIdB).RowsAffected
	if affected == 0 {
		err = tx.Create(follow).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if affected != 1 {
		tx.Rollback()
		hlog.Errorf("affected rows is not 1, affected: %d, dto: %+v", affected, follow)
		return constant.UpdateNotEqualOneErr
	}
	affected = tx.Model(follower).
		Update("deleted", 0).
		Where("user_id_a = ? AND user_id_b = ?", follower.UserIdA, follower.UserIdB).RowsAffected
	if affected == 0 {
		err = tx.Create(follower).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if affected != 1 {
		tx.Rollback()
		hlog.Errorf("affected rows is not 1, affected: %d, dto: %+v", affected, follower)
		return constant.UpdateNotEqualOneErr
	}
	return nil
}

func CancelFollowAndFollower(follow *Follow, follower *Follower) error {
	tx := DB.Begin(&sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	err = tx.Model(follow).
		Update("deleted", 1).
		Where("user_id_a = ? AND user_id_b = ?", follow.UserIdA, follow.UserIdB).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(follower).
		Update("deleted", 1).
		Where("user_id_a = ? AND user_id_b = ?", follower.UserIdA, follower.UserIdB).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	hlog.Info("cancel follow and follower success")
	return nil
}

package dto

import (
	"fmt"
	"strings"
	"time"
	"user_service/internal/constant"
	"user_service/internal/model"
	"user_service/internal/util"
)

type Base struct {
	UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime;column:updated_time"`
	CreatedAt *time.Time `json:"-" gorm:"autoCreateTime;column:created_time"`
	Deleted   int        `json:"-" gorm:"column:deleted"`
}

type User struct {
	Base
	ID              int64  `json:"id" gorm:"column:id"`
	UserName        string `json:"name" gorm:"column:username"`
	Password        string `gorm:"column:password"`
	Name            string `gorm:"column:name"`
	FollowCount     int64  `json:"follow_count" gorm:"column:follow_count"`
	FollowerCount   int64  `json:"follower_count" gorm:"column:follower_count"`
	IsFollow        bool   `gorm:"-"`
	Avatar          string `json:"avatar" gorm:"column:avatar"`
	BackgroundImage string `json:"background_image" gorm:"column:background_image"`
	Signature       string `json:"signature" gorm:"column:signature"`
	TotalFavorited  int64  `json:"total_favorited" gorm:"column:total_favorited"`
	WorkCount       int64  `json:"work_count" gorm:"column:work_count"`
	FavoriteCount   int64  `json:"favorite_count" gorm:"column:favorite_count"`
}

func (*User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	return model.DB.Create(&u).Error
}

func (u *User) Update() error {
	affected := model.DB.
		Model(u).
		Updates(u).
		Where("deleted = 0").
		Limit(2).
		RowsAffected
	if affected != 1 {
		return constant.AffectedRowIsNotEqualOne.WithMessage("update User err")
	}
	return nil
}

func (u *User) SelectById(id int64) error {
	return model.DB.
		Where("id = ? AND deleted = 0", id).
		First(&u).
		Error
}

func (u *User) SelectByUsername(username string) error {
	return model.DB.
		Where("username = ? AND deleted = 0", username).
		First(&u).
		Error
}

func (u *User) UpdateByCountMap(origin, changes map[string]int) error {
	var (
		wheres  []string
		equals  []interface{}
		updates = map[string]interface{}{}
	)
	for k, v := range origin {
		wheres = append(wheres, fmt.Sprintf("%s = ?", utils.CameCaseToUnderscore(k)))
		equals = append(equals, v)
		updates[k] = v + changes[k]
	}
	wheres = append(wheres, "deleted = 0")
	result := model.DB.
		Model(u).
		Updates(updates).
		Where(strings.Join(wheres, " AND "), equals...).
		Limit(2)
	affected := result.RowsAffected
	if affected != 1 {
		return constant.AffectedRowIsNotEqualOne.WithMessage(result.Error.Error())
	}
	return nil
}

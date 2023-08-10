package dto

import (
	"gorm.io/gorm"
	"user_service/biz/model"
)

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

func (u *User) UpdateByUser(tx *gorm.DB) (err error) {

}

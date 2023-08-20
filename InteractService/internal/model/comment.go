package model

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Base struct {
	UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime;column:updated_time"`
	CreatedAt *time.Time `json:"-" gorm:"autoCreateTime;column:created_time"`
	Deleted   int        `json:"-" gorm:"column:deleted"`
}

type Comment struct {
	Base
	ID      int64  `json:"id" gorm:"column:id"`
	UserId  int64  `json:"user_id" gorm:"column:user_id"`
	VideoId int64  `json:"video_id" gorm:"column:video_id"`
	Content string `json:"content" gorm:"content"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Create() error {
	err = DB.Create(c).Error
	if err != nil {
		sqlerr := new(mysql.MySQLError)
		if errors.As(err, &sqlerr); sqlerr.Number == 1062 {
			// 唯一索引冲突，返回正常值
			return DB.Model(c).
				Where("user_id = ? and video_id = ? and content = ? and deleted = 0",
					c.UserId, c.VideoId, c.Content).
				First(c).Error
		}
		return err
	}
	return nil
}

func (c *Comment) SelectByVideoId() (list []Comment, err error) {
	err = DB.Model(c).
		Where("video_id = ? and deleted = 0", c.VideoId).
		Order("created_time desc").
		Find(&list).Error
	return list, err
}

func (c *Comment) SelectByUserIdAndVideoId() (list []Comment, err error) {
	err = DB.Model(c).
		Where("user_id = ? and video_id = ? and deleted = 0",
			c.UserId, c.VideoId).
		Order("created_time desc").
		Find(&list).Error
	return list, err
}

func (c *Comment) DeleteById() error {
	return DB.Model(c).
		Update("deleted", 1).
		Where("id = ? and user_id = ? and video_id = ? and deleted = 0",
			c.ID, c.UserId, c.VideoId).
		Error
}

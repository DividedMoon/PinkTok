package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"user_service/biz"
)

func TestUpdateUserInfo(t *testing.T) {
	var (
		user = &biz.UserInfo{
			Id:              1,
			Name:            "Lhj",
			FollowCount:     0,
			FollowerCount:   1,
			IsFollow:        false,
			Avatar:          "",
			BackgroundImage: "",
			Signature:       "",
			TotalFavorited:  0,
			WorkCount:       1,
			FavoriteCount:   0,
		}
		changes = map[string]int{
			"FollowCount": 3,
			"WorkCount":   4,
		}
	)
	_, err := UpdateUserInfo(user, changes)
	assert.Assert(t, err == nil)
	fmt.Println(user)
}

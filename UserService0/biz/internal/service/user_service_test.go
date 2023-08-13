package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"user_service/biz/model/client"
)

func TestUpdateUserInfo(t *testing.T) {
	var (
		user = &client.UserInfo{
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
	err := UpdateUserInfo(user, changes)
	assert.Assert(t, err == nil)
	fmt.Println(user)
}

package redis

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"user_service/biz"
)

func TestMain(m *testing.M) {
	InitRedis()
	m.Run()
	CloseRedis()
}

func TestAddUser(t *testing.T) {
	user := biz.UserInfo{
		Id:              2,
		Name:            "LHJWEZ",
		FollowCount:     1,
		FollowerCount:   1,
		IsFollow:        true,
		Avatar:          "Nihao",
		BackgroundImage: "Haha",
		Signature:       "Huhu",
		TotalFavorited:  3,
		WorkCount:       40,
		FavoriteCount:   92,
	}

	AddUser(&user)
	existUser := ExistUser(2)
	assert.Assert(t, existUser)
}

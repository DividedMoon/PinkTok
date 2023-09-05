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

func TestGetUser(t *testing.T) {
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
	info := GetUser(user.Id)
	assert.NotNil(t, info)
	assert.Assert(t, info.Id == user.Id)
	assert.Assert(t, info.Name == user.Name)
}

func TestUpdateUserByMap(t *testing.T) {
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
	change := map[string]int{
		"FollowCount": 2,
		"WorkCount":   50,
	}
	UpdateUserByMap(user.Id, change)
	info := GetUser(user.Id)
	assert.Assert(t, info.FollowCount == 3)
	assert.Assert(t, info.WorkCount == 90)
}

package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"math/rand"
	"reflect"
	"time"
	"user_service/biz"
	"user_service/internal/constant"
	"user_service/internal/model/dto"
)

func RegisterUser(username, password string) (u *biz.UserInfo, err error) {
	var user = &dto.User{
		UserName:        username,
		Password:        password,
		Name:            generateRandomName(5),
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRG.img",
		Signature:       "这个人很懒，什么都没有留下。",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	err = user.Create()
	if err != nil {
		return nil, err
	}
	u = &biz.UserInfo{
		Id:              user.ID,
		Name:            user.UserName,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return u, nil
}

func GetUserInfo(userId int64) (u *biz.UserInfo, err error) {
	var user = &dto.User{
		ID: userId,
	}
	err = user.SelectById(userId)
	if err != nil {
		return nil, err
	}
	hlog.Info("user", user)
	u = &biz.UserInfo{
		Id:              user.ID,
		Name:            user.UserName,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return u, nil
}

func UpdateUserInfo(u *biz.UserInfo, changes map[string]int) (string, error) {
	var (
		user = &dto.User{
			ID:             u.Id,
			FollowCount:    u.FollowCount,
			FollowerCount:  u.FollowerCount,
			IsFollow:       false,
			TotalFavorited: u.TotalFavorited,
			WorkCount:      u.WorkCount,
			FavoriteCount:  u.FavoriteCount,
		}
		origin = map[string]int{}
	)
	// 设置初始值
	elem := reflect.ValueOf(user).Elem()
	for k := range changes {
		i := elem.FieldByName(k).Interface().(int64)
		origin[k] = int(i)
	}
	// 更新数据库
	err := user.UpdateByCountMap(origin, changes)
	if err != nil {
		return "", err
	}
	return user.Base.UpdatedAt.Format(constant.DateFormat), nil
}

func generateRandomName(n int) string {
	name := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range name {
		se := 19968 + rand.Int63n(40869-19968)
		name[i] = rune(se)
	}
	return string(name)
}

package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reflect"
	"strconv"
	"user_service/biz"
)

const (
	userSuffix = "user:"
)

func getUserKey(userId int64) string {
	return userSuffix + strconv.FormatInt(userId, 10)
}

// AddUser 添加用户
func AddUser(info *biz.UserInfo) {
	if rdb == nil {
		return
	}
	key := getUserKey(info.Id)
	userInfoValue := reflect.Indirect(reflect.ValueOf(info))
	userInfoType := userInfoValue.Type()
	tx := rdb.TxPipeline()
	for i := 0; i < userInfoValue.NumField(); i++ {
		hlog.Infof("fieldName = %+v", userInfoType.Field(i).Name)
		fieldName := userInfoType.Field(i).Name
		if fieldName == "state" || fieldName == "sizeCache" || fieldName == "unknownFields" {
			continue
		}
		fieldValue := userInfoValue.Field(i).Interface()

		err := rdb.HSet(key, fieldName, fieldValue).Err()
		if err != nil {
			hlog.Errorf("HSet failed, err = %+v", err)
			_ = tx.Discard()
			return
		}
	}
	_, _ = tx.Exec()
}

// GetUser 获取用户
func GetUser(userId int64) (info *biz.UserInfo) {
	if rdb == nil {
		return nil
	}
	key := getUserKey(userId)
	userInfoValue := reflect.ValueOf(&biz.UserInfo{}).Elem()
	userInfoType := userInfoValue.Type()

	for i := 0; i < userInfoValue.NumField(); i++ {
		fieldName := userInfoType.Field(i).Name

		fieldValue, err := rdb.HGet(key, fieldName).Result()
		if err != nil {
			hlog.Errorf("HGet failed, err = %+v", err)
			return nil
		}

		fieldKind := userInfoValue.Field(i).Kind()

		switch fieldKind {
		case reflect.String:
			userInfoValue.Field(i).SetString(fieldValue)
		case reflect.Int:
			fieldValueInt, err := strconv.Atoi(fieldValue)
			if err != nil {
				hlog.Errorf("Atoi failed, err = %+v", err)
				return nil
			}
			userInfoValue.Field(i).SetInt(int64(fieldValueInt))
		default:
			// do nothing
		}
	}

	userInfo := userInfoValue.Interface().(*biz.UserInfo)
	return userInfo
}

// DelUser 删除用户
func DelUser(userId int64) {
	del(rdb, getUserKey(userId), userId)
}

// ExistUser 检查用户是否存在
func ExistUser(userId int64) bool {
	if rdb == nil {
		return false
	}
	return rdb.Exists(getUserKey(userId)).Val() == 1
}

package redis

import "strconv"

const (
	followSuffix   = "follow_set:"
	followerSuffix = "follower_set:"
)

// AddFollow 在userIdA的关注列表中添加userIdB
func AddFollow(userIdA, userIdB int64) {
	if rdbFollow == nil {
		return
	}
	add(rdbFollow, followSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// DelFollow 在userIdA的关注列表中删除userIdB
func DelFollow(userIdA, userIdB int64) {
	del(rdbFollow, followSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// ExistFollow 检查userIdA的关注列表中是否存在userIdB
func ExistFollow(userIdA, userIdB int64) bool {
	if rdbFollow == nil {
		return false
	}
	return exist(rdbFollow, followSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// CountFollow 获取userIdA的关注列表的长度
func CountFollow(userIdA int64) (int64, error) {
	return count(rdbFollow, followSuffix+strconv.FormatInt(userIdA, 10))
}

// GetFollow 获取userIdA的关注列表
func GetFollow(userIdA int64) []int64 {
	return get(rdbFollow, followSuffix+strconv.FormatInt(userIdA, 10))
}

// AddFollower 在userIdA的粉丝列表中添加userIdB
func AddFollower(userIdA, userIdB int64) {
	if rdbFollow == nil {
		return
	}
	add(rdbFollow, followerSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// DelFollower 在userIdA的粉丝列表中删除userIdB
func DelFollower(userIdA, userIdB int64) {
	del(rdbFollow, followerSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// ExistFollower 检查userIdA的粉丝列表中是否存在userIdB
func ExistFollower(userIdA, userIdB int64) bool {
	if rdbFollow == nil {
		return false
	}
	return exist(rdbFollow, followerSuffix+strconv.FormatInt(userIdA, 10), userIdB)
}

// CountFollower 获取userIdA的粉丝列表的长度
func CountFollower(userIdA int64) (int64, error) {
	return count(rdbFollow, followerSuffix+strconv.FormatInt(userIdA, 10))
}

// GetFollower 获取userIdA的粉丝列表
func GetFollower(userIdA int64) []int64 {
	return get(rdbFollow, followerSuffix+strconv.FormatInt(userIdA, 10))
}

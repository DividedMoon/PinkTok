package redis

import "strconv"

const (
	followSuffix   = "follow_set:"
	followerSuffix = "follower_set:"
)

func getFollowKey(userId int64) string {
	return followSuffix + "user:" + strconv.FormatInt(userId, 10)
}

func getFollowerKey(userId int64) string {
	return followerSuffix + "user:" + strconv.FormatInt(userId, 10)
}

// AddFollow 在userIdA的关注列表中添加userIdB
func AddFollow(userIdA, userIdB int64) {
	if rdb == nil {
		return
	}
	add(rdb, getFollowKey(userIdA), userIdB)
}

// DelFollow 在userIdA的关注列表中删除userIdB
func DelFollow(userIdA, userIdB int64) {
	del(rdb, getFollowKey(userIdA), userIdB)
}

// ExistFollow 检查userIdA的关注列表中是否存在userIdB
func ExistFollow(userIdA, userIdB int64) bool {
	if rdb == nil {
		return false
	}
	return exist(rdb, getFollowKey(userIdA), userIdB)
}

// CountFollow 获取userIdA的关注列表的长度
func CountFollow(userIdA int64) (int64, error) {
	return count(rdb, getFollowKey(userIdA))
}

// GetFollow 获取userIdA的关注列表
func GetFollow(userIdA int64) []int64 {
	return get(rdb, getFollowKey(userIdA))
}

// AddFollower 在userIdA的粉丝列表中添加userIdB
func AddFollower(userIdA, userIdB int64) {
	if rdb == nil {
		return
	}
	add(rdb, getFollowerKey(userIdA), userIdB)
}

// DelFollower 在userIdA的粉丝列表中删除userIdB
func DelFollower(userIdA, userIdB int64) {
	del(rdb, getFollowerKey(userIdA), userIdB)
}

// ExistFollower 检查userIdA的粉丝列表中是否存在userIdB
func ExistFollower(userIdA, userIdB int64) bool {
	if rdb == nil {
		return false
	}
	return exist(rdb, getFollowerKey(userIdA), userIdB)
}

// CountFollower 获取userIdA的粉丝列表的长度
func CountFollower(userIdA int64) (int64, error) {
	return count(rdb, getFollowerKey(userIdA))
}

// GetFollower 获取userIdA的粉丝列表
func GetFollower(userIdA int64) []int64 {
	return get(rdb, getFollowerKey(userIdA))
}

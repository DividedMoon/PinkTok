package redis

import (
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
	"video_service/biz/internal/constants"
)

//TODO expireTime从数据库中查到数据然后更新的时候会用到
var (
	expireTime    = time.Hour
	rdFavorite    *redis.Client
	rdComment     *redis.Client
	redisAddr     = constants.RedisAddr
	redisPassword = constants.RedisPassword
)

func InitRedis() {
	rdFavorite = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	rdComment = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
}

func countSetSize(c *redis.Client, k string) (sum int64, err error) {
	if sum, err = c.SCard(k).Result(); err != nil {
		return 0, err
	}
	return sum, nil
}

// checkSetExist check the relation k and v if checkSetExist
func checkSetExist(c *redis.Client, k string, v int64) (bool, error) {
	e, err := c.SIsMember(k, v).Result()
	if err != nil {
		return false, err
	} else {
		c.Expire(k, expireTime)
		return e, err
	}
}

//getStringValue 获取key对应的value
func getStringValue(c *redis.Client, k string) (int64, error) {
	v, err := c.Get(k).Result()
	if err != nil {
		return 0, err
	}
	vt, _ := strconv.ParseInt(v, 10, 64)
	return vt, nil
}

func initSet(c *redis.Client, k string, v []int64) error {
	redisValues := make([]interface{}, len(v))
	for _, val := range v {
		redisValues = append(redisValues, val)
	}

	tx := c.TxPipeline() // 高并发场景下使用管道优化插入操作
	tx.SAdd(k, redisValues...)
	tx.Expire(k, expireTime)
	_, err := tx.Exec()
	return err
}

func initString(c *redis.Client, k string, v int64) error {
	tx := c.TxPipeline()
	tx.Set(k, v, expireTime)
	_, err := tx.Exec()
	return err
}

//
//func addSet(c *redis.Client, k string, v int64) {
//	tx := c.TxPipeline()
//	tx.SAdd(k, v)
//	tx.Expire(k, expireTime)
//	_, _ = tx.Exec()
//}
//
//// del k & v
//func del(c *redis.Client, k string, v int64) {
//	tx := c.TxPipeline()
//	tx.SRem(k, v)
//	tx.Expire(k, expireTime)
//	_, _ = tx.Exec()
//}
//
//// check the set of k if checkSetExist
//func check(c *redis.Client, k string) bool {
//	if e, _ := c.Exists(k).Result(); e > 0 {
//		return true
//	}
//	return false
//}
//
//// checkSetExist check the relation k and v if checkSetExist
//func checkSetExist(c *redis.Client, k string, v int64) bool {
//	if e, _ := c.SIsMember(k, v).Result(); e {
//		c.Expire(k, expireTime)
//		return true
//	}
//	return false
//}
//
//func getStringValue(c *redis.Client, k string) (vt []int64) {
//	v, _ := c.SMembers(k).Result()
//	c.Expire(k, expireTime)
//	for _, vs := range v {
//		vI64, _ := strconv.ParseInt(vs, 10, 64)
//		vt = append(vt, vI64)
//	}
//	return vt
//}

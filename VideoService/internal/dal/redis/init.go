package redis

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
	"video_service/internal/constants"
)

// TODO expireTime从数据库中查到数据然后更新的时候会用到
var (
	expireTime    = time.Hour
	rdFavorite    *redis.Client
	rdComment     *redis.Client
	rdVideo       *redis.Client
	rdRobfig      *redis.Client
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
	rdVideo = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
	rdRobfig = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	//err := InitChangedVideoSet4Robfig()
	//if err != nil {
	//	hlog.Error("InitChangedVideoSet4Robfig error", err.Error())
	//	panic(err)
	//}
}
func CloseRedis() {
	err := rdVideo.Close()
	if err != nil {
		hlog.Error("rdVideo.Close() error", err.Error())
	}
	err = rdComment.Close()
	if err != nil {
		hlog.Error("rdComment.Close() error", err.Error())
	}
	err = rdFavorite.Close()
	if err != nil {
		hlog.Error("rdFavorite.Close() error", err.Error())
	}
	err = rdRobfig.Close()
	if err != nil {
		hlog.Error("rdRobfig.Close() error", err.Error())
	}

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

// getStringValue 获取key对应的value
func getStringValue(c *redis.Client, k string) (int64, error) {
	v, err := c.Get(k).Result()
	if err != nil {
		return 0, err
	}
	vt, _ := strconv.ParseInt(v, 10, 64)
	return vt, nil
}

func initSet(c *redis.Client, k string, v []int64) error {
	if len(v) == 0 {
		return fmt.Errorf("initSet v is empty")
	}
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

func initHash(c *redis.Client, k string, v map[string]interface{}) error {
	tx := c.TxPipeline()
	tx.HMSet(k, v)
	tx.Expire(k, expireTime)
	_, err := tx.Exec()
	return err
}

func getHash(c *redis.Client, k string) (map[string]string, error) {
	v, err := c.HGetAll(k).Result()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func getHashField(c *redis.Client, k string, filed string) (string, error) {
	v, err := c.HGet(k, filed).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func setHashField(c *redis.Client, k string, filed string, v interface{}) error {
	tx := c.TxPipeline()
	tx.HSet(k, filed, v)
	tx.Expire(k, expireTime)
	_, err := tx.Exec()
	return err
}

func querySetExist(c *redis.Client, k string) (bool, error) {
	if e, err := c.Exists(k).Result(); err != nil {
		return false, err
	} else {
		return e > 0, nil
	}
}

func getAllFromSetAndClear(c *redis.Client, k string) ([]int64, error) {
	elements := make([]int64, 0)
	tx := c.TxPipeline()
	members := tx.SMembers(k)
	tx.Del(k)
	_, err := tx.Exec()
	if err != nil {
		hlog.Error("getAllFromSetAndClear error", err.Error())
		return nil, err
	}
	for _, v := range members.Val() {
		vt, _ := strconv.ParseInt(v, 10, 64)
		elements = append(elements, vt)
	}
	return elements, nil
}

func addIntoSet(c *redis.Client, k string, v int64) error {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, expireTime)
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

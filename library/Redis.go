package library

import (
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

// 获取redis连接
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     ConfigParam.Redis.RedisHost + ":" + ConfigParam.Redis.RedisPort,
		Password: ConfigParam.Redis.RedisPassword,
		DB:       ConfigParam.Redis.RedisDb,
	})
	defer RedisClient.Close()
}

// 设置字符串
func SetRedisString(keys string, values interface{}, expiration time.Duration) {
	RedisClient.Set(keys, values, expiration)
}

// 获取字符串
func GetRedisString(name string) interface{} {
	return RedisClient.Get(name)
}

// 删除字符串
func DelRedisString(name string) {
	RedisClient.Del(name)
}

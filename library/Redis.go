package library

import "github.com/go-redis/redis"

// 获取redis连接
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     ConfigParam.Redis.RedisHost + ":" + ConfigParam.Redis.RedisPort,
		Password: ConfigParam.Redis.RedisPassword,
		DB:       ConfigParam.Redis.RedisDb,
	})
	return client
}

package library

import "github.com/go-redis/redis"


// redis配置信息
type Redis struct {
	RedisHost		string	`json:"redis_host"`		// redis地址
	RedisPort		string	`json:"redis_port"`		// redis端口
	RedisPassword	string	`json:"redis_password"`	// redis密码
	RedisDb			int		`json:"redis_db"`		// redis选择数据库
}

// 获取redis连接
func (redisObj Redis) RedisClient() *redis.Client {

	config, _ := config.ParseConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.RedisHost + ":" + config.Redis.RedisPort,
		Password: config.Redis.RedisPassword,
		DB:       config.Redis.RedisDb,
	})

	return client

}
package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func conn() *redis.Client {
	redisConfig := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	rdb := redis.NewClient(&redisConfig)
	return rdb
}

//Ping test connection
func Ping() string {
	RedisClient := conn()
	result, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	return result
}

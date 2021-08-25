package redicli

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var redisConfig = redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
}

var instance *redis.Client

//conn get redis connection
func conn() *redis.Client {
	if instance == nil {
		instance = redis.NewClient(&redisConfig)
	}
	return instance
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

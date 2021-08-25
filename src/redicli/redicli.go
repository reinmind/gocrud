package redicli

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

var redisConfig = redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
}

var instance *redis.Client

//init get redis connection
func init() {
	if instance == nil {
		instance = redis.NewClient(&redisConfig)
	}
}

//Ping test connection
func Ping() string {

	result, err := instance.Ping().Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	return result
}

func Set(key string, value interface{}) {
	instance.Set(key, value, 6*time.Minute)
}

func Get(key string) string {
	get := instance.Get(key)
	result, err := get.Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

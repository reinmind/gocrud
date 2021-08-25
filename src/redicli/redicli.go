package redicli

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redicli struct{}

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
func (Redicli) Ping() string {
	result, err := instance.Ping().Result()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return result
}

func (Redicli) Set(key string, value interface{}) {
	instance.Set(key, value, 6*time.Minute)
}

func (Redicli) Get(key string) string {
	get := instance.Get(key)
	result, err := get.Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

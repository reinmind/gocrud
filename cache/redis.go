package cache

import (
	_ "gocrud/config"
	"log"
	"time"

	"strconv"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

//Redicli redis client
type Redicli struct{}

var instance *redis.Client

//init get redis connection
func init() {
	configMap := viper.GetViper()
	var redisConfig = redis.Options{
		Addr:     configMap.GetString("redis.host") + ":" + strconv.Itoa(configMap.GetInt("redis.port")),
		Password: configMap.GetString("redis.password"),
		DB:       configMap.GetInt("redis.db"),
	}
	if instance == nil {
		instance = redis.NewClient(&redisConfig)
	}
}

//Ping test connection
func (Redicli) Ping() string {
	result, err := instance.Ping().Result()
	if err != nil {
		log.Fatalf("%v", err)
	}
	return result
}

//Set set key value
func (Redicli) Set(key string, value interface{}) {
	//set cache default expire time
	instance.Set(key, value, 6*time.Minute)
}

//Get get key value
func (Redicli) Get(key string) string {
	get := instance.Get(key)
	result, err := get.Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

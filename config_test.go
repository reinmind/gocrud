package main

import (
	"gocrud/cache"
	_ "gocrud/config"
	"gocrud/fs"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	s := viper.GetString("server.port")
	t.Logf(`port value: %v`, s)
}

func TestMinio(t *testing.T) {
	minioClient := fs.GetClient()
	//list all buckets
	bi, err := minioClient.ListBuckets(&gin.Context{})
	if err != nil {
		t.Errorf(`error: %v`, err)
	}
	for _, bucket := range bi {
		t.Logf(`bucket: %v`, bucket.Name)
	}
}

func TestRedis(t *testing.T) {
	redisClient := cache.Redicli{}
	s := redisClient.Ping()
	t.Logf(`ping: %v`, s)
	redisClient.Set("foo", "bar")
	redisClient.Set("hello", "world")

	s2 := redisClient.Get("foo")
	s3 := redisClient.Get("hello")

	if s2 == "bar" && s3 == "world" {
		t.Logf(`redis testing passed!`)
		return
	}
	t.Errorf(`redis testing failed!`)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrud/cache"
	"gocrud/db"
	"gocrud/entity"
	"gocrud/fs"
	"gocrud/rpc"
	"log"
	"net/http"
)

func main() {

	// 启动rpc服务器
	go rpc.StartServer()

	redi := cache.Redicli{}
	r := gin.Default()

	r.GET("/db/ping", func(c *gin.Context) {
		m1 := query()
		c.JSON(http.StatusOK, gin.H{
			"value": db.Response{Msg: "hello", Code: 2, Value: m1},
		})
	})

	ping := func(c *gin.Context) {
		message := redi.Ping()
		c.JSON(http.StatusOK, gin.H{
			"value": message,
		})
	}
	showContext := func(context *gin.Context) {
		log.Printf(`context.Keys: %v\n`, context.Keys)
		log.Printf(`context.Request: %v`, context.Request)

		context.JSON(http.StatusOK, gin.H{"URL": context.Request.URL,
			"requestURI": context.Request.RequestURI,
			"method":     context.Request.Method,
			"header":     context.Request.Header})
	}

	r.GET("/context/show", showContext)

	r.GET("/redis/ping", ping)

	r.GET("/ping/minio", func(context *gin.Context) {
		client := fs.GetClient()
		buckets, err := client.ListBuckets(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			context.JSON(200, gin.H{
				"buckets": buckets,
			})
		}
	})

	r.GET("/redis/test", func(ctx *gin.Context) {
		keyQuery := ctx.Query("keyQuery")
		valQuery := ctx.Query("valQuery")
		redi.Set(keyQuery, valQuery)
		valGot := redi.Get(keyQuery)
		success := valGot == valQuery
		ctx.JSON(http.StatusOK, gin.H{
			"success": success,
			"value":   valGot,
		})
	})

	// 启动gin服务
	err := r.Run(":8080")
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

}

//query get first tradeCount
func query() entity.TradeCount {
	connection := db.Conn()
	// resultList := map[string]entity.TradeCount{}
	result := entity.TradeCount{}
	err := connection.Table("entity_trade_count").Take(&result).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return result
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrud/dbcli"
	"gocrud/entity"
	"gocrud/miniocli"
	"gocrud/redicli"
	"gocrud/rpc"
	"net/http"
)

func main() {

	// 启动rpc服务器
	go rpc.StartServer()

	redi := redicli.Redicli{}
	r := gin.Default()

	r.GET("/db/ping", func(c *gin.Context) {
		m1 := query()
		c.JSON(http.StatusOK, gin.H{
			"value": dbcli.Response{Msg: "hello", Code: 2, Value: m1},
		})
	})

	r.GET("/redis/ping", func(c *gin.Context) {
		message := redi.Ping()
		c.JSON(http.StatusOK, gin.H{
			"value": message,
		})
	})

	r.GET("/ping/minio", func(context *gin.Context) {
		client := miniocli.Conn()
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
	connection := dbcli.Conn()
	// resultList := map[string]entity.TradeCount{}
	result := entity.TradeCount{}
	err := connection.Table("entity_trade_count").Take(&result).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return result
}

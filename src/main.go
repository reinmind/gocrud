package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrud/dbcli"
	"gocrud/entity"
	"gocrud/miniocli"
	"gocrud/redicli"
)

func main() {

	r := gin.Default()

	r.GET("/ping/db", func(c *gin.Context) {
		m1 := query()
		c.JSON(200, gin.H{
			"value": dbcli.Response{"hello", 2, m1},
		})
	})

	r.GET("/ping/redis", func(c *gin.Context) {
		message := redicli.Ping()
		c.JSON(200, gin.H{
			"value": message,
		})
	})

	r.GET("/ping/minio", func(context *gin.Context) {
		client := miniocli.New()
		buckets, err := client.ListBuckets(context)
		if err != nil {
			context.JSON(500, gin.H{
				"message": err,
			})
		} else {
			context.JSON(200, gin.H{
				"buckets": buckets,
			})
		}
	})

	err := r.Run()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

}

//query get first tradeCount
func query() entity.TradeCount {
	connection := dbcli.New()
	// resultList := map[string]entity.TradeCount{}
	result := entity.TradeCount{}
	err := connection.Table("entity_trade_count").Take(&result).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return result
}

package main

import (
	"crud-demo/src/database"
	"crud-demo/src/database/entity"

	"fmt"
	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/ping",func(c *gin.Context){
		m1 := query()
		c.JSON(200,gin.H{
			"value" : database.Response{"hello",2,m1},
		})
	})
	err := r.Run()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
// map[string]interface
func query() entity.TradeCount{
	connection := database.Connetct()

	// resultList := map[string]entity.TradeCount{}
	result := entity.TradeCount{}
	err := connection.Table("entity_trade_count").Take(&result).Error
	if err != nil {
		fmt.Printf("%v\n",err)
	}
	return result
}
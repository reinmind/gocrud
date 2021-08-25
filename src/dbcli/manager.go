package dbcli

import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "log"

var connection *gorm.DB

func init() {
	var err error
	url := "hyperchain:K5bernetes@tcp(rm-bp1db5ttsxnkk51p1vo.mysql.rds.aliyuncs.com:3306)/carbon_manager?charset=utf8&parseTime=True&loc=Local"
	connection, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("connection failed!")
	}
}

//Conn 获取数据库连接
func Conn() *gorm.DB {
	return connection
}

package database
import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "log"

func Connetct() * gorm.DB{
	url := "hyperchain:K5bernetes@tcp(rm-bp1db5ttsxnkk51p1vo.mysql.rds.aliyuncs.com:3306)/carbon_manager?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("connection failed!")
	}
	return open
}
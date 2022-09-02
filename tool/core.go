package tool

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func init() {
	mysqlInfo := Config{}.GetMysqlEnv()
	mysqlConfig := mysqlInfo.Account + ":" + mysqlInfo.Password + "@(" + mysqlInfo.Host + ":" + mysqlInfo.Port + ")/" + mysqlInfo.Db_name
	DB, err = gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

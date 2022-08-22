package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//安装 `go get -u github.com/jinzhu/gorm`
//mysql连接

//quick-start

func main() {
	dsn := "root:luoying123..@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	defer sqlDb.Close()

}

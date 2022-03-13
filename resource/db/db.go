package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化 MySQL 链接
func InitDB(user, password, host, port, dbName string) {
	log.Println("connecting MySQL ... ", host)
	//dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if mdb == nil {
		panic("failed to connect database")
	}
	log.Println("connected")
	db = mdb
	return
}

// GetDB 获取数据库链接实例
func GetDB() *gorm.DB {
	return db
}

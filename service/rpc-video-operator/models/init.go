package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := "root:root@tcp(192.168.1.28:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm init error: ", err)
		panic("fail to connect database")
	}
	DB = db
	err = db.AutoMigrate(&VideoInfo{})
	if err != nil {
		panic("fail to create table")
		return
	}
}

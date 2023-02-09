package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//var RDB = InitRedisDB()

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
		panic("fail to connect database")
	}
	DB = db
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		panic("fail to create table")
	}

	err = db.AutoMigrate(&Message{})
	if err != nil {
		panic("fail to create table")
	}

	err = db.AutoMigrate(&Relation{})
	if err != nil {
		panic("fail to create table")
	}
}

//func InitRedisDB() *redis.Client {
//	return redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//}
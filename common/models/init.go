package models

import (
	"log"
	"paigu1902/douyin/common/config"
	"strconv"
	"time"

	mc "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	mysqlConfig := mc.Config{
		User:                 config.C.Mysql.Username,
		Passwd:               config.C.Mysql.Password,
		Net:                  "tcp",
		Addr:                 config.C.Mysql.Host + ":" + strconv.Itoa(config.C.Mysql.Port),
		DBName:               config.C.Mysql.Dbname,
		Loc:                  time.Local,
		ParseTime:            true,
		Collation:            "utf8mb4_general_ci",
		AllowNativePasswords: true,
	}
	dsn := mysqlConfig.FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Println("gorm Init Error : ", err)
		panic("fail to connect database")
	}
	DB = db
	migrateDBTable(&UserInfo{})
	migrateDBTable(&Message{})
	migrateDBTable(&Relation{})
	migrateDBTable(&VideoInfo{})
	migrateDBTable(&UserFavo{})
	migrateDBTable(&UserComm{})
}

func migrateDBTable(dst ...interface{}) {
	err := DB.Set("gorm:table_options", "CHARSET=utf8mb4_general_ci").AutoMigrate(dst...)
	if err != nil {
		panic("fail to create table")
		return
	}
}

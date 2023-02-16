package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RDB = InitRedisDB()

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

var RdbVCid *redis.Client // video -> comment 一对多
var RdbCVid *redis.Client // comment -> video 一对一
var Ctx = context.Background()

func InitRedis() {
	RdbCVid = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "xx",
		DB:       1,
	})
	RdbVCid = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "xx",
		DB:       2,
	})
}

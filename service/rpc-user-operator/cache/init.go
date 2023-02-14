package cache

import (
	"github.com/go-redis/redis/v8"
)

var RdbFavoUser *redis.Client  //key:UserId,value:VideoId
var RdbFavoVideo *redis.Client //key:VideoId,value:UserId

func InitRedis() {
	RdbFavoUser = redis.NewClient(
		&redis.Options{
			//TODO
			Addr:     "0.0.0.0",
			Password: "",
			DB:       0,
		})
	RdbFavoVideo = redis.NewClient(
		&redis.Options{
			//TODO
			Addr:     "0.0.0.0",
			Password: "",
			DB:       0,
		})
}

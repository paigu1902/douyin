package cache

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func init() {
	opt, err := redis.ParseURL("redis://localhost:6379/<db>")
	if err != nil {
		panic(err)
	}
	RDB = redis.NewClient(opt)
}

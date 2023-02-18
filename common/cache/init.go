package cache

import (
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

var RdbUserOp *redis.Client

// video -> comment "VideoIdToComments:*" 一对多
// video -> user "VideoIdsToUserIdsIds:*" 多对多
// user -> video "UserIdsToVideoIds:*" 多对多

func init() {
	RdbUserOp = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
}

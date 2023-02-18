package cache

import (
	"github.com/go-redis/redis/v8"
	"paigu1902/douyin/common/config"
)

var RDB = InitRedisDB()

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Host + ":6379",
		Password: config.C.Redis.Password, // no password set
		DB:       config.C.Redis.Db,       // use default DB
	})
}

var RdbUserOp *redis.Client

// video -> comment "VideoIdToComments:*" 一对多
// video -> user "VideoIdsToUserIdsIds:*" 多对多
// user -> video "UserIdsToVideoIds:*" 多对多

func init() {
	RdbUserOp = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Host + ":6379",
		Password: config.C.Redis.Password, // no password set
		DB:       config.C.Redis.Db,       // use default DB
	})
}

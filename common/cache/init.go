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

// video -> comment "VideoIdToCommentIds:*" 一对多
// comment -> video "CommentIdToVideoId:*" 一对一

var RdbFavoUser *redis.Client  //key:UserId,value:VideoId
var RdbFavoVideo *redis.Client //key:VideoId,value:UserId

func init() {
	RdbUserOp = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	RdbFavoUser = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       3,
	})
	RdbFavoVideo = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       4,
	})
}

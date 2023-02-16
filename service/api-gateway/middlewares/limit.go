package middlewares

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"paigu1902/douyin/common/cache"
	"strconv"
	"time"
)

type LimitConfig struct {
	KeyPre     string
	Expiration time.Duration
	MaxNum     int
}

func LimitAction(c context.Context, ctx *app.RequestContext, config LimitConfig) {
	value, exists := ctx.Get("from_id")
	if exists != true {
		ctx.JSON(400, errors.New("token解析失败"))
	}
	fromId := value.(uint)
	redisKey := config.KeyPre + strconv.Itoa(int(fromId))
	i, err := cache.RDB.Get(c, redisKey).Int()
	if err == redis.Nil {
		cache.RDB.Set(c, redisKey, 1, config.Expiration)
		ctx.Next(c)
	} else if err != nil {
		hlog.Warn("redis缓存错误")
		ctx.Next(c)
	} else if i > config.MaxNum {
		ctx.AbortWithMsg("请求次数过多", 400)
		return
	} else {
		cache.RDB.Incr(c, redisKey)
		ctx.Next(c)
	}
}

func LimitFollowAction() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		LimitAction(c, ctx, LimitConfig{KeyPre: "FollowLim:", Expiration: time.Second * 10, MaxNum: 5})
	}
}

package middlewares

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"net/http"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/utils"
	"strconv"
)

type Auth struct {
	Token string `json:"token" form:"token"`
}

func AuthUserCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token, ok := c.GetQuery("token")
		if !ok {
			auth := new(Auth)
			err := c.Bind(auth)
			if err != nil {
				hlog.Warn("Unauthorized Authorization")
				c.AbortWithMsg("Unauthorized Authorization", http.StatusUnauthorized)
				return
			}
			token = auth.Token
		}
		userClaim, err := utils.AnalyseToken(token)
		if err != nil {
			hlog.Warn("Unauthorized Authorization")
			c.AbortWithMsg("Unauthorized Authorization", http.StatusUnauthorized)
			return
		}
		if userClaim == nil {
			c.AbortWithMsg("Unauthorized Authorization", http.StatusUnauthorized)
			return
		}

		forbiddenKey := "ForbiddenToken:" + strconv.Itoa(int(userClaim.ID))
		time, err := cache.RDB.Get(ctx, forbiddenKey).Int64()
		if err != nil && err != redis.Nil {
			hlog.Warn("Redis Error")
		} else if err != redis.Nil && time > userClaim.ExpiresAt {
			hlog.Warn("Unauthorized Authorization")
			c.AbortWithMsg("Unauthorized Authorization", http.StatusUnauthorized)
			return
		}
		c.Set("user_claims", userClaim)
		c.Set("from_id", userClaim.ID)
		c.Next(ctx)
	}
}

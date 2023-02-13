package middlewares

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"paigu1902/douyin/common/utils"
)

type Auth struct {
	Token string `json:"token"`
}

func AuthUserCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token, ok := c.GetQuery("token")
		if !ok {
			auth := new(Auth)
			err := c.Bind(auth)
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "Unauthorized Authorization",
				})
				return
			}
			token = auth.Token
		}
		userClaim, err := utils.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		c.Set("user_claims", userClaim)
		c.Next(ctx)
	}
}

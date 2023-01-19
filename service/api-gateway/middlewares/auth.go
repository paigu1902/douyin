package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paigu1902/douyin/common/utils"
)

type Auth struct {
	Token string `json:"token"`
}

func AuthUserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.GetQuery("token")
		if !ok {
			auth := new(Auth)
			err := c.BindJSON(auth)
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
		c.Next()
	}
}

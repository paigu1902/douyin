package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var myKey = []byte("paigu1902")

type UserClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken
// 生成 token
func GenerateToken(id uint, name string) (string, error) {
	UserClaim := &UserClaims{
		ID:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	// 1小时过期
	UserClaim.ExpiresAt = time.Now().Add(time.Second * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

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
		userClaim, err := AnalyseToken(token)
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

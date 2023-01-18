package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"paigu1902/douyin/middlewares"
	"paigu1902/douyin/models"
	"paigu1902/douyin/utils"
)

type RegisterResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     uint   `json:"user_id"`     // 用户id
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := new(models.UserInfo)
	user.UserName = username

	err := models.DB.Where(user).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user.Password = utils.MD5Encode(password)
		models.DB.Create(&user)
		token, err := middlewares.GenerateToken(user.ID, user.UserName)
		if err == nil {
			response := RegisterResponse{
				StatusCode: 0,
				StatusMsg:  "注册成功",
				Token:      token,
				UserID:     user.ID,
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := RegisterResponse{
				StatusCode: 1,
				StatusMsg:  "token生成失败",
				Token:      token,
				UserID:     user.ID,
			}
			c.JSON(http.StatusOK, response)
		}
	}
	response := RegisterResponse{
		StatusCode: -1,
		StatusMsg:  "注册失败",
	}
	c.JSON(http.StatusOK, response)
}

//func Login(c *gin.Context) {
//	username := c.Query("username")
//	password := c.Query("password")
//
//}

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"paigu1902/douyin/models"
	"paigu1902/douyin/utils"
)

type UserResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     uint   `json:"user_id"`     // 用户id
}

// Register
// Summary 用户注册
// @Tags         用户模块
// @param username query string false "用户名"
// @param password query string false "密码"
// @Success      200  {string}  json{"code","message"}
// @Router       /douyin/user/register [post]
func Register(c *gin.Context) {
	user := models.UserInfo{UserName: c.Query("username"), Password: c.Query("password")}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(user.Password, salt)
	user.Salt = salt
	err := models.CreateUser(user)
	if err != nil {
		response := UserResponse{
			StatusCode: -1,
			StatusMsg:  "用户名已经存在",
		}
		c.JSON(http.StatusOK, response)
		return
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err == nil {
		response := UserResponse{
			StatusCode: 0,
			StatusMsg:  "注册成功",
			Token:      token,
			UserID:     user.ID,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := UserResponse{
			StatusCode: 1,
			StatusMsg:  "token生成失败",
			Token:      token,
			UserID:     user.ID,
		}
		c.JSON(http.StatusOK, response)
	}
}

// Login
// Summary		用户登录
// @Tags         用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success      200  {string}  json{"code","message"}
// @Router       /douyin/user/login [post]
func Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.UserName == "" {
		response := UserResponse{
			StatusCode: -1,
			StatusMsg:  "用户名不存在",
		}
		c.JSON(http.StatusOK, response)
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		response := UserResponse{
			StatusCode: -1,
			StatusMsg:  "密码错误",
		}
		c.JSON(http.StatusOK, response)
		return
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err == nil {
		response := UserResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			Token:      token,
			UserID:     user.ID,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := UserResponse{
			StatusCode: 1,
			StatusMsg:  "token生成失败",
			Token:      token,
			UserID:     user.ID,
		}
		c.JSON(http.StatusOK, response)
	}
}

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"paigu1902/douyin/models"
	"paigu1902/douyin/utils"
)

// Register
// Summary 用户注册
// @Tags         用户模块
// @param username query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success      200  {string}  json{"code","message"}
// @Router       /douyin/user/register [post]
func Register(c *gin.Context) {
	user := models.UserInfo{}
	user.UserName = c.Query("username")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	err := models.CreateUser(user)
	fmt.Println(err)
	fmt.Println(err == nil)
	if err != nil {
		c.JSON(-1, gin.H{
			"message": "用户名已被使用",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "新增用户成功",
		"userID":  user.ID,
		"token":   12345,
	})
}

// Login
// Summary		用户登录
// @Tags         用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success      200  {string}  json{"code","message"}
// @Router       /douyin/user/login [post]
func Login(c *gin.Context) {
	data := models.UserInfo{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.UserName == "" {
		c.JSON(-1, gin.H{
			"code":    -1, //0 成功 -1失败
			"message": "查询为空",
			"data":    data,
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(-1, gin.H{
			"code":    -1, //0 成功 -1失败
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserBynameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, //0 成功 -1失败
		"message": "登陆成功",
		"data":    data,
	})
	return
}

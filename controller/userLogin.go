package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"paigu1902/douyin/models"
	"paigu1902/douyin/utils"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := new(models.UserInfo)
	user.UserName = username

	err := models.DB.Where(user).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user.Password = utils.CalMD5(password)
		models.DB.Create(&user)
		c.JSON(http.StatusOK, "ok")
	} else {
		c.JSON(http.StatusOK, "no ok")
	}

}

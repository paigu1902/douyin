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

	var user models.UserInfo
	err := models.DB.Where("user_name = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		models.DB.Create(models.UserInfo{UserName: username, Password: utils.CalMD5(password)})
	}

	c.JSON(http.StatusOK, "ok")
}

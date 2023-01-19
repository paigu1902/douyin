package router

import (
	"github.com/gin-gonic/gin"
	"paigu1902/douyin/service/api-gateway/controller"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

}

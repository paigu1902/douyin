package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"paigu1902/douyin/docs"
	"paigu1902/douyin/service"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", service.Register)
	apiRouter.POST("/user/login/", service.Login)

}

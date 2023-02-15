// Code generated by hertz generator.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	userInfo "paigu1902/douyin/service/api-gateway/biz/handler/userInfoPb"
	videoOperator "paigu1902/douyin/service/api-gateway/biz/handler/videoOperatorPb"
	"paigu1902/douyin/service/api-gateway/middlewares"
)

// customizeRegister registers customize routers.
func Register(r *server.Hertz) {
	v1 := r.Group("/v1")
	v1.POST("/login", userInfo.LoginMethod)
	v1.POST("/register", userInfo.RegisterMethod)
	v2 := r.Group("/v2")
	v2.Use(middlewares.AuthUserCheck())
	v2.GET("/info", userInfo.InfoMethod)

	publishGroup := r.Group("/douyin/publish")
	publishGroup.GET("/list", videoOperator.PublishListMethod)
	publishGroup.POST("/action", videoOperator.PublishActionMethod)

	publishGroup.POST("/action", videoOperator.UploadMethod)

	v4 := r.Group("/v4")
	v4.GET("/feed", videoOperator.FeedMethod)
}

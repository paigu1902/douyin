// Code generated by hertz generator.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	userComm "paigu1902/douyin/service/api-gateway/biz/handler/UserCommPb"
	userFavo "paigu1902/douyin/service/api-gateway/biz/handler/userFavoPb"
	userInfo "paigu1902/douyin/service/api-gateway/biz/handler/userInfoPb"
	userRelation "paigu1902/douyin/service/api-gateway/biz/handler/userRelationPb"
	videoOperator "paigu1902/douyin/service/api-gateway/biz/handler/videoOperatorPb"
	"paigu1902/douyin/service/api-gateway/middlewares"
)

// Register customizeRegister registers customize routers.
func Register(r *server.Hertz) {
	userinfoGroup := r.Group("/douyin/user")
	userinfoGroup.Use(middlewares.AccessLog())
	userinfoGroup.POST("/login/", userInfo.LoginMethod)
	userinfoGroup.POST("/register/", userInfo.RegisterMethod)
	userinfoGroup.Use(middlewares.AuthUserCheck())
	userinfoGroup.GET("", userInfo.InfoMethod)

	publishGroup := r.Group("/douyin/publish/")
	publishGroup.Use(middlewares.AccessLog())
	publishGroup.GET("/list/", videoOperator.PublishListMethod)
	publishGroup.Use(middlewares.AuthUserCheck())
	publishGroup.POST("/action/", videoOperator.PublishActionMethod)

	userMessageGroup := r.Group("/douyin/message/")
	userMessageGroup.Use(middlewares.AccessLog())
	userMessageGroup.Use(middlewares.AuthUserCheck())
	userMessageGroup.GET("/chat/", userRelation.MessageHistory)
	userMessageGroup.POST("/action/", userRelation.MessageAction)

	userRelationGroup := r.Group("/douyin/relation/")
	userRelationGroup.Use(middlewares.AccessLog())
	userRelationGroup.Use(middlewares.AuthUserCheck())
	userRelationGroup.POST("/action/", userRelation.FollowAction, middlewares.LimitFollowAction())
	userRelationGroup.GET("/follow/list/", userRelation.FollowList)
	userRelationGroup.GET("/follower/list/", userRelation.FollowerList)
	userRelationGroup.GET("/friend/list/", userRelation.FriendList)

	feedGroup := r.Group("/douyin")
	feedGroup.Use(middlewares.AccessLog())
	feedGroup.GET("/feed", videoOperator.FeedMethod)

	favoriteGroup := r.Group("/douyin/favorite/")
	favoriteGroup.Use(middlewares.AccessLog())
	favoriteGroup.GET("/list/", userFavo.FavoListMethod)
	favoriteGroup.Use(middlewares.AuthUserCheck())
	favoriteGroup.POST("/action/", userFavo.FavoActionMethod, middlewares.LimitFavorAction())

	userCommGroup := r.Group("/douyin/comment/")
	userCommGroup.Use(middlewares.AccessLog())
	userCommGroup.Use(middlewares.AuthUserCheck())
	userCommGroup.GET("/list/", userComm.CommentGetListMethod)
	userCommGroup.POST("/action/", userComm.CommentActionMethod, middlewares.LimitCommentAction())
}

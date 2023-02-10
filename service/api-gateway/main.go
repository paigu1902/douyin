// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	userInfo "paigu1902/douyin/service/api-gateway/biz/handler/userInfoPb"
)

func main() {
	// TODO:暂时没有研究hertz将网关也注册上去
	svr := server.New(
		server.WithHostPorts(":3002"),
		//server.WithRegistry(r),

	)
	v1 := svr.Group("/v1")
	v1.GET("/login", userInfo.LoginMethod)
	v1.GET("/register", userInfo.RegisterMethod)
	svr.Spin()
}

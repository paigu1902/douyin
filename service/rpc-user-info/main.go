package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}
	svr := userInfoPb.NewServer(
		new(UserInfoImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: 50051}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userInfoImpl"}),
		server.WithRegistry(r),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

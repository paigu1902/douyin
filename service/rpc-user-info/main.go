package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"paigu1902/douyin/common/nacos"
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
)

func main() {
	svr := userInfoPb.NewServer(
		new(UserInfoImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userInfoImpl"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
		server.WithServiceAddr(&net.TCPAddr{Port: 50051, IP: net.IPv4(127, 0, 0, 1)}),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

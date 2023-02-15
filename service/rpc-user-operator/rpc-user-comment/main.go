package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"paigu1902/douyin/common/nacos"
	UserCommPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb/usercommrpc"
)

func main() {
	svr := UserCommPb.NewServer(
		new(UserCommRpcImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: 12345, IP: net.IPv4(127, 0, 0, 1)}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "UserComment"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

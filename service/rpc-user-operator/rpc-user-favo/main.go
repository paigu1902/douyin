package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"paigu1902/douyin/common/nacos"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb/userfavorpc"
)

func main() {
	svr := userFavoPb.NewServer(
		new(UserFavoRpcImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "UserFavoImpl"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
		server.WithServiceAddr(&net.TCPAddr{Port: 50056}),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

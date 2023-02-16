package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"paigu1902/douyin/common/nacos"
	videoOperatorPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb/videooperator"
)

func main() {
	svr := videoOperatorPb.NewServer(
		new(VideoOperatorImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoOperatorImpl"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
		server.WithServiceAddr(&net.TCPAddr{Port: 50053, IP: net.IPv4(127, 0, 0, 1)}),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

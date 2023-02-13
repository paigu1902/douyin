package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	videoOperatorPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb/videooperator"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}
	svr := videoOperatorPb.NewServer(
		new(VideoOperatorImpl),
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

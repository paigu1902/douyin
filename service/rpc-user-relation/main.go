package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	UserRelationPb "paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userRelation"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}
	svr := UserRelationPb.NewServer(
		new(UserRelationImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: 50052}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userRelation"}),
		server.WithRegistry(r),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

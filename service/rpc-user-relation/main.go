package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"os"
	"paigu1902/douyin/common/nacos"
	UserRelationPb "paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
)

func main() {
	f, err := os.OpenFile("./user-relation.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)
	svr := UserRelationPb.NewServer(
		new(UserRelationImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: 50052}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userRelation"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

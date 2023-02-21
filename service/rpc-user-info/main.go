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
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
)

func main() {
	f, err := os.OpenFile("./user-info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)
	svr := userInfoPb.NewServer(
		new(UserInfoImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userInfoImpl"}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Cli)),
		server.WithServiceAddr(&net.TCPAddr{Port: 50051}),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

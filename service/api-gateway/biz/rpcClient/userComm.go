package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb/usercommrpc"
)

var UserComm usercommrpc.Client

func init() {
	c, err := usercommrpc.NewClient("UserCommRpcImpl", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	UserComm = c
}

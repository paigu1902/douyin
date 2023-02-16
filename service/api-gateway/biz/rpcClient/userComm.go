package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb/usercommrpc"
	"time"
)

var UserComm usercommrpc.Client

func init() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c := usercommrpc.MustNewClient(
		"UserCommRpcImpl",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*5),
	)
	UserComm = c
}

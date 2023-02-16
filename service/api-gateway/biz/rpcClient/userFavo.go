package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb/userfavorpc"
	"time"
)

var UserFavo userfavorpc.Client

func init() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c := userfavorpc.MustNewClient(
		"userFavoRpcImpl",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*5),
	)
	UserFavo = c
}

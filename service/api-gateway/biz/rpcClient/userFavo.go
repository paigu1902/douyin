package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb/userfavorpc"
)

var UserFavo userfavorpc.Client

func init() {
	c, err := userfavorpc.NewClient("UserFavoriteImpl", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	UserFavo = c
}

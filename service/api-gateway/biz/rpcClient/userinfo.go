package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
)

var UserInfo userinfo.Client

func init() {
	c, err := userinfo.NewClient("userInfoImpl", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	UserInfo = c
}

package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb/videooperator"
)

var VideoOpClient videooperator.Client

func init() {
	c, err := videooperator.NewClient("videoOperatorImpl", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	VideoOpClient = c
}

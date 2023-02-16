package rpcClient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb/videooperator"
	"time"
)

var VideoOperatorClient videooperator.Client

func init() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c := videooperator.MustNewClient(
		"videoOperatorImpl",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*5),
	)
	VideoOperatorClient = c
}

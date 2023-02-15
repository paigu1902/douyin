package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
)

var UserRelation userrelation.Client

func init() {
	c, err := userrelation.NewClient("userRelation", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	UserRelation = c
}

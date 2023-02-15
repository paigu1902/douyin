package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
)

func main() {
	c, err := userrelation.NewClient("userRelation", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	action, err := c.FollowAction(ctx, &userRelationPb.FollowActionReq{FromId: 11, ToId: 13, Type: "1"})
	defer ctx.Done()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(action.GetStatusMsg())
	fmt.Println(action.GetStatusCode())

}

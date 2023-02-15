package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/common/nacos"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
	"time"
)

func main() {
	c, err := userrelation.NewClient("userRelation", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer ctx.Done()
	//action, err := c.FollowAction(ctx, &userRelationPb.FollowActionReq{FromId: 19, ToId: 13, Type: "1"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(action.GetStatusMsg())
	//fmt.Println(action.GetStatusCode())

	follow, err := c.IsFollowList(ctx, &userRelationPb.IsFollowListReq{FromId: 12, ToId: []uint64{13, 11, 12}})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(follow.GetIsFollow())
	time.Sleep(time.Second)
}

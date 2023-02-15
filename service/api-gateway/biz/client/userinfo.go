package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
	"time"
)

var UserInfo userinfo.Client

func init() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c := userinfo.MustNewClient(
		"userInfoImpl",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)
	UserInfo = c
}

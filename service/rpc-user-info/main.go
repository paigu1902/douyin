package main

import (
	"log"
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
)

func main() {
	svr := userInfoPb.NewServer(new(UserInfoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

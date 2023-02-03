package main

import (
	"log"
	userRelationPb "paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
)

func main() {
	svr := userRelationPb.NewServer(new(UserRelationImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

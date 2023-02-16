package main

import (
	"log"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb/userfavorpc"
)

func main() {
	svr := userFavoPb.NewServer(new(UserFavoRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

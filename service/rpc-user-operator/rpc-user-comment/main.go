package main

import (
	"log"
	UserCommPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb/usercommrpc"
)

func main() {
	svr := UserCommPb.NewServer(new(UserCommRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

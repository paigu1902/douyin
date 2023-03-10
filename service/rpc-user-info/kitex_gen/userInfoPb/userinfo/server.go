// Code generated by Kitex v0.4.4. DO NOT EDIT.
package userinfo

import (
	server "github.com/cloudwego/kitex/server"
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler userInfoPb.UserInfo, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

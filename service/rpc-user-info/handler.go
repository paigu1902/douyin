package main

import (
	"context"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-info/logic"
)

// UserInfoImpl implements the last service interface defined in the IDL.
type UserInfoImpl struct{}

// Register implements the UserInfoImpl interface.
func (s *UserInfoImpl) Register(ctx context.Context, req *userInfoPb.RegisterReq) (resp *userInfoPb.RegisterResp, err error) {
	// TODO: Your code here...
	return logic.Register(req)
}

// Login implements the UserInfoImpl interface.
func (s *UserInfoImpl) Login(ctx context.Context, req *userInfoPb.LoginReq) (resp *userInfoPb.LoginResp, err error) {
	// TODO: Your code here...
	return logic.Login(req)
}

// Info implements the UserInfoImpl interface.
func (s *UserInfoImpl) Info(ctx context.Context, req *userInfoPb.UserInfoReq) (resp *userInfoPb.UserInfoResp, err error) {
	// TODO: Your code here...
	return logic.Info(ctx, req)
}

// ActionDB implements the UserInfoImpl interface.
func (s *UserInfoImpl) ActionDB(ctx context.Context, req *userInfoPb.ActionDBReq) (resp *userInfoPb.ActionDBResp, err error) {
	// TODO: Your code here...
	return logic.ActionDB(ctx, req)
}

// BatchInfo implements the UserInfoImpl interface.
func (s *UserInfoImpl) BatchInfo(ctx context.Context, req *userInfoPb.BatchUserReq) (resp *userInfoPb.BtachUserResp, err error) {
	// TODO: Your code here...
	return logic.BatchInfo(ctx, req)
}

// FavDB implements the UserInfoImpl interface.
func (s *UserInfoImpl) FavDB(ctx context.Context, req *userInfoPb.FavDBReq) (resp *userInfoPb.FavDBResp, err error) {
	// TODO: Your code here...
	return logic.FavDB(ctx, req)
}

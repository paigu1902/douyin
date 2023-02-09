package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"paigu1902/douyin/common/utils"
	userInfoPb "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-info/models"
)

// UserInfoImpl implements the last service interface defined in the IDL.
type UserInfoImpl struct{}

// Register implements the UserInfoImpl interface.
func (s *UserInfoImpl) Register(ctx context.Context, req *userInfoPb.RegisterReq) (resp *userInfoPb.RegisterResp, err error) {
	// TODO: Your code here...
	user := models.UserInfo{UserName: req.UserName, Password: req.Password}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(user.Password, salt)
	user.Salt = salt
	err = models.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return nil, err
	}
	return &userInfoPb.RegisterResp{StatusCode: 1, StatusMsg: "成功", UserId: uint64(user.ID), Token: token}, nil
}

// Login implements the UserInfoImpl interface.
func (s *UserInfoImpl) Login(ctx context.Context, req *userInfoPb.LoginReq) (resp *userInfoPb.LoginResp, err error) {
	// TODO: Your code here...
	name := req.UserName
	password := req.Password
	user := models.FindUserByName(name)
	isValid := utils.ValidPassword(password, user.Salt, user.Password)
	if !isValid {
		return nil, errors.New("密码错误")
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return nil, err
	}
	return &userInfoPb.LoginResp{StatusCode: 1, StatusMsg: "成功", UserId: uint64(int64(user.ID)), Token: token}, nil
}

// Info implements the UserInfoImpl interface.
func (s *UserInfoImpl) Info(ctx context.Context, req *userInfoPb.UserInfoReq) (resp *userInfoPb.UserInfoResp, err error) {
	// TODO: Your code here...
	ID := req.UserId
	//token := req.Token
	user := models.FindUserByID(ID)
	if user.UserName == "" {
		return nil, errors.New("用户不存在")
	}
	userDetail := &userInfoPb.User{UserId: uint64(user.ID), UserName: user.UserName, FollowCount: string(user.FollowCount), FollowerCount: string(user.FollowedCount), Is_Follow: "false"}
	return &userInfoPb.UserInfoResp{StatusCode: 1, StatusMsg: "成功", User: userDetail}, nil
}

// ActionDB implements the UserInfoImpl interface.
func (s *UserInfoImpl) ActionDB(ctx context.Context, req *userInfoPb.ActionDBReq) (resp *userInfoPb.ActionDBResp, err error) {
	// TODO: Your code here...
	fromId := req.FromId
	toId := req.ToId
	actionType := req.Type
	switch actionType {
	case 0:
		err = models.AddActcion(fromId, toId)
	case 1:
		err = models.DeleteActcion(fromId, toId)
	default:
		return nil, errors.New("用户操作异常")
	}
	if err != nil {
		return nil, err
	}
	return &userInfoPb.ActionDBResp{StatusCode: 1, StatusMsg: "成功"}, nil
}

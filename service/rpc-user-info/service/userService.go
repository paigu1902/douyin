package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-user-info/models"
	pb "paigu1902/douyin/service/rpc-user-info/userInfoPb"
)

type UserService struct {
	pb.UnimplementedUserInfoServer
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	user := models.UserInfo{UserName: req.UserName, Password: req.Password}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(user.Password, salt)
	user.Salt = salt
	err := models.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{StatusCode: 1, StatusMsg: "成功", UserId: uint64(user.ID), Token: token}, nil
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
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
	return &pb.LoginResp{StatusCode: 1, StatusMsg: "成功", UserId: uint64(int64(user.ID)), Token: token}, nil
}

func (s *UserService) Info(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	ID := req.UserId
	//token := req.Token
	user := models.FindUserByID(ID)
	if user.UserName == "" {
		return nil, errors.New("用户不存在")
	}
	userDetail := &pb.User{UserId: uint64(user.ID), UserName: user.UserName, FollowCount: string(user.FollowCount), FollowerCount: string(user.FollowedCount), IsFollow: "true"}
	return &pb.UserInfoResp{StatusCode: 1, StatusMsg: "成功", User: userDetail}, nil
}

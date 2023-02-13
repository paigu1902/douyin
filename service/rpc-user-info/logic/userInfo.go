package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb/userrelation"
	"strconv"
	"time"
)

func Login(req *userInfoPb.LoginReq) (resp *userInfoPb.LoginResp, err error) {
	resp = new(userInfoPb.LoginResp)
	user := models.UserInfo{}
	userReq := models.UserInfo{UserName: req.UserName, Password: req.Password}
	err = models.DB.Where("user_name = ?", userReq.UserName).First(&user).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "查询失败"
		return resp, err
	}
	isValid := utils.ValidPassword(userReq.Password, user.Salt, user.Password)
	if !isValid {
		resp.StatusCode = 1
		resp.StatusMsg = "密码错误"
		return resp, err
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "生成token失败"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	resp.UserId = uint64(int64(user.ID))
	resp.Token = token
	return resp, nil
}

func Register(req *userInfoPb.RegisterReq) (resp *userInfoPb.RegisterResp, err error) {
	resp = new(userInfoPb.RegisterResp)
	user := models.UserInfo{UserName: req.UserName, Password: req.Password}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(user.Password, salt)
	user.Salt = salt
	err = models.DB.Create(&user).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "注册失败"
		return resp, err
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "生成token失败"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	resp.UserId = uint64(int64(user.ID))
	resp.Token = token
	return resp, nil
}

func Info(ctx context.Context, req *userInfoPb.UserInfoReq) (resp *userInfoPb.UserInfoResp, err error) {
	resp = new(userInfoPb.UserInfoResp)
	ID := req.UserId
	token := req.Token
	user, isFollow, err := InfoRDB(ctx, ID, token)
	if err != nil || user.UserName == "" {
		resp.StatusCode = 1
		resp.StatusMsg = "查询错误"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	userDetail := &userInfoPb.User{UserId: uint64(user.ID), IsFollow: isFollow, UserName: user.UserName, FollowCount: user.FollowCount, FollowerCount: user.FollowedCount}
	resp.User = userDetail
	log.Println("user", resp)
	return resp, nil
}

func Actcion(fromId uint64, toId uint64, actionType string) error {
	var err error
	err = models.RDB.Del(context.Background(), strconv.Itoa(int(fromId))).Err()
	err = models.RDB.Del(context.Background(), strconv.Itoa(int(toId))).Err()
	if err != nil {
		return errors.New("删除缓存失败")
	}
	defer func() {
		go func() {
			time.Sleep(time.Second * 3)
			err = models.RDB.Del(context.Background(), strconv.Itoa(int(fromId))).Err()
			err = models.RDB.Del(context.Background(), strconv.Itoa(int(toId))).Err()
		}()
	}()

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&models.UserInfo{}).Where("id", fromId).Update("follow_count", gorm.Expr("follow_count"+actionType+"?", 1)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.Model(&models.UserInfo{}).Where("id", toId).Update("followed_count", gorm.Expr("followed_count"+actionType+"?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	return err
}

func InfoRDB(ctx context.Context, userId uint64, token string) (*models.UserInfo, bool, error) {
	var userinfo models.UserInfo
	var err error
	userClaim, _ := utils.AnalyseToken(token)
	fromId := uint64(userClaim.ID)
	// TODO：需要改造成nacos解析格式
	client, err := userrelation.NewClient("UserRelationImpl", client.WithHostPorts("0.0.0.0:8888"))
	isFollowResp, err := client.IsFollow(ctx, &userRelationPb.IsFollowReq{FromId: fromId, ToId: userId}, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		return &userinfo, false, err
	}
	userCache, err := models.RDB.Do(context.Background(), "get", userId).Text()
	if err == nil {
		err := json.Unmarshal([]byte(userCache), &userinfo)
		if err == nil {
			return &userinfo, isFollowResp.GetIsFollow(), nil
		}
	}
	err = models.DB.Where("id = ?", userId).First(&userinfo).Error
	u, _ := json.Marshal(userinfo)
	err = models.RDB.Do(ctx, "setex", userinfo.ID, 1000, string(u)).Err()
	if err != nil {
		return &userinfo, isFollowResp.GetIsFollow(), errors.New("写入异常")
	}
	return &userinfo, isFollowResp.GetIsFollow(), nil
}

func ActionDB(ctx context.Context, req *userInfoPb.ActionDBReq) (resp *userInfoPb.ActionDBResp, err error) {
	fromId := req.FromId
	toId := req.ToId
	actionType := req.Type
	switch actionType {
	case 0:
		err = Actcion(fromId, toId, "+")
	case 1:
		err = Actcion(fromId, toId, "-")
	default:
		return nil, errors.New("用户操作异常")
	}
	if err != nil {
		return nil, err
	}
	return &userInfoPb.ActionDBResp{StatusCode: 1, StatusMsg: "成功"}, nil
}

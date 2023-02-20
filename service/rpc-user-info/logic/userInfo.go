package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"gorm.io/gorm"
	"math/rand"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-user-info/client"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
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
	fromID := req.FromId
	toId := req.ToId
	user, isFollow, err := InfoRDB(ctx, fromID, toId)
	if err != nil || user.UserName == "" {
		resp.StatusCode = 1
		resp.StatusMsg = "查询错误"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	userDetail := &userInfoPb.User{UserId: uint64(user.ID), IsFollow: isFollow, UserName: user.UserName, FollowCount: user.FollowCount, FollowerCount: user.FollowedCount}
	resp.User = userDetail
	log.Info("resp", resp)
	return resp, nil
}

func Actcion(ctx context.Context, fromId uint64, toId uint64, actionType string) error {
	var err error
	err = cache.RDB.Del(ctx, "UserInfo:"+strconv.Itoa(int(fromId))).Err()
	if err != nil {
		return errors.New(strconv.Itoa(int(fromId)) + "删除缓存失败")
	}
	err = cache.RDB.Del(ctx, "UserInfo:"+strconv.Itoa(int(toId))).Err()
	if err != nil {
		return errors.New(strconv.Itoa(int(toId)) + "删除缓存失败")
	}
	defer func() {
		go func() {
			time.Sleep(time.Second * 3)
			cache.RDB.Del(ctx, "UserInfo:"+strconv.Itoa(int(fromId)))
			cache.RDB.Del(ctx, "UserInfo:"+strconv.Itoa(int(toId)))
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

func InfoRDB(ctx context.Context, fromId uint64, toId uint64) (*models.UserInfo, bool, error) {
	var userinfo models.UserInfo
	isFollowResp, err := client.UserRelation.IsFollow(ctx, &userRelationPb.IsFollowReq{FromId: fromId, ToId: toId})
	if err != nil {
		return nil, false, err
	}
	userCache, err := cache.RDB.Do(ctx, "get", "UserInfo:"+strconv.Itoa(int(toId))).Text()
	if err == nil {
		err = json.Unmarshal([]byte(userCache), &userinfo)
		if err == nil {
			return &userinfo, isFollowResp.GetIsFollow(), nil
		}
	}
	err = models.DB.Where("id = ?", toId).First(&userinfo).Error
	if err != nil {
		return nil, false, errors.New("用户不存在")
	}
	u, _ := json.Marshal(userinfo)
	err = cache.RDB.Do(ctx, "setex", "UserInfo:"+strconv.Itoa(int(userinfo.ID)), 1000, string(u)).Err()
	if err != nil {
		log.Warn("缓存写入异常")
	}
	return &userinfo, isFollowResp.GetIsFollow(), nil
}

func ActionDB(ctx context.Context, req *userInfoPb.ActionDBReq) (resp *userInfoPb.ActionDBResp, err error) {
	fromId := req.FromId
	toId := req.ToId
	actionType := req.Type
	switch actionType {
	case 0:
		err = Actcion(ctx, fromId, toId, "-")
	case 1:
		err = Actcion(ctx, fromId, toId, "+")
	default:
		return nil, errors.New("用户操作异常")
	}
	if err != nil {
		return nil, err
	}
	log.Info("操作成功", resp)
	return &userInfoPb.ActionDBResp{StatusCode: 1, StatusMsg: "成功"}, nil
}

func BatchInfo(ctx context.Context, req *userInfoPb.BatchUserReq) (resp *userInfoPb.BtachUserResp, err error) {
	var userinfo models.UserInfo
	var limitids []uint64
	var userinfors []*models.UserInfo
	countLimit := make(map[uint64]int, len(req.Batchids))
	respTmp := make(map[uint64]*userInfoPb.User, len(req.Batchids))
	resp = new(userInfoPb.BtachUserResp)
	isfollows := make(map[uint64]bool, len(req.Batchids))
	batchIds := req.Batchids
	if req.Fromid == 0 {
		for _, v := range batchIds {
			isfollows[v] = false
		}
	} else {
		isfollowresp, err := client.UserRelation.IsFollowList(ctx, &userRelationPb.IsFollowListReq{FromId: req.Fromid, ToId: batchIds})
		if err != nil {
			return nil, err
		}
		for i, v := range isfollowresp.GetIsFollow() {
			isfollows[uint64(i)] = v
		}
	}

	for _, id := range batchIds {
		u, err := cache.RDB.Get(ctx, "UserInfo:"+strconv.Itoa(int(id))).Result()
		if err == nil {
			err := json.Unmarshal([]byte(u), &userinfo)
			if err == nil {
				respTmp[id] = &userInfoPb.User{
					UserId:        id,
					UserName:      userinfo.UserName,
					FollowCount:   userinfo.FollowCount,
					FollowerCount: userinfo.FollowedCount,
					IsFollow:      isfollows[id],
				}
				continue
			}
		}
		limitids = append(limitids, id)
		countLimit[id] = 1
	}
	err = models.DB.Where("id IN ?", limitids).Find(&userinfors).Error
	for _, user := range userinfors {
		u, _ := json.Marshal(user)
		err = cache.RDB.Set(ctx, "UserInfo:"+strconv.Itoa(int(user.ID)), string(u), 1000).Err()
		if err != nil {
			log.Warn("写入缓存失败")
		}
		respTmp[uint64(user.ID)] = &userInfoPb.User{
			UserId:        uint64(user.ID),
			UserName:      user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowedCount,
			IsFollow:      isfollows[uint64(user.ID)],
		}
	}
	for _, id := range batchIds {
		resp.Batchusers = append(resp.Batchusers, respTmp[id])
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	log.Info("resp", resp)
	return resp, nil
}

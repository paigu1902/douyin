package logic

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	models2 "paigu1902/douyin/common/models"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
)

// rpc FollowAction(FollowActionReq) returns (FollowActionResp);
// rpc FollowList(FollowListReq) returns (FollowListResp);
// rpc FollowerList(FollowerListReq) returns (FollowerListResp);
// rpc FriendList (FriendListReq) returns (FriendListResp);

func followIds(id uint64) (ids []uint64, err error) {
	result := make([]models2.Relation, 0)
	err = models2.DB.Where(&models2.Relation{FromId: id}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	ids = make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.ToId
	}
	return ids, nil
}

func followerIds(id uint64) (ids []uint64, err error) {
	result := make([]models2.Relation, 0)
	err = models2.DB.Where(&models2.Relation{ToId: id}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	ids = make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.FromId
	}
	return ids, nil
}

func idsToMap(ids []uint64) map[uint64]struct{} {
	idMap := make(map[uint64]struct{}, len(ids))
	for _, v := range ids {
		idMap[v] = struct{}{}
	}
	return idMap
}

func isFollow(followMap map[uint64]struct{}, id uint64) bool {
	_, ok := followMap[id]
	return ok
}

func IsFollow(req *userRelationPb.IsFollowReq) (resp *userRelationPb.IsFollowResp, err error) {
	resp = new(userRelationPb.IsFollowResp)
	if req.FromId == req.ToId {
		resp.IsFollow = false
		return resp, nil
	}
	result := new(models2.Relation)

	err = models2.DB.Where(&models2.Relation{FromId: req.FromId, ToId: req.ToId}).First(result).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		resp.IsFollow = false
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	resp.IsFollow = true
	return resp, nil
}

func FollowAction(req *userRelationPb.FollowActionReq) (resp *userRelationPb.FollowActionResp, err error) {
	resp = new(userRelationPb.FollowActionResp)
	if req.FromId == req.ToId {
		resp.StatusCode = 1
		resp.StatusMsg = "关注失败"
		return resp, err
	}
	if req.Type == "1" {
		err = models2.DB.Create(&models2.Relation{FromId: req.FromId, ToId: req.ToId}).Error
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "关注失败"
			return resp, err
		}
	} else if req.Type == "0" {
		err = models2.DB.Where(&models2.Relation{FromId: req.FromId, ToId: req.ToId}).Delete(&models2.Relation{}).Error
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "取消关注失败"
			return resp, err
		}
	} else {
		resp.StatusCode = 1
		resp.StatusMsg = "未知操作"
		return resp, errors.New("unknown action type")
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	return resp, nil
}

func FollowList(req *userRelationPb.FollowListReq) (resp *userRelationPb.FollowListResp, err error) {
	resp = new(userRelationPb.FollowListResp)
	ids, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	userInfos := make([]models2.UserInfo, len(ids))
	err = models2.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(ids))
	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   fmt.Sprint(v.FollowCount),
			FollowerCount: fmt.Sprint(v.FollowedCount),
			IsFollow:      true,
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.UserList = userList
	return resp, nil
}

func FollowerList(req *userRelationPb.FollowerListReq) (resp *userRelationPb.FollowerListResp, err error) {
	resp = new(userRelationPb.FollowerListResp)
	ids, err := followerIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(ids))
	if len(ids) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "操作成功"
		resp.UserList = userList
		return resp, nil
	}
	userInfos := make([]models2.UserInfo, len(ids))
	err = models2.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	followIds, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	followMap := idsToMap(followIds)
	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   fmt.Sprint(v.FollowCount),
			FollowerCount: fmt.Sprint(v.FollowedCount),
			IsFollow:      isFollow(followMap, uint64(v.ID)),
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.UserList = userList
	return resp, nil
}

func FriendList(req *userRelationPb.FriendListReq) (resp *userRelationPb.FriendListResp, err error) {
	resp = new(userRelationPb.FriendListResp)
	followerIds, err := followerIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}

	followIds, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	followMap := idsToMap(followIds)

	friendIds := make([]uint64, 0)
	for _, v := range followerIds {
		if isFollow(followMap, v) {
			friendIds = append(friendIds, v)
		}
	}

	userInfos := make([]models2.UserInfo, len(friendIds))
	err = models2.DB.Where(&friendIds).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}

	userList := make([]*userRelationPb.User, len(friendIds))
	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   fmt.Sprint(v.FollowCount),
			FollowerCount: fmt.Sprint(v.FollowedCount),
			IsFollow:      true,
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.UserList = userList
	return resp, nil
}
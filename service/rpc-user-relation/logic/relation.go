package logic

import (
	"fmt"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"paigu1902/douyin/service/rpc-user-relation/models"
)

// rpc FollowAction(FollowActionReq) returns (FollowActionResp);
// rpc FollowList(FollowListReq) returns (FollowListResp);
// rpc FollowerList(FollowerListReq) returns (FollowerListResp);
// rpc FriendList (FriendListReq) returns (FriendListResp);

func followIds(id uint64) (ids []uint64, err error) {
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{From_id: id}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	ids = make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.To_id
	}
	return ids, nil
}

func followerIds(id uint64) (ids []uint64, err error) {
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{To_id: id}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	ids = make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.From_id
	}
	return ids, nil
}

func isFollow(followMap map[uint64]struct{}, id uint64) bool {
	_, ok := followMap[id]
	return ok
}

func FollowAction(req *userRelationPb.FollowActionReq) (resp *userRelationPb.FollowActionResp, err error) {
	resp = new(userRelationPb.FollowActionResp)
	if req.FromId == req.ToId {
		resp.StatusCode = 1
		resp.StatusMsg = "关注失败"
		return resp, err
	}
	if req.Type == "1" {
		err = models.DB.Create(&models.Relation{From_id: req.FromId, To_id: req.ToId}).Error
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "关注失败"
			return resp, err
		}
	} else {
		err = models.DB.Where(&models.Relation{From_id: req.FromId, To_id: req.ToId}).Delete(&models.Relation{}).Error
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "取消关注失败"
			return resp, err
		}
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
	userInfos := make([]models.UserInfo, len(ids))
	err = models.DB.Where(&ids).Find(&userInfos).Error
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
	userInfos := make([]models.UserInfo, len(ids))
	err = models.DB.Where(&ids).Find(&userInfos).Error
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
	followMap := make(map[uint64]struct{}, len(followIds))
	for _, v := range followIds {
		followMap[v] = struct{}{}
	}
	userList := make([]*userRelationPb.User, len(ids))
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
	followMap := make(map[uint64]struct{}, len(followIds))

	friendIds := make([]uint64, 0)
	for _, v := range followerIds {
		if isFollow(followMap, v) {
			friendIds = append(friendIds, v)
		}
	}

	userInfos := make([]models.UserInfo, len(friendIds))
	err = models.DB.Where(&friendIds).Find(&userInfos).Error
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

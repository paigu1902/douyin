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
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{From_id: req.UserId}).Find(&result).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	ids := make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.To_id
	}
	userInfos := make([]models.UserInfo, len(result))
	err = models.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(result))
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
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{To_id: req.UserId}).Find(&result).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	ids := make([]uint64, len(result))
	for i, v := range result {
		ids[i] = v.From_id
	}
	userInfos := make([]models.UserInfo, len(result))
	err = models.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "获取失败"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(result))
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

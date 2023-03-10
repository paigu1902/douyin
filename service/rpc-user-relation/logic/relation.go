package logic

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-relation/client"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"strconv"
	"time"
)

func followIds(id uint64) (ids []uint64, err error) {
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{FromId: id}).Find(&result).Error
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
	result := make([]models.Relation, 0)
	err = models.DB.Where(&models.Relation{ToId: id}).Find(&result).Error
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

func IsFollow(ctx context.Context, req *userRelationPb.IsFollowReq) (resp *userRelationPb.IsFollowResp, err error) {
	resp = new(userRelationPb.IsFollowResp)
	if req.FromId == req.ToId {
		resp.IsFollow = false
		return resp, nil
	}
	result := new(models.Relation)
	redisKey := "UserRelation:" + strconv.FormatUint(req.FromId, 10) + "-" + strconv.FormatUint(req.ToId, 10)
	text, err := cache.RDB.Get(ctx, redisKey).Result()
	if err == nil {
		if text == "1" {
			resp.IsFollow = true
			return resp, nil
		} else {
			resp.IsFollow = false
			return resp, nil
		}
	}
	err = models.DB.Where(&models.Relation{FromId: req.FromId, ToId: req.ToId}).First(result).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = cache.RDB.Set(ctx, redisKey, "0", time.Second*1000).Err()
		if err != nil {
			log.Warn("???????????????")
		}
		resp.IsFollow = false
		return resp, nil
	} else if err != nil {
		return nil, err
	}

	err = cache.RDB.Do(ctx, "setex", redisKey, 1000, "1").Err()
	if err != nil {
		log.Warn("???????????????")
	}
	resp.IsFollow = true
	return resp, nil
}

func IsFollowList(ctx context.Context, req *userRelationPb.IsFollowListReq) (resp *userRelationPb.IsFollowListResp, err error) {
	resp = new(userRelationPb.IsFollowListResp)
	relationSet := make(map[uint64]struct{}, 0)
	idInDb := make([]uint64, 0)
	for _, v := range req.ToId {
		redisKey := "UserRelation:" + strconv.FormatUint(req.FromId, 10) + "-" + strconv.FormatUint(v, 10)
		text, err := cache.RDB.Get(ctx, redisKey).Result()
		if err == nil {
			if text == "1" {
				relationSet[v] = struct{}{}
			}
		} else if err == redis.Nil {
			idInDb = append(idInDb, v)
		} else {
			log.Warn("??????????????????")
		}
	}

	var relations []models.Relation
	err = models.DB.Where("from_id = ?", req.FromId).Where("to_id IN ?", idInDb).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	for _, v := range relations {
		relationSet[v.ToId] = struct{}{}
		redisKey := "UserRelation:" + strconv.FormatUint(req.FromId, 10) + "-" + strconv.FormatUint(v.ToId, 10)
		err = cache.RDB.Set(ctx, redisKey, "1", time.Second*1000).Err()
		if err != nil {
			log.Warn("??????????????????")
		}
	}

	for _, v := range idInDb {
		if _, ok := relationSet[v]; ok == false {
			redisKey := "UserRelation:" + strconv.FormatUint(req.FromId, 10) + "-" + strconv.FormatUint(v, 10)
			err = cache.RDB.Set(ctx, redisKey, "0", time.Second*1000).Err()
			if err != nil {
				log.Warn("??????????????????")
			}
		}
	}

	res := make([]bool, len(req.ToId))
	for i, v := range req.ToId {
		if _, ok := relationSet[v]; ok == true {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	resp.IsFollow = res
	return resp, nil
}

func FollowAction(ctx context.Context, req *userRelationPb.FollowActionReq) (resp *userRelationPb.FollowActionResp, err error) {
	resp = new(userRelationPb.FollowActionResp)
	if req.FromId == req.ToId {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	redisKey := "UserRelation:" + strconv.FormatUint(req.FromId, 10) + "-" + strconv.FormatUint(req.ToId, 10)
	cache.RDB.Del(ctx, redisKey)
	defer func() {
		go func() {
			time.Sleep(time.Second)
			cache.RDB.Del(ctx, redisKey)
		}()
	}()
	if req.Type == "1" {
		err = models.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&models.Relation{FromId: req.FromId, ToId: req.ToId}).Error; err != nil {
				return err
			}
			_, err2 := client.UserInfoClient.ActionDB(ctx, &userInfoPb.ActionDBReq{FromId: req.FromId, ToId: req.ToId, Type: 1})
			if err2 != nil {
				return err2
			}
			return nil
		})
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "????????????"
			return resp, err
		}

	} else if req.Type == "2" {
		err = models.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where(&models.Relation{FromId: req.FromId, ToId: req.ToId}).Delete(&models.Relation{}).Error; err != nil {
				return err
			}
			_, err2 := client.UserInfoClient.ActionDB(ctx, &userInfoPb.ActionDBReq{FromId: req.FromId, ToId: req.ToId, Type: 0})
			if err2 != nil {
				return err2
			}
			return nil
		})
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "??????????????????"
			return resp, err
		}
	} else {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, errors.New("unknown action type")
	}
	resp.StatusCode = 0
	resp.StatusMsg = "????????????"
	return resp, nil
}

func FollowList(req *userRelationPb.FollowListReq) (resp *userRelationPb.FollowListResp, err error) {
	resp = new(userRelationPb.FollowListResp)
	ids, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(ids))
	if len(ids) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "????????????"
		resp.UserList = userList
		return resp, nil
	}
	userInfos := make([]models.UserInfo, len(ids))
	err = models.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowedCount,
			IsFollow:      true,
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "????????????"
	resp.UserList = userList
	return resp, nil
}

func FollowerList(req *userRelationPb.FollowerListReq) (resp *userRelationPb.FollowerListResp, err error) {
	resp = new(userRelationPb.FollowerListResp)
	ids, err := followerIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	userList := make([]*userRelationPb.User, len(ids))
	if len(ids) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "????????????"
		resp.UserList = userList
		return resp, nil
	}
	userInfos := make([]models.UserInfo, len(ids))
	err = models.DB.Where(&ids).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	followIds, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	followMap := idsToMap(followIds)
	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowedCount,
			IsFollow:      isFollow(followMap, uint64(v.ID)),
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "????????????"
	resp.UserList = userList
	return resp, nil
}

func FriendList(req *userRelationPb.FriendListReq) (resp *userRelationPb.FriendListResp, err error) {
	resp = new(userRelationPb.FriendListResp)
	followerIds, err := followerIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}

	followIds, err := followIds(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}
	followMap := idsToMap(followIds)

	friendIds := make([]uint64, 0)
	for _, v := range followerIds {
		if isFollow(followMap, v) {
			friendIds = append(friendIds, v)
		}
	}
	userList := make([]*userRelationPb.User, len(friendIds))

	userInfos := make([]models.UserInfo, len(friendIds))
	if len(friendIds) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "????????????"
		resp.UserList = userList
		return resp, nil
	}

	err = models.DB.Where(&friendIds).Find(&userInfos).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "????????????"
		return resp, err
	}

	for i, v := range userInfos {
		userList[i] = &userRelationPb.User{
			UserId:        uint64(v.ID),
			UserName:      v.UserName,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowedCount,
			IsFollow:      true,
		}
	}
	resp.StatusCode = 0
	resp.StatusMsg = "????????????"
	resp.UserList = userList
	return resp, nil
}

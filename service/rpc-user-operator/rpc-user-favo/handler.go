package main

import (
	"context"
	"errors"
	"log"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/common/rabbitmq"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
	VideoOptPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
	"strconv"
	"strings"
	"time"
)

// UserFavoRpcImpl implements the last service interface defined in the IDL.
type UserFavoRpcImpl struct{}

// FavoAction implements the UserFavoRpcImpl interface. 执行点赞、取消赞操作 actionType==1->点赞 actionType==2->取消点赞
func (s *UserFavoRpcImpl) FavoAction(ctx context.Context, req *userFavoPb.FavoActionReq) (resp *userFavoPb.FavoActionResp, err error) {
	// TODO: Your code here..
	if req.Type != 1 && req.Type != 2 {
		return &userFavoPb.FavoActionResp{
			StatusCode: 1,
			StatusMsg:  "Failed",
		}, errors.New("FavoAction Parameters Error")
	}
	user := strconv.FormatInt(req.UserId, 10)
	video := strconv.FormatInt(req.VideoId, 10)
	err1 := s.userFavoActionImpl(ctx, int(req.Type), user, video)
	err2 := s.videoFavoActionImpl(ctx, int(req.Type), user, video)
	if err1 != nil || err2 != nil {
		return &userFavoPb.FavoActionResp{
				StatusCode: 1,
				StatusMsg:  "Failed",
			},
			errors.New("FavoAction Failed")
	}
	return &userFavoPb.FavoActionResp{
			StatusCode: 0,
			StatusMsg:  "Succeed",
		},
		nil
}

// FavoList implements the UserFavoRpcImpl interface. 获取用户的点赞视频列表
func (s *UserFavoRpcImpl) FavoList(ctx context.Context, req *userFavoPb.FavoListReq) (resp *userFavoPb.FavoListResp, err error) {
	// TODO: Your code here...
	user := strconv.FormatInt(req.UserId, 10)
	key := "UserIdsToVideoIds" + user
	// 1. 查询cache
	ext, err1 := cache.RDB.Exists(ctx, key).Result()
	if err1 != nil {
		log.Println("FavoList Exists Error")
	}
	// 2. cache中存在点赞用户信息 获取视频列表
	if ext > 0 {
		videoIdListStr, err2 := cache.RDB.SMembers(ctx, key).Result()
		if err2 != nil {
			log.Println("FavoList SMembers Error")
		}
		var videoIdList []uint64
		for index, str := range videoIdListStr {
			id, _ := strconv.Atoi(str)
			videoIdList[index] = uint64(id)
		}
		// 获取videoOperator客户端 读取视频对象
		resp1, err3 := s.GetVideoList(ctx, videoIdList)
		if err3 != nil {
			return &userFavoPb.FavoListResp{
					StatusCode: 1,
					StatusMsg:  "Failed",
					VideoList:  nil},
				errors.New("FavoList GetVidoeList Failed")
		}
		return resp1, nil
	}
	// 3. cache中不存在用户信息 查询MySQL加入原有视频信息后更新
	err4 := s.readRecordsToCache(ctx, 1, key, user, "")
	if err4 != nil {
		log.Println("FavoList readRecordsToCache Error")
	}
	// 4. 重新读取缓存中的video列表
	videoIdListStr, err5 := cache.RDB.SMembers(ctx, key).Result()
	if err5 != nil {
		log.Println("FavoList SMembers Error")
	}
	var videoIdList []uint64
	for index, str := range videoIdListStr {
		id, _ := strconv.Atoi(str)
		videoIdList[index] = uint64(id)
	}
	// 5. 获取videoOperator客户端 读取视频对象
	resp2, err := s.GetVideoList(ctx, videoIdList)
	if err != nil {
		return &userFavoPb.FavoListResp{
				StatusCode: 1,
				StatusMsg:  "Failed",
				VideoList:  nil},
			errors.New("FavoList GetVidoeList Failed")
	}
	return resp2, nil
}

// FavoStatus 查询用户对某条视频的点赞状态
func (s *UserFavoRpcImpl) FavoStatus(ctx context.Context, req *userFavoPb.FavoStatusReq) (resp *userFavoPb.FavoStatusResp, err error) {
	user := strconv.FormatInt(req.UserId, 10)
	video := strconv.FormatInt(req.VideoId, 10)
	keyU := "UserIdsToVideoIds" + user
	keyV := "VideoIdsToUserIds" + video
	// 1. 查询RDB(key:user, value:video)
	ext1, err1 := cache.RDB.Exists(ctx, keyU).Result()
	if err1 != nil {
		log.Println("FavoStatus Exists Error")
	}
	if ext1 > 0 {
		res, err := cache.RDB.SIsMember(ctx, keyU, video).Result()
		if err != nil {
			log.Println("FavoStatus Exists Error")
		}
		return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, nil
	}
	// 2. 若RDB(key:user, value:video)中不存在点赞记录 查询RDB(key:video, value:user)
	ext2, err2 := cache.RDB.Exists(ctx, keyV).Result()
	if err2 != nil {
		log.Println("FavoStatus Exists Error")
	}
	if ext2 > 0 {
		res, err := cache.RDB.SIsMember(ctx, keyV, user).Result()
		if err != nil {
			log.Println("FavoStatus SIsMember Error")
		}
		return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, nil
	}
	//3. 若cache中不存在点赞记录 查询MySQL加入原有视频信息后更新 仅更新RDB(key:user, value:video)
	err3 := s.readRecordsToCache(ctx, 1, keyU, user, video)
	if err3 != nil {
		log.Println("FavoStatus readRecordsToCache Error")
	}
	// 4. 再次查询cache
	res, err := cache.RDB.SIsMember(ctx, keyU, video).Result()
	if err != nil {
		log.Println("FavoStatus SIsMember Error")
	}
	return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, nil
}

// FavoCount 查询视频被点赞总数
func (s *UserFavoRpcImpl) FavoCount(ctx context.Context, videoId int64) (int64, error) {
	video := strconv.FormatInt(videoId, 10)
	key := "VideoIdsToUserIds" + video
	// 1. 查询RDB(key:video, value:user)
	ext, err1 := cache.RDB.Exists(ctx, key).Result()
	if err1 != nil {
		log.Println("FavoCount Exists Error")
	}
	if ext > 0 {
		res, err2 := cache.RDB.SCard(ctx, key).Result()
		if err2 != nil {
			log.Println("FavoCount SCard Error")
		}
		return res - 1, nil //减去 default value
	}
	// 2. 若cache中不存在点赞记录 查询MySQL加入原有视频信息后更新 仅更新RDB(key:video, value:user)
	err3 := s.readRecordsToCache(ctx, 2, key, "", video)
	if err3 != nil {
		log.Println("FavoCount readRecordsToCache Error")
	}
	// 3. 再次查询cache
	res, err4 := cache.RDB.SCard(ctx, key).Result()
	if err4 != nil {
		log.Println("FavoCount SCard Error")
	}
	return res - 1, nil //减去 default value
}

// userFavoActionImpl 根据点赞类型操作RDB(key:userId, value:videoId)
func (s *UserFavoRpcImpl) userFavoActionImpl(ctx context.Context, actionType int, user string, video string) error {
	msg := strings.Builder{}
	msg.WriteString(user)
	msg.WriteString(" ")
	msg.WriteString(video)
	key := "UserIdsToVideoIds" + user
	// 1. 查询cache
	ext, err0 := cache.RDB.Exists(ctx, key).Result()
	if err0 != nil {
		log.Println("FavoAction Rdb.Exist Error")
	}
	// 2. 若cache中存在点赞用户信息 更新cache和数据库
	if ext > 0 {
		if actionType == 1 { //点赞
			_, err1 := cache.RDB.SAdd(ctx, key, video).Result()
			if err1 != nil {
				log.Println("FavoAction Rdb.SAdd Error")
			}
			rabbitmq.RmqFavoAdd.Publish(msg.String())
		} else { // 取消点赞
			_, err2 := cache.RDB.SRem(ctx, key, video).Result()
			if err2 != nil {
				log.Println("FavoAction Rdb.SRem Error")
			}
			rabbitmq.RmqFavoDel.Publish(msg.String())
		}
		return nil
	}
	// 3. 若cache中不存在用户信息 查询MySQL加入原有视频信息后更新RDB(key:user, value:video)
	err3 := s.readRecordsToCache(ctx, 1, key, user, video)
	if err3 != nil {
		log.Println("userFavoActionImpl readRecordsToCache Error")
	}
	// 4. 将本次点赞操作更新到cache和数据库
	if actionType == 1 { // 点赞
		_, err4 := cache.RDB.SAdd(ctx, key, video).Result()
		if err4 != nil {
			cache.RDB.Del(ctx, key)
			log.Println("FavoAction Rdb.SAdd Error")
		}
	} else if actionType == 2 { // 取消点赞
		_, err5 := cache.RDB.SRem(ctx, key, video).Result()
		if err5 != nil {
			cache.RDB.Del(ctx, key)
			log.Println("FavoAction Rdb.SRem Error")
		}
	} else {
		return errors.New("userFavoActionImpl Parameters Error")
	}
	rabbitmq.RmqFavoAdd.Publish(msg.String())
	return nil
}

// videoFavoActionImpl 根据点赞类型操作RDB(key:videoId, value:userId) 仅当取消点赞时更新
func (s *UserFavoRpcImpl) videoFavoActionImpl(ctx context.Context, actionType int, user string, video string) error {
	if actionType == 1 {
		return nil
	} else if actionType != 2 {
		return errors.New("videoFavoActionImpl Parameters Error")
	}
	key := "VideoIdsToUserIds" + video
	// 1. 查询cache
	ext, err0 := cache.RDB.Exists(ctx, key).Result()
	if err0 != nil {
		log.Println("FavoAction Rdb.Exists Error")
	}
	// 2. 若cache中存在点赞用户信息 更新cache
	if ext > 0 {
		_, err1 := cache.RDB.SRem(ctx, key, user).Result()
		if err1 != nil {
			log.Println("FavoAction Rdb.SRem Error")
		}
		return nil
	}
	// 3. 若cache中不存在用户信息 查询MySQL加入原有视频信息后更新RDB(key:video, value:user)
	err2 := s.readRecordsToCache(ctx, 2, key, user, video)
	if err2 != nil {
		log.Println("userFavoActionImpl readRecordsToCache Error")
	}
	// 4. 将本次点赞操作更新到cache
	_, err3 := cache.RDB.SRem(ctx, key, user).Result()
	if err3 != nil {
		cache.RDB.Del(ctx, key)
		log.Println("FavoAction GetFavoUserId Error")
	}
	return nil
}

// readRecordsToCache 查询MySQL加入原有视频信息后更新RDB redType==1->更新UserIdsToVideoIds readType==2->更新VideoIdsToUserIds
func (s *UserFavoRpcImpl) readRecordsToCache(ctx context.Context, readType int, key string, user string, video string) error {
	// 1. 加入Default Value 防止脏数据
	_, err1 := cache.RDB.SAdd(ctx, key, -1).Result()
	if err1 != nil {
		cache.RDB.Del(ctx, key)
		log.Println("FavoAction Rdb.SAdd Error")
	}
	// 4. 设置数据有效期
	_, err2 := cache.RDB.Expire(ctx, key, time.Duration(30)*time.Second).Result()
	if err2 != nil {
		cache.RDB.Del(ctx, key)
		log.Println("FavoAction Rdb.Expire Error")
	}
	// 5. 查询MySQL原有信息 加入cache
	if readType == 1 { //读取视频信息 更新UserIdsToVideoIds
		userId, _ := strconv.Atoi(user)
		videoIdList, err3 := models.GetFavoVideoId(uint64(userId))
		if err3 != nil {
			log.Println("FavoAction GetFavoVideoId Error")
		}
		for _, favoVideoId := range videoIdList {
			_, err := cache.RDB.SAdd(ctx, key, favoVideoId).Result()
			if err != nil {
				cache.RDB.Del(ctx, key)
				log.Println("FavoAction Rdb.SAdd Error")
			}
		}
	} else if readType == 2 { // 读取用户信息 更新VideoIdsToUserIds
		videoId, _ := strconv.Atoi(video)
		UserIdList, err4 := models.GetFavoUserId(uint64(videoId))
		if err4 != nil {
			log.Println("FavoAction GetFavoUserId Error")
		}
		for _, favoUserId := range UserIdList {
			_, err := cache.RDB.SAdd(ctx, key, favoUserId).Result()
			if err != nil {
				cache.RDB.Del(ctx, key)
				log.Println("FavoAction SAdd Error")
			}
		}
	} else {
		return errors.New("readRecordsToCache Parameters Error")
	}
	return nil
}

// GetVideoList 获取videoOperator客户端 读取视频对象
func (s *UserFavoRpcImpl) GetVideoList(ctx context.Context, videoIdList []uint64) (resp *userFavoPb.FavoListResp, err error) {
	myResp, err6 := rpcClient.VideoOpClient.VideoList(ctx, &VideoOptPb.VideoListReq{VideoId: videoIdList})
	if err6 != nil {
		return &userFavoPb.FavoListResp{
				StatusCode: 1,
				StatusMsg:  "Failed",
				VideoList:  nil},
			errors.New("FavoList GetVideoList Failed")
	}
	var favoList []*userFavoPb.Video
	for _, respVideo := range myResp.VideoList {
		author := userFavoPb.User{
			Id:            respVideo.Author.Id,
			Name:          respVideo.Author.Name,
			FollowCount:   respVideo.Author.FollowCount,
			FollowerCount: respVideo.Author.FollowerCount,
			IsFollow:      false,
		}
		video := userFavoPb.Video{
			Id:            respVideo.Id,
			Author:        &author,
			PlayUrl:       respVideo.PlayUrl,
			CoverUrl:      respVideo.CoverUrl,
			FavoriteCount: respVideo.FavoriteCount,
			CommentCount:  respVideo.CommentCount,
			IsFavorite:    respVideo.IsFavorite,
			Title:         respVideo.Title,
		}
		favoList = append(favoList, &video)
	}
	return &userFavoPb.FavoListResp{
			StatusCode: 0,
			StatusMsg:  "Success",
			VideoList:  favoList},
		nil
}

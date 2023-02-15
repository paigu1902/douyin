package main

import (
	"context"
	"fmt"
	"paigu1902/douyin/service/rpc-user-operator/cache"
	"paigu1902/douyin/service/rpc-user-operator/models"
	"paigu1902/douyin/service/rpc-user-operator/rabbitmq"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
	"strings"
	"time"
)

// UserFavoRpcImpl implements the last service interface defined in the IDL.
type UserFavoRpcImpl struct{}

// 执行点赞、取消赞操作
// FavoAction implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoAction(ctx context.Context, req *userFavoPb.FavoActionReq) (resp *userFavoPb.FavoActionResp, err error) {
	// TODO: Your code here...
	//	if req.Type != 1 && req.Type != 2 || req.UserId == 0 || req.VideoId == 0 {
	//		return nil, errors.New("Parameter Error")
	//	}
	//	likeRecord, err := models.GetLikeRecord(req.UserId, req.VideoId)
	//	if err != nil {
	//		err := models.CreateLikeRecord(req.UserId, req.VideoId)
	//		if err != nil {
	//			return nil, errors.New("Create LikeRecord Failed")
	//		}
	//	} else {
	//		if likeRecord.Status == 0 {
	//			err := models.UpdateLikeStatus(req.UserId, req.VideoId, 1)
	//			if err != nil {
	//				return nil, errors.New("Update LikeRecord Failed")
	//			}
	//		} else {
	//			err := models.UpdateLikeStatus(req.UserId, req.VideoId, 0)
	//			if err != nil {
	//				return nil, errors.New("Update LikeRecord Failed")
	//			}
	//		}
	//	}
	//	return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Succeed"}, nil
	user := fmt.Sprintf("%s", req.UserId)
	video := fmt.Sprintf("%s", req.VideoId)
	msg := strings.Builder{}
	msg.WriteString(user)
	msg.WriteString(" ")
	msg.WriteString((video))
	if req.Type == 1 { //点赞操作
		ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
		}
		if ext > 0 { //cache中存在点赞用户信息 加入视频信息
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			} else { //为避免脏数据 仅当cache操作成功后再操作MySQL
				//mq优化
				rabbitmq.RmqFavoAdd.Publish(msg.String())
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			videoIdList, err := models.GetFavoVideoId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			for _, favoVideoId := range videoIdList {
				_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
				if err != nil {
					cache.RdbFavoUser.Del(context.Background(), user)
					return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
				}
			}
			//将本次点赞视频信息加入cache
			_, errA := cache.RdbFavoUser.SAdd(context.Background(), user, video).Result()
			if errA != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errA
			} else {
				//mq优化
				rabbitmq.RmqFavoAdd.Publish(msg.String())
			}
		}
	} else { //取消点赞
		extU, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
		}
		if extU > 0 { //cache中存在点赞用户信息 删除视频信息
			_, err := cache.RdbFavoUser.SRem(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			} else { //为避免脏数据 仅当cache操作成功后再操作MySQL
				//mq优化
				rabbitmq.RmqFavoDel.Publish(msg.String())
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			videoIdList, err := models.GetFavoVideoId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			for _, favoVideoId := range videoIdList {
				_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
				if err != nil {
					cache.RdbFavoUser.Del(context.Background(), user)
					return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
				}
			}
			//将本次取消点赞视频信息加入cache
			_, errD := cache.RdbFavoUser.SRem(context.Background(), user, video).Result()
			if errD != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errD
			} else {
				//mq优化
				rabbitmq.RmqFavoDel.Publish(msg.String())
			}
		}
		//查询 RdbFavoVideo key:VideoId-value:UderId 中是否缓存此信息
		extV, err := cache.RdbFavoVideo.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
		}
		if extV > 0 { //cache中存在点赞用户信息 删除视频信息
			_, err := cache.RdbFavoVideo.SRem(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			//为避免脏数据，加入value:DefaultRedisValue，过期删除
			_, err := cache.RdbFavoVideo.SAdd(context.Background(), video, -1).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoVideo.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoVideo.Del(context.Background(), video)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			UserIdList, err := models.GetFavoUserId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
			}
			for _, favoUserId := range UserIdList {
				_, err := cache.RdbFavoVideo.SAdd(context.Background(), video, favoUserId).Result()
				if err != nil {
					cache.RdbFavoVideo.Del(context.Background(), video)
					return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, err
				}
			}
			//将本次取消点赞视频信息加入cache
			_, errD := cache.RdbFavoVideo.SRem(context.Background(), video, user).Result()
			if errD != nil {
				cache.RdbFavoVideo.Del(context.Background(), video)
				return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Failed"}, errD
			}
		}
	}
	return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Succeed"}, nil
}

// 获取用户的点赞视频列表
// FavoList implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoList(ctx context.Context, req *userFavoPb.FavoListReq) (resp *userFavoPb.FavoListResp, err error) {
	// TODO: Your code here...
	//	if req.UserId == 0 {
	//		return nil, errors.New("Parameter Error")
	//	}
	//	videoIdList, err := models.GetFavoVideoId(req.UserId)
	//	if err != nil {
	//		if err.Error() == "record not found" {
	//			return nil, err
	//		} else {
	//			return nil, errors.New("Get VideoList Failed")
	//		}
	//	}
	//	return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Succeed", VideoList: videoIdList}, nil
	user := fmt.Sprintf("%s", req.UserId)
	ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
	if err != nil {
		return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
	}
	if ext > 0 { //cache中存在点赞用户信息 获取视频列表
		videoIdList, err := cache.RdbFavoUser.SMembers(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
		}
		//调用VideoOperator
		myReq := VideoOperator.VideoListReq{VideoId: videoIdList}
		myResp, err := VideoOperator.VideoList(context.Background(), &myReq)
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
		}
		////用协程实现高并发查询Video类型对象并返回
		//favoVideoList := new([]userFavoPb.Video)
		//size := len(videoIdList) - 1 //去除DefaultValue
		//var wg sync.WaitGroup
		//wg.Add(size)
		//for i := 0; i <= size; i++ {
		//	if videoIdList[i] == -1 {
		//		continue
		//	}
		//	go AddVideo(req.UserId, videoIdList[i], favoVideoList, &wg)
		//}
		//wg.Wait()
		return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Succeed", VideoList: myResp.videoList}, nil
	} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
		_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
		if err != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
		}
		//设置数据有效期
		_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
		if errT != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, errT
		}
		//查询MySQL原有视频信息 加入cache
		videoIdList, errV := models.GetFavoVideoId(uint64(req.UserId))
		if errV != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, errV
		}
		for _, favoVideoId := range videoIdList {
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
			}
		}
		//调用VideoOperator
		myReq := VideoOperator.VideoListReq{VideoId: videoIdList}
		myResp, err := VideoOperator.VideoList(context.Background(), &myReq)
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Failed", VideoList: nil}, err
		}
		////用协程实现高并发查询Video类型对象并返回
		//favoVideoList := new([]userFavoPb.Video)
		//size := len(videoIdList) - 1 //去除DefaultValue
		//var wg sync.WaitGroup
		//wg.Add(size)
		//for i := 0; i <= size; i++ {
		//	if videoIdList[i] == -1 {
		//		continue
		//	}
		//	go AddVideo(req.UserId, videoIdList[i], favoVideoList, &wg)
		//}
		//wg.Wait()
		return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Succeed", VideoList: myResp.videoList}, nil
	}
}

//// TODO
//// 在用户点赞视频列表中添加视频对象
//func (s *UserFavoRpcImpl) AddVideo(userId int64, videoId int64, favoVideoList *[]userFavoPb.Video, wg *sync.WaitGroup) (bool, error) {
//	defer wg.Done()
//	video, err := GetVideo(userId, videoId)
//	if err != nil {
//		return false, err
//	}
//	*favoVideoList = append(*favoVideoList, video)
//	return true, nil
//}

// 查询用户对某条视频的点赞状态
func (s *UserFavoRpcImpl) FavoStatus(userId int64, videoId int64) (bool, error) {
	user := fmt.Sprintf("%s", userId)
	video := fmt.Sprintf("%s", videoId)
	//查询user-video表
	ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
	if err != nil {
		//return false, errors.New("Function IsFavorite Error")
		return false, err
	}
	if ext > 0 { //在user-video表中查到记录
		res, err := cache.RdbFavoUser.SIsMember(context.Background(), user, video).Result()
		if err != nil {
			return false, err
		}
		return res, nil
	} else { //查询video-user表
		ext, err := cache.RdbFavoVideo.Exists(context.Background(), video).Result()
		if err != nil {
			return false, err
		}
		if ext > 0 { //在video-user表中查到记录
			res, err := cache.RdbFavoVideo.SIsMember(context.Background(), video, user).Result()
			if err != nil {
				return false, err
			}
			return res, nil
		} else { //在cache中未查到记录 加入 default value 查询数据库
			//加入value:DefaultRedisValue，过期删除
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return false, err
			}
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return false, errT
			}
			//将数据查询结果加入cache
			videoList, err := models.GetFavoVideoId(uint64(userId))
			if err != nil {
				return false, err
			}
			for _, favoVideo := range videoList {
				str := fmt.Sprintf("%s", favoVideo)
				cache.RdbFavoUser.SAdd(context.Background(), user, str)
			}
			//再次查询cache
			res, err := cache.RdbFavoUser.SIsMember(context.Background(), user, video).Result()
			if err != nil {
				return false, err
			}
			return res, nil
		}
	}
}

// 查询视频被点赞总数
func (s *UserFavoRpcImpl) FavoCount(videoId int64) (int64, error) {
	video := fmt.Sprintf("%s", videoId)
	//查询video-user表
	ext, err := cache.RdbFavoVideo.Exists(context.Background(), video).Result()
	if err != nil {
		//return false, errors.New("Function IsFavorite Error")
		return 0, err
	}
	if ext > 0 { //在user-video表中查到记录
		res, err := cache.RdbFavoVideo.SCard(context.Background(), video).Result()
		if err != nil {
			return 0, err
		}
		return res - 1, nil //减去 default value
	} else { //在cache中未查到记录 加入 default value 查询数据库
		//加入value:DefaultRedisValue，过期删除
		_, err := cache.RdbFavoVideo.SAdd(context.Background(), video, -1).Result()
		if err != nil {
			cache.RdbFavoVideo.Del(context.Background(), video)
			return 0, err
		}
		_, errT := cache.RdbFavoVideo.Expire(context.Background(), video, time.Duration(30)*time.Second).Result()
		if errT != nil {
			cache.RdbFavoVideo.Del(context.Background(), video)
			return 0, errT
		}
		//将数据查询结果加入cache
		userList, err := models.GetFavoUserId(uint64(videoId))
		if err != nil {
			return 0, err
		}
		for _, favoUser := range userList {
			str := fmt.Sprintf("%s", favoUser)
			cache.RdbFavoUser.SAdd(context.Background(), video, str)
		}
		//再次查询cache
		res, err := cache.RdbFavoUser.SCard(context.Background(), video).Result()
		if err != nil {
			return 0, err
		}
		return res - 1, nil //减去 default value
	}
}

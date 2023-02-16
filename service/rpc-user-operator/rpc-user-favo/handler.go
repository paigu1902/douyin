package main

import (
	"context"
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

// 执行点赞、取消赞操作
// FavoAction implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoAction(ctx context.Context, req *userFavoPb.FavoActionReq) (resp *userFavoPb.FavoActionResp, err error) {
	// TODO: Your code here..
	user := strconv.FormatInt(req.UserId, 10)
	video := strconv.FormatInt(req.VideoId, 10)
	msg := strings.Builder{}
	msg.WriteString(user)
	msg.WriteString(" ")
	msg.WriteString((video))
	if req.Type == 1 { //点赞操作
		ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
		}
		if ext > 0 { //cache中存在点赞用户信息 加入视频信息
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			} else { //为避免脏数据 仅当cache操作成功后再操作MySQL
				//mq优化
				rabbitmq.RmqFavoAdd.Publish(msg.String())
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			videoIdList, err := models.GetFavoVideoId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			for _, favoVideoId := range videoIdList {
				_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
				if err != nil {
					cache.RdbFavoUser.Del(context.Background(), user)
					return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
				}
			}
			//将本次点赞视频信息加入cache
			_, errA := cache.RdbFavoUser.SAdd(context.Background(), user, video).Result()
			if errA != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errA
			} else {
				//mq优化
				rabbitmq.RmqFavoAdd.Publish(msg.String())
			}
		}
	} else { //取消点赞
		extU, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
		}
		if extU > 0 { //cache中存在点赞用户信息 删除视频信息
			_, err := cache.RdbFavoUser.SRem(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			} else { //为避免脏数据 仅当cache操作成功后再操作MySQL
				//mq优化
				rabbitmq.RmqFavoDel.Publish(msg.String())
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			videoIdList, err := models.GetFavoVideoId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			for _, favoVideoId := range videoIdList {
				_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
				if err != nil {
					cache.RdbFavoUser.Del(context.Background(), user)
					return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
				}
			}
			//将本次取消点赞视频信息加入cache
			_, errD := cache.RdbFavoUser.SRem(context.Background(), user, video).Result()
			if errD != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errD
			} else {
				//mq优化
				rabbitmq.RmqFavoDel.Publish(msg.String())
			}
		}
		//查询 RdbFavoVideo key:VideoId-value:UderId 中是否缓存此信息
		extV, err := cache.RdbFavoVideo.Exists(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
		}
		if extV > 0 { //cache中存在点赞用户信息 删除视频信息
			_, err := cache.RdbFavoVideo.SRem(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
		} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
			//为避免脏数据，加入value:DefaultRedisValue，过期删除
			_, err := cache.RdbFavoVideo.SAdd(context.Background(), video, -1).Result()
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			//设置数据有效期
			_, errT := cache.RdbFavoVideo.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoVideo.Del(context.Background(), video)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errT
			}
			//查询MySQL原有视频信息 加入cache
			UserIdList, err := models.GetFavoUserId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
			}
			for _, favoUserId := range UserIdList {
				_, err := cache.RdbFavoVideo.SAdd(context.Background(), video, favoUserId).Result()
				if err != nil {
					cache.RdbFavoVideo.Del(context.Background(), video)
					return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, err
				}
			}
			//将本次取消点赞视频信息加入cache
			_, errD := cache.RdbFavoVideo.SRem(context.Background(), video, user).Result()
			if errD != nil {
				cache.RdbFavoVideo.Del(context.Background(), video)
				return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Failed"}, errD
			}
		}
	}
	return &userFavoPb.FavoActionResp{StatusCode: 0, StatusMsg: "Succeed"}, nil
}

// 获取用户的点赞视频列表
// FavoList implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoList(ctx context.Context, req *userFavoPb.FavoListReq) (resp *userFavoPb.FavoListResp, err error) {
	// TODO: Your code here...
	user := strconv.FormatInt(req.UserId, 10)
	ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
	if err != nil {
		return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
	}
	if ext > 0 { //cache中存在点赞用户信息 获取视频列表
		videoIdListStr, err := cache.RdbFavoUser.SMembers(context.Background(), user).Result()
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
		}
		var videoIdList []uint64
		for index, str := range videoIdListStr {
			id, _ := strconv.Atoi(str)
			videoIdList[index] = uint64(id)
		}
		//获取videoOperator客户端
		resp, err := rpcClient.VideoOpClient.VideoList(ctx, &VideoOptPb.VideoListReq{VideoId: videoIdList})
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
		}
		var favoList []*userFavoPb.Video
		for _, respVideo := range resp.VideoList {
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
		return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Success", VideoList: favoList}, err
	} else { //cache中不存在用户信息 查询MySQL加入原有视频信息后更新
		_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
		if err != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
		}
		//设置数据有效期
		_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
		if errT != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, errT
		}
		//查询MySQL原有视频信息 加入cache
		videoIdList, errV := models.GetFavoVideoId(uint64(req.UserId))
		if errV != nil {
			cache.RdbFavoUser.Del(context.Background(), user)
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, errV
		}
		for _, favoVideoId := range videoIdList {
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, favoVideoId).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
			}
		}
		//获取videoOperator客户端
		resp, err := rpcClient.VideoOpClient.VideoList(ctx, &VideoOptPb.VideoListReq{VideoId: videoIdList})
		if err != nil {
			return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Failed", VideoList: nil}, err
		}
		var favoList []*userFavoPb.Video
		for _, respVideo := range resp.VideoList {
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
		return &userFavoPb.FavoListResp{StatusCode: 0, StatusMsg: "Success", VideoList: favoList}, err
	}
}

// 查询用户对某条视频的点赞状态
func (s *UserFavoRpcImpl) FavoStatus(ctx context.Context, req *userFavoPb.FavoStatusReq) (resp *userFavoPb.FavoStatusResp, err error) {
	user := strconv.FormatInt(req.UserId, 10)
	video := strconv.FormatInt(req.VideoId, 10)
	//查询user-video表
	ext, err := cache.RdbFavoUser.Exists(context.Background(), user).Result()
	if err != nil {
		return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
	}
	if ext > 0 { //在user-video表中查到记录
		res, err := cache.RdbFavoUser.SIsMember(context.Background(), user, video).Result()
		if err != nil {
			return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
		}
		return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, err
	} else { //查询video-user表
		ext, err := cache.RdbFavoVideo.Exists(context.Background(), video).Result()
		if err != nil {
			return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
		}
		if ext > 0 { //在video-user表中查到记录
			res, err := cache.RdbFavoVideo.SIsMember(context.Background(), video, user).Result()
			if err != nil {
				return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
			}
			return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, err
		} else { //在cache中未查到记录 加入 default value 查询数据库
			//加入value:DefaultRedisValue，过期删除
			_, err := cache.RdbFavoUser.SAdd(context.Background(), user, -1).Result()
			if err != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
			}
			_, errT := cache.RdbFavoUser.Expire(context.Background(), user, time.Duration(30)*time.Second).Result()
			if errT != nil {
				cache.RdbFavoUser.Del(context.Background(), user)
				return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, errT
			}
			//将数据查询结果加入cache
			videoList, err := models.GetFavoVideoId(uint64(req.UserId))
			if err != nil {
				return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
			}
			for _, favoVideo := range videoList {
				str := strconv.FormatInt(int64(favoVideo), 10)
				cache.RdbFavoUser.SAdd(context.Background(), user, str)
			}
			//再次查询cache
			res, err := cache.RdbFavoUser.SIsMember(context.Background(), user, video).Result()
			if err != nil {
				return &userFavoPb.FavoStatusResp{StatusCode: 1, StatusMsg: "Failed", IsFavorite: false}, err
			}
			return &userFavoPb.FavoStatusResp{StatusCode: 0, StatusMsg: "Success", IsFavorite: res}, err
		}
	}
}

// 查询视频被点赞总数
func (s *UserFavoRpcImpl) FavoCount(videoId int64) (int64, error) {
	video := strconv.FormatInt(videoId, 10)
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
			str := strconv.FormatInt(int64(favoUser), 10)
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

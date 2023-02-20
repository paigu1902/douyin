package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
	"paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// VideoOperatorImpl implements the last service interface defined in the IDL.
type VideoOperatorImpl struct{}

type VideoInfoCache struct {
	Id            uint64 `redis:"Id"`
	AuthorId      uint64 `redis:"AuthorId"`
	Title         string `redis:"Title"`
	PlayUrl       string `redis:"PlayUrl"`
	CoverUrl      string `redis:"CoverUrl"`
	FavoriteCount int64  `redis:"FavoriteCount"`
	CommentCount  int64  `redis:"CommentCount"`
}

func writeVideoCache(ctx context.Context, redisKey string, videoDbList []models.VideoInfo) {
	for _, v := range videoDbList {
		err := cache.RDB.ZAdd(ctx, redisKey, &redis.Z{
			Score:  float64(v.CreatedAt.UnixMilli()),
			Member: v.ID,
		}).Err()
		if err != nil {
			klog.Error(err)
		}
		if err != nil {
			klog.Error("写入缓存错误")
		}
		err = cache.RDB.HSet(ctx, "VideoInfo:"+strconv.FormatUint(uint64(v.ID), 10), map[string]interface{}{
			"AuthorId":      v.AuthorId,
			"Title":         v.Title,
			"PlayUrl":       v.PlayUrl,
			"CoverUrl":      v.CoverUrl,
			"FavoriteCount": v.FavoriteCount,
			"CommentCount":  v.CommentCount,
		}).Err()
		if err != nil {
			klog.Error(err)
		}
	}
}

// Upload implements the VideoOperatorImpl interface.
func (s *VideoOperatorImpl) Upload(ctx context.Context, req *videoOperatorPb.VideoUploadReq) (resp *videoOperatorPb.VideoUploadResp, err error) {
	// TODO: Your code here...
	token := req.Token
	title := req.Title
	data := req.Data
	claims, err := utils.AnalyseToken(token)
	if err != nil {
		// debug 专用
		//claims = &utils.UserClaims{
		//	ID:             0,
		//	Name:           "wqxtest",
		//	StandardClaims: jwt.StandardClaims{},
		//}
		return nil, err
	}

	fileName := fmt.Sprintf("%d_%d.mp4", time.Now().UnixNano(), claims.ID)

	playUrl, err := utils.Upload(data, fileName, "videos")
	if err != nil {
		return nil, err
	}
	coverTmpPath, err := extractCover(playUrl)
	if err != nil {
		return nil, err
	}
	coverUrl, err := utils.Upload(coverTmpPath, filepath.Base(coverTmpPath), "covers")
	if err != nil {
		return nil, err
	}
	info := models.VideoInfo{
		AuthorId: uint64(claims.ID),
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	err = models.CreateVideoInfo(&info)
	if err != nil {
		return nil, err
	}
	redisKey := "VideoFeed"
	writeVideoCache(ctx, redisKey, []models.VideoInfo{info})
	return &videoOperatorPb.VideoUploadResp{Status: 0, StatusMsg: "成功"}, nil
}

// Feed implements the VideoOperatorImpl interface.
func (s *VideoOperatorImpl) Feed(ctx context.Context, req *videoOperatorPb.FeedReq) (resp *videoOperatorPb.FeedResp, err error) {
	// TODO: Your code here...
	if req == nil {
		req = &videoOperatorPb.FeedReq{
			LatestTime: 0,
			Token:      "",
		}
	}
	var userId uint64
	if req.Token != "" {
		u, err := strconv.Atoi(req.Token)
		if err != nil {
			return nil, errors.New("id不正确")
		}
		userId = uint64(u)
	}

	limit := 30
	redisKey := "VideoFeed"

	timestamp := req.LatestTime
	if timestamp == 0 {
		timestamp = time.Now().UnixMilli()
	}

	var videoList []models.VideoInfo

	//1. redis cache
	videoIds, err := cache.RDB.ZRevRangeByScoreWithScores(ctx, redisKey, &redis.ZRangeBy{Count: 31, Max: strconv.FormatInt(timestamp, 10)}).Result()
	if err != nil {
		klog.Error("redis缓存错误")
	}
	videoCacheList := make([]models.VideoInfo, len(videoIds))
	videoWithoutInfo := make(map[uint64]int, 0)
	for i, v := range videoIds {
		videoHash, err := cache.RDB.HGetAll(ctx, "VideoInfo:"+v.Member.(string)).Result()
		if err != nil {
			id, _ := strconv.Atoi(v.Member.(string))
			videoWithoutInfo[uint64(id)] = i
			klog.Error("redis缓存错误")
			continue
		}
		authorId, err := strconv.Atoi(videoHash["AuthorId"])
		favoriteCount, err := strconv.Atoi(videoHash["FavoriteCount"])
		commentCount, err := strconv.Atoi(videoHash["CommentCount"])

		vid, err := strconv.Atoi(v.Member.(string))
		if err != nil {
			klog.Error("缓存错误")
			continue
		}

		videoCacheList[i] = models.VideoInfo{
			Model:         gorm.Model{ID: uint(vid)},
			AuthorId:      uint64(authorId),
			Title:         videoHash["Title"],
			PlayUrl:       videoHash["PlayUrl"],
			CoverUrl:      videoHash["CoverUrl"],
			FavoriteCount: int64(favoriteCount),
			CommentCount:  int64(commentCount),
		}
	}
	ids := make([]uint64, len(videoWithoutInfo))
	for k, _ := range videoWithoutInfo {
		ids = append(ids, k)
	}
	var infos []models.VideoInfo
	err = models.DB.Where(ids).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	writeVideoCache(ctx, redisKey, infos)

	var latestTime string
	if len(videoIds) == 0 {
		latestTime = time.UnixMilli(timestamp).Format("2006-01-02 15:04:05")
	} else {
		latestTime = time.UnixMilli(int64(videoIds[len(videoIds)-1].Score)).Format("2006-01-02 15:04:05")
	}
	videoList = videoCacheList

	//2. db
	if len(videoIds) < 31 {
		var videoDbList []models.VideoInfo
		err = models.GetVideoInfo(latestTime, limit+1-len(videoIds), &videoDbList)
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, videoDbList...)
		writeVideoCache(ctx, redisKey, videoDbList)
	}
	nextTime := int64(0)
	//没滑到底
	if len(videoList) > limit {
		videoList = videoList[:len(videoList)-1]
		nextTime = videoList[len(videoList)-1].CreatedAt.UnixMilli()
	} else { //到底了
		nextTime = 1
	}

	var videoRespList []*videoOperatorPb.Video

	//获取作者信息
	authorIdList := make([]uint64, len(videoList))
	for i, v := range videoList {
		authorIdList[i] = v.AuthorId
	}
	r, err := rpcClient.UserInfo.BatchInfo(ctx, &userInfoPb.BatchUserReq{
		Fromid:   userId,
		Batchids: authorIdList,
	})
	authorInfos := r.GetBatchusers()

	for i, videoInfo := range videoList {
		author := videoOperatorPb.User{
			Id:            videoInfo.AuthorId,
			Name:          authorInfos[i].UserName,
			FollowCount:   authorInfos[i].FollowCount,
			FollowerCount: authorInfos[i].FollowerCount,
			IsFollow:      authorInfos[i].IsFollow,
		}
		video := videoOperatorPb.Video{
			Id:            uint64(videoInfo.ID),
			Author:        &author,
			PlayUrl:       videoInfo.PlayUrl,
			CoverUrl:      videoInfo.CoverUrl,
			FavoriteCount: videoInfo.FavoriteCount,
			CommentCount:  videoInfo.CommentCount,
			IsFavorite:    false,
			Title:         videoInfo.Title,
		}

		//用户登录状态的话，查询用户是否点赞视频
		if req.Token != "" {
			userFavoResp, err := rpcClient.UserFavo.FavoStatus(ctx, &userFavoPb.FavoStatusReq{
				UserId:  int64(userId),
				VideoId: int64(video.Id),
			})
			if err != nil {
				return nil, err
			}
			video.IsFavorite = userFavoResp.GetIsFavorite()
		}

		videoRespList = append(videoRespList, &video)
	}
	feedResp := &videoOperatorPb.FeedResp{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videoRespList,
		NextTime:   nextTime,
	}
	return feedResp, nil
}

// 截取视频第一秒截图，保存在本地临时文件中并返回文件地址
// 该方法成功执行需要本地安装ffmpeg
func extractCover(playUrl string) (string, error) {
	tmpCoverDir := filepath.Join("tmp", "cover")
	coverTmpPath := filepath.Join(tmpCoverDir, strings.Split(filepath.Base(playUrl), ".")[0]+".jpg")
	err := os.MkdirAll(tmpCoverDir, 0777)
	if err != nil {
		return "", err
	}
	cmd := exec.Command("ffmpeg", "-ss", "00:00:01", "-i", playUrl, "-frames:v", "1", coverTmpPath, "-r", "1", "-an", "-y", "-f", "mjpeg")
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return coverTmpPath, nil
}

// PublishList implements the VideoOperatorImpl interface.
func (s *VideoOperatorImpl) PublishList(ctx context.Context, req *videoOperatorPb.PublishListReq) (resp *videoOperatorPb.PublishListResp, err error) {
	//1. 根据user_id,获取author的信息
	authorId, userId := req.AuthorId, req.UserId
	authInfoReq := userInfoPb.UserInfoReq{FromId: userId, ToId: authorId}
	authorInfo, err := rpcClient.UserInfo.Info(ctx, &authInfoReq)
	if err != nil {
		resp = &videoOperatorPb.PublishListResp{
			StatusCode: 1,
			StatusMsg:  "author_id 不存在",
		}
		return resp, err
	}
	// 2.根据author信息，查询发布的视频
	var videoList []models.VideoInfo
	err = models.GetVideoListByAuthorId(authorInfo.GetUser().GetUserId(), &videoList)
	if err != nil {
		resp = &videoOperatorPb.PublishListResp{
			StatusCode: 1,
			StatusMsg:  "查询视频错误",
		}
		return resp, err
	}
	followCnt := authorInfo.GetUser().GetFollowCount()
	followerCnt := authorInfo.GetUser().GetFollowerCount()
	//3.需要判断用户是否关注该作者
	//isFollowResp, err := rpcClient.UserRelationClient.IsFollow(ctx, &userRelationPb.IsFollowReq{
	//	FromId: userId,
	//	ToId:   authorId,
	//})
	//if err != nil {
	//	resp = &videoOperatorPb.PublishListResp{
	//		StatusCode: 1,
	//		StatusMsg:  "查询用户关注错误",
	//	}
	//	return resp, err
	//}
	author := &videoOperatorPb.User{
		Id:            authorInfo.User.GetUserId(),
		FollowCount:   followCnt,
		FollowerCount: followerCnt,
		IsFollow:      authorInfo.User.IsFollow,
	}

	var videos []*videoOperatorPb.Video
	for _, v := range videoList {
		//TODO: isFavourite字段需要后续，根据userFavo获取
		video := v.TransToVideo()
		video.Author = author
		favoStatusResp, er := rpcClient.UserFavo.FavoStatus(ctx, &userFavoPb.FavoStatusReq{UserId: int64(userId), VideoId: int64(v.ID)})
		if er != nil {
			resp = &videoOperatorPb.PublishListResp{
				StatusCode: 1,
				StatusMsg:  "查询用户点赞错误",
			}
			return resp, er
		}
		video.IsFavorite = favoStatusResp.IsFavorite
		videos = append(videos, video)
	}
	resp = &videoOperatorPb.PublishListResp{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videos,
	}
	return resp, nil
}

// VideoList implements the VideoOperatorImpl interface.
func (s *VideoOperatorImpl) VideoList(ctx context.Context, req *videoOperatorPb.VideoListReq) (resp *videoOperatorPb.VideoListResp, err error) {
	videoIdList := req.GetVideoId()
	var videos []models.VideoInfo
	if err = models.GetVideosByIds(videoIdList, &videos); err != nil {
		resp = &videoOperatorPb.VideoListResp{
			StatusCode: 1,
			StatusMsg:  "失败",
		}
		return resp, err
	}
	var videoList []*videoOperatorPb.Video
	for _, v := range videos {
		video := v.TransToVideo()
		video.Author = &videoOperatorPb.User{Id: v.AuthorId}
		videoList = append(videoList, video)
	}
	resp = &videoOperatorPb.VideoListResp{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videoList,
	}
	return resp, nil
}

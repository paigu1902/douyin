package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"os"
	"os/exec"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb/userinfo"
	videoOperatorPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
	"path/filepath"
	"strings"
	"time"
)

// VideoOperatorImpl implements the last service interface defined in the IDL.
type VideoOperatorImpl struct{}

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
		AuthorId: int64(claims.ID),
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	err = models.CreateVideoInfo(&info)
	if err != nil {
		return nil, err
	}
	return &videoOperatorPb.VideoUploadResp{Status: 1, StatusMsg: "成功"}, nil
}

// Feed implements the VideoOperatorImpl interface.
func (s *VideoOperatorImpl) Feed(ctx context.Context, req *videoOperatorPb.FeedReq) (resp *videoOperatorPb.FeedResp, err error) {
	// TODO: Your code here...
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	userInfoClient := userinfo.MustNewClient(
		"userInfoImpl",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*5),
	)

	limit := 30
	token := req.Token
	// todo: timestamp to UTC time format，抓包看一下timestamp的形式
	timestamp := req.LatestTime
	latestTime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	var videoList []models.VideoInfo
	err = models.GetVideoInfo(latestTime, limit+1, &videoList)
	if err != nil {
		return nil, err
	}
	nextTime := videoList[len(videoList)-1].CreatedAt.Unix()
	videoList = videoList[:len(videoList)-1]
	var videoRespList []*videoOperatorPb.Video
	for _, videoInfo := range videoList {
		userInfoReq := userInfoPb.UserInfoReq{
			UserId: uint64(videoInfo.AuthorId),
			Token:  token,
		}
		authorInfo, err := userInfoClient.Info(ctx, &userInfoReq)
		if err != nil {
			return nil, err
		}
		author := videoOperatorPb.User{
			Id:            videoInfo.AuthorId,
			Name:          authorInfo.User.UserName,
			FollowCount:   authorInfo.User.FollowCount,
			FollowerCount: authorInfo.User.FollowerCount,
			IsFollow:      false,
		}
		videoRespList = append(videoRespList, &videoOperatorPb.Video{
			Id:            int64(videoInfo.ID),
			Author:        &author,
			PlayUrl:       videoInfo.PlayUrl,
			CoverUrl:      videoInfo.CoverUrl,
			FavoriteCount: videoInfo.FavoriteCount,
			CommentCount:  videoInfo.CommentCount,
			IsFavorite:    false, //todo: 修改
			Title:         videoInfo.Title,
		})
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

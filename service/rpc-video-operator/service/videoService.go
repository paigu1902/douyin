package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-video-operator/models"
	pb "paigu1902/douyin/service/rpc-video-operator/videoOperatorPb"
	"path/filepath"
	"strings"
	"time"
)

type VideoService struct {
	pb.UnimplementedVideoOperatorServer
}

func (s *VideoService) Upload(ctx context.Context, req *pb.VideoUploadReq) (*pb.VideoUploadResp, error) {
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
	return &pb.VideoUploadResp{Status: 1, StatusMsg: "成功"}, nil
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

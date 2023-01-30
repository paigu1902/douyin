package service

import (
	"context"
	"gorm.io/gorm"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-video-operator/models"
	pb "paigu1902/douyin/service/rpc-video-operator/videoOperatorPb"
	"strconv"
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
	info := models.VideoInfo{
		AuthorId: int64(claims.ID),
		Title:    title,
		PlayUrl:  "",
		CoverUrl: "",
	}
	err = uploadTx(&info, data, strconv.FormatInt(int64(claims.ID), 10)+".mp4")
	if err != nil {
		return nil, err
	}
	return &pb.VideoUploadResp{Status: 1, StatusMsg: "成功"}, nil
}

func uploadTx(info *models.VideoInfo, data []byte, name string) error {
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(info).Error; err != nil {
			return err
		}
		playUrl, err := utils.Upload(data, name)
		if err != nil {
			return err
		}
		info.PlayUrl = playUrl
		if err := tx.Model(info).Update("play_url", info.PlayUrl).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

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
		return nil, err
	}
	info := models.VideoInfo{
		AuthorId: int64(claims.ID),
		Title:    title,
		PlayUrl:  "",
		CoverUrl: "",
	}
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if err := models.CreateVideoInfo(&info); err != nil {
			return err
		}
		filename, err := utils.Upload(data, strconv.FormatInt(int64(claims.ID), 10)+".mp4")
		if err != nil {
			return err
		}
		info.PlayUrl = filename
		if err := models.UpdateVideoInfo(&info); err != nil {
			return err
		}
		return nil
	})
	return &pb.VideoUploadResp{Status: 1, StatusMsg: "成功"}, nil
}

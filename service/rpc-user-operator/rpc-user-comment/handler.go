package main

import (
	"context"
	"errors"
	"paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/rpc-user-operator/models"
	UserCommPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
	"time"
)

// UserCommRpcImpl implements the last service interface defined in the IDL.
type UserCommRpcImpl struct{}

// InsertComment implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) InsertComment(ctx context.Context, req *UserCommPb.DouyinCommentActionRequest) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	token := req.Token
	commentTxt := req.CommentText
	commentId := req.CommentId
	video_id := req.VideoId
	claims, err := utils.AnalyseToken(token)
	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 3,
			StatusMsg:  "USER TOKEN ERROR",
		}, errors.New("Analyse Token Error")
	}
	comment := UserCommPb.Comment{
		Id: commentId,
		User: &UserCommPb.User{
			Id:   int64(claims.ID),
			Name: claims.Name,
		},
		Content:    commentTxt,
		CreateDate: time.Now().Format("1-2"),
	}
	if req.CommentId == 1 {
		// 发表评论
		err := models.InsertComment(models.UserComm{
			UserName: claims.Name,
			VideoId:  video_id,
			CommText: commentTxt,
		})
		if err != nil {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 2,
				StatusMsg:  "OTHER_ERROR",
			}, err
		} else {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 0,
				StatusMsg:  "SUCCESS",
				Comment:    &comment,
			}, nil
		}
	} else {
		// 删除评论
		err := models.DeleteComment(req.CommentId)
		if err != nil {
			if err.Error() == "del comment is not exist" {
				return &UserCommPb.DouyinCommentActionResponse{
					StatusCode: 1,
					StatusMsg:  "NOT_EXIST_ERROR",
				}, err
			} else {
				return &UserCommPb.DouyinCommentActionResponse{
					StatusCode: 2,
					StatusMsg:  "OTHER_ERROR",
				}, err
			}
		} else {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 0,
				StatusMsg:  "SUCCESS",
				Comment:    &comment,
			}, nil
		}
	}
}

// GetCommentsByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentsByVideo(ctx context.Context, req *UserCommPb.DouyinCommentListRequest) (resp *UserCommPb.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	videoId := req.VideoId
	commentList, err := models.GetCommentsByVideo(videoId)
	if err != nil {
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	} else {
		respCommentList, err := utils.FillCommentListFields(commentList)
		if err != nil {
			// 评论为空，此时应该只是提示，不报错
			return &UserCommPb.DouyinCommentListResponse{
				StatusCode: 1,
				StatusMsg:  "NOT_EXIST_LIST",
				//CommentList: UserCommPb.Comment{
				//	Id: comment_list
				//},
				CommentList: respCommentList,
			}, nil
		}
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode: 0,
			StatusMsg:  "SUCCESS",
			//CommentList: UserCommPb.Comment{
			//	Id: comment_list
			//},
			CommentList: respCommentList,
		}, nil
	}
}

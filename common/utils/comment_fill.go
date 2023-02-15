package utils

import (
	"context"
	"errors"
	UserInfo "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-info/logic"
	"paigu1902/douyin/service/rpc-user-operator/models"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
)

func FillCommentListFields(comments []models.UserComm) ([]*UserCommPb.Comment, error) {
	size := len(comments)
	var commentListPb []*UserCommPb.Comment
	if comments == nil || size == 0 {
		return commentListPb, errors.New("Find List is Empty")
	}
	var userids []uint64
	for _, com := range comments {
		userids = append(userids, com.UserId)
	}
	myReq := UserInfo.BatchUserReq{
		Batchids: userids,
		Fromid:   uint64(comments[0].VideoId), // TODO: Find AuthorId
		// 查找videoid的author
	}
	myRes, _ := logic.BatchInfo(context.Background(), &myReq)
	for i, v := range comments {
		//userid := v.UserId
		//user := UserInfo.FindUserByID(uint64(userid))
		user := myRes.Batchusers[i]
		commentListPb = append(commentListPb, &UserCommPb.Comment{
			Id: int64(v.ID),
			User: &UserCommPb.User{
				Id:            int64(user.UserId),
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			Content:    v.CommText,
			CreateDate: v.CreatedAt.Format("1-2"),
		})
	}
	return commentListPb, nil
}

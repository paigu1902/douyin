package utils

import (
	"errors"
	UserInfo "paigu1902/douyin/service/rpc-user-info/models"
	"paigu1902/douyin/service/rpc-user-operator/models"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
)

func FillCommentListFields(comments []models.UserComm) ([]*UserCommPb.Comment, error) {
	size := len(comments)
	var commentListPb []*UserCommPb.Comment
	if comments == nil || size == 0 {
		return commentListPb, errors.New("Find List is Empty")
	}

	for _, v := range comments {
		userid := v.UserId
		user := UserInfo.FindUserByID(uint64(userid))
		commentListPb = append(commentListPb, &UserCommPb.Comment{
			Id: int64(v.ID),
			User: &UserCommPb.User{
				Id:            int64(user.ID),
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowedCount,
				IsFollow:      false,
			},
			Content:    v.CommText,
			CreateDate: v.CreatedAt.Format("1-2"),
		})
	}
	return commentListPb, nil
}

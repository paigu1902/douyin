// Code generated by hertz generator.

package userRelationPb

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	dyUtils "paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"strconv"
)

type UserHttp struct {
	UserId        uint64 `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type FollowActionReq struct {
	ToId       string `query:"to_user_id"`
	ActionType string `query:"action_type"`
}

type MessageActionReq struct {
	ToId       string `query:"to_user_id"`
	ActionType string `query:"action_type"`
	Content    string `query:"content"`
}

type MessageHistoryReq struct {
	ToId string `query:"to_user_id"`
}

func getUsers(users []*userRelationPb.User) []*UserHttp {
	res := make([]*UserHttp, len(users))
	for i, v := range users {
		res[i] = &UserHttp{
			UserId:        v.GetUserId(),
			UserName:      v.GetUserName(),
			FollowCount:   v.GetFollowCount(),
			FollowerCount: v.GetFollowerCount(),
			IsFollow:      v.GetIsFollow(),
		}
	}
	return res
}

func stringToUint64(intStr string) (uint64, error) {
	atoi, err := strconv.Atoi(intStr)
	if err != nil {
		return 0, err
	}
	return uint64(atoi), nil
}

func FollowAction(ctx context.Context, c *app.RequestContext) {
	req := new(FollowActionReq)
	// 1. 绑定校验参数
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id异常"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	// 2.调用rpc
	resp, err := rpcClient.UserRelationClient.FollowAction(ctx, &userRelationPb.FollowActionReq{FromId: fromId, ToId: toId, Type: req.ActionType})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg()},
	)
	return
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	id, exist := c.GetQuery("user_id")
	if exist != true {
		c.JSON(400, errors.New("参数userid异常"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("to_id异常"))
		return
	}
	resp, err := rpcClient.UserRelationClient.FollowList(ctx, &userRelationPb.FollowListReq{UserId: userId})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user_list":   getUsers(resp.GetUserList())},
	)
	return
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	id, exist := c.GetQuery("user_id")
	if exist != true {
		c.JSON(400, errors.New("参数userid异常"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("id异常"))
		return
	}
	resp, err := rpcClient.UserRelationClient.FollowerList(ctx, &userRelationPb.FollowerListReq{UserId: userId})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user_list":   getUsers(resp.GetUserList())},
	)
	return
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	id, exist := c.GetQuery("user_id")
	if exist != true {
		c.JSON(400, errors.New("参数userid异常"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("id异常"))
		return
	}
	resp, err := rpcClient.UserRelationClient.FriendList(ctx, &userRelationPb.FriendListReq{UserId: userId})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user_list":   getUsers(resp.GetUserList())},
	)
	return
}

func MessageAction(ctx context.Context, c *app.RequestContext) {
	req := new(MessageActionReq)
	// 1. 绑定校验参数
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id异常"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	// 2.调用rpc
	resp, err := rpcClient.UserRelationClient.SendMessage(ctx, &userRelationPb.SendMessageReq{FromId: fromId, ToId: toId, Type: req.ActionType, Content: req.Content})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg()},
	)
	return
}

func MessageHistory(ctx context.Context, c *app.RequestContext) {
	req := new(MessageHistoryReq)
	// 1. 绑定校验参数
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var toId uint64
	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id异常"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}
	// 2.调用rpc
	resp, err := rpcClient.UserRelationClient.HistoryMessage(ctx, &userRelationPb.HistoryMessageReq{FromId: fromId, ToId: toId})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code":  resp.GetStatusCode(),
		"status_msg":   resp.GetStatusMsg(),
		"message_list": resp.GetMessageList()},
	)
	return
}

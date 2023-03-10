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
	UserId          uint64 `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar" default:"https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c"`
	BackgroundImage string `json:"background_image" default:"https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c"`
	Signature       string `json:"signature" default:"666"`
	TotalFavorited  string `json:"total_favorited" default:"0"`
	WorkCount       int64  `json:"work_count" default:"0"`
	FavoriteCount   int64  `json:"favorite_count" default:"0"`
}

type MessageHttp struct {
	Id         int    `json:"id"`
	ToUserId   int    `json:"to_user_id"`
	FromUserId int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
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
			UserId:          v.GetUserId(),
			UserName:        v.GetUserName(),
			FollowCount:     v.GetFollowCount(),
			FollowerCount:   v.GetFollowerCount(),
			IsFollow:        v.GetIsFollow(),
			Avatar:          "https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c",
			BackgroundImage: "https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c",
			Signature:       "666",
			TotalFavorited:  "0",
			WorkCount:       0,
			FavoriteCount:   0,
		}
	}
	return res
}

func getMessage(messages []*userRelationPb.MessageContent) []*MessageHttp {
	res := make([]*MessageHttp, len(messages))
	for i, v := range messages {
		res[i] = &MessageHttp{
			Id:         int(v.GetId()),
			ToUserId:   int(v.GetToId()),
			FromUserId: int(v.GetFromId()),
			Content:    v.GetContent(),
			CreateTime: v.GetCreateTime(),
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
	// 1. ??????????????????
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id??????"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	// 2.??????rpc
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
		c.JSON(400, errors.New("??????userid??????"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("to_id??????"))
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
		c.JSON(400, errors.New("??????userid??????"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("id??????"))
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
		c.JSON(400, errors.New("??????userid??????"))
		return
	}
	userId, err := stringToUint64(id)
	if err != nil {
		c.JSON(400, errors.New("id??????"))
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
	// 1. ??????????????????
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id??????"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	// 2.??????rpc
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
	// 1. ??????????????????
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var toId uint64
	toId, err := stringToUint64(req.ToId)
	if err != nil {
		c.JSON(400, errors.New("to_id??????"))
		return
	}

	fromId, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err)
		return
	}
	// 2.??????rpc
	resp, err := rpcClient.UserRelationClient.HistoryMessage(ctx, &userRelationPb.HistoryMessageReq{FromId: fromId, ToId: toId})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, utils.H{
		"status_code":  resp.GetStatusCode(),
		"status_msg":   resp.GetStatusMsg(),
		"message_list": getMessage(resp.GetMessageList())},
	)
	return
}

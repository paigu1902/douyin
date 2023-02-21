// Code generated by hertz generator.

package UserCommPb

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	dyUtils "paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
	"strconv"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
)

type CommentActionReq struct {
	CommentId   string `query:"comment_id"`
	VideoId     string `query:"video_id"`
	ActionType  string `query:"action_type"`
	CommentText string `query:"comment_text" vd:"len($)>=0 && len($)<100"`
}

type CommentsInfoReq struct {
	UserId  int64 `query:"user_id"`
	VideoId int64 `query:"video_id"`
}

type CommentHttp struct {
	Id         int64     `json:"id"`
	User       *UserHttp `json:"user"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date"`
}

type UserHttp struct {
	UserId          int64  `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count" default:"0"`
	FollowerCount   int64  `json:"follower_count" default:"0"`
	IsFollow        bool   `json:"is_follow" default:"false"`
	Avatar          string `json:"avatar" default:"https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c"`
	BackgroundImage string `json:"background_image" default:"https://img0.baidu.com/it/u=1705694933,4002952892&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1677085200&t=327023c8f454fb913a8a32d5485f403c"`
	Signature       string `json:"signature" default:"666"`
	TotalFavorited  string `json:"total_favorited" default:"0"`
	WorkCount       int64  `json:"work_count" default:"0"`
	FavoriteCount   int64  `json:"favorite_count" default:"0"`
}

func getUserHttp(user *userInfoPb.User) *UserHttp {
	return &UserHttp{
		UserId:          int64(user.GetUserId()),
		UserName:        user.GetUserName(),
		FollowCount:     user.GetFollowCount(),
		FollowerCount:   user.GetFollowerCount(),
		IsFollow:        user.GetIsFollow(),
		Avatar:          user.GetAvatar(),
		BackgroundImage: user.GetBackgroundImage(),
		Signature:       user.GetSignature(),
		TotalFavorited:  strconv.Itoa(int(user.GetTotalFavorited())),
		FavoriteCount:   user.GetFavoriteCount(),
		WorkCount:       user.GetWorkCount(),
	}
}

func getComments(ctx context.Context, comments []*UserCommPb.Comment, FromId int64) ([]*CommentHttp, error) {
	res := make([]*CommentHttp, len(comments))
	ids := make([]uint64, len(comments))
	//log.Println(comments)
	for i, v := range comments {
		ids[i] = uint64(v.GetUser().GetId())
	}
	Info_resp, err := rpcClient.UserInfo.BatchInfo(ctx, &userInfoPb.BatchUserReq{
		Fromid:   uint64(FromId),
		Batchids: ids,
	})
	if err != nil {
		return res, err
	}
	for i, v := range comments {
		res[i] = &CommentHttp{
			Id: v.GetId(),
			User: getUserHttp(Info_resp.GetBatchusers()[i]),
			Content:    v.GetContent(),
			CreateDate: v.GetCreateDate(),
		}
	}
	return res, nil
}

func CommentActionMethod(ctx context.Context, c *app.RequestContext) {
	req := new(CommentActionReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	id, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err.Error())
	}
	userId := id

	actionType, _ := strconv.Atoi(req.ActionType)
	videoId, _ := strconv.Atoi(req.VideoId)
	commentId, _ := strconv.Atoi(req.CommentId)
	resp, err := rpcClient.UserComm.CommentAction(ctx, &UserCommPb.DouyinCommentActionRequest{
		UserId:      int64(userId),
		VideoId:     int64(videoId),
		ActionType:  int32(actionType),
		CommentText: req.CommentText,
		CommentId:   int64(commentId),
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"comment":     resp.GetComment(),
	})
	return
}

func CommentGetListMethod(ctx context.Context, c *app.RequestContext) {
	req := new(CommentsInfoReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	id, err := dyUtils.GetFromId(c)
	if err != nil {
		c.JSON(400, err.Error())
	}
	userId := id
	resp, err := rpcClient.UserComm.GetCommentsByVideo(ctx, &UserCommPb.DouyinCommentListRequest{
		UserId:  int64(userId),
		VideoId: req.VideoId,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	commenList, err := getComments(ctx, resp.GetCommentList(), req.UserId)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code":  resp.GetStatusCode(),
		"status_msg":   resp.GetStatusMsg(),
		"comment_list": commenList,
	})
	return
}

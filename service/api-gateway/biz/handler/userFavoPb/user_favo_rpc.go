// Code generated by hertz generator.

package userFavo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
)

type FavoActionReq struct {
	UserId int64 `query:"userId"`
	VideId int64 `query:"videoId"`
	Type   int32 `query:"type" `
}

type FavoListReq struct {
	UserId int64 `query:"userId"`
}

type FavoStatusReq struct {
	UserId int64 `query:"userId"`
	VideId int64 `query:"videoId"`
}

type UserHttp struct {
	UserId          uint64 `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar" default:""`
	BackgroundImage string `json:"background_image" default:""`
	Signature       string `json:"signature" default:""`
	TotalFavorited  string `json:"total_favorited" default:""`
	WorkCount       int64  `json:"work_count" default:"0"`
	FavoriteCount   int64  `json:"favorite_count" default:"0"`
}

type VideoListHttp struct {
	Id            uint64    `json:"id"`
	User          *UserHttp `json:"user"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
}

func GetVideoList(videolist []*userFavoPb.Video) []*VideoListHttp {
	res := make([]*VideoListHttp, len(videolist))
	for i, v := range videolist {
		res[i] = &VideoListHttp{
			Id: v.Id,
			User: &UserHttp{
				UserId:        v.GetAuthor().GetId(),
				UserName:      v.GetAuthor().GetName(),
				FollowCount:   v.GetAuthor().GetFollowCount(),
				FollowerCount: v.GetAuthor().GetFollowerCount(),
				IsFollow:      v.GetAuthor().GetIsFollow(),
			},
			PlayUrl:       v.GetPlayUrl(),
			CoverUrl:      v.GetCoverUrl(),
			FavoriteCount: v.GetFavoriteCount(),
			CommentCount:  v.GetCommentCount(),
			IsFavorite:    v.GetIsFavorite(),
			Title:         v.GetTitle(),
		}
	}
	return res
}

func FavoActionMethod(ctx context.Context, c *app.RequestContext) {
	req := new(FavoActionReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := rpcClient.UserFavo.FavoAction(ctx, &userFavoPb.FavoActionReq{
		UserId:  req.UserId,
		VideoId: req.VideId,
		Type:    req.Type,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
	})
	return
}

func FavoListMethod(ctx context.Context, c *app.RequestContext) {
	req := new(FavoListReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := rpcClient.UserFavo.FavoList(ctx, &userFavoPb.FavoListReq{
		UserId: req.UserId,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"favo_list":   GetVideoList(resp.GetVideoList()),
	})
	return
}

func FavoStatusMethod(ctx context.Context, c *app.RequestContext) {
	req := new(FavoStatusReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := rpcClient.UserFavo.FavoStatus(ctx, &userFavoPb.FavoStatusReq{
		UserId:  req.UserId,
		VideoId: req.VideId,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"favo_status": resp.GetIsFavorite(),
	})
	return
}

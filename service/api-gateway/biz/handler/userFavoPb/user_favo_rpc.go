// Code generated by hertz generator.

package userFavo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
	"strconv"
)

func FavoActionMethod(ctx context.Context, c *app.RequestContext) {
	userId, err1 := c.GetQuery("userId")
	videoId, err2 := c.GetQuery("VideoId")
	actionType, err3 := c.GetQuery("type")
	if !(err1 && err2 && err3) {
		c.String(400, "Get parameters failed")
		return
	}
	userid, _ := strconv.Atoi(userId)
	videoid, _ := strconv.Atoi(videoId)
	actiontype, _ := strconv.Atoi(actionType)
	resp, err := rpcClient.UserFavo.FavoAction(ctx, &userFavoPb.FavoActionReq{
		UserId:  int64(userid),
		VideoId: int64(videoid),
		Type:    int32(actiontype),
	})
	if err != nil {
		c.String(400, err.Error())
	}
	c.JSON(200, &resp)
	return
}

func FavoListMethod(ctx context.Context, c *app.RequestContext) {
	userId, err1 := c.GetQuery("userId")
	if !err1 {
		c.String(400, "Get parameters failed")
		return
	}
	userid, _ := strconv.Atoi(userId)
	resp, err := rpcClient.UserFavo.FavoList(ctx, &userFavoPb.FavoListReq{
		UserId: int64(userid),
	})
	if err != nil {
		c.String(400, err.Error())
	}
	c.JSON(200, &resp)
	return
}

func FavoStatusMethod(ctx context.Context, c *app.RequestContext) {
	userId, err1 := c.GetQuery("userId")
	videoId, err2 := c.GetQuery("VideoId")
	if !(err1 && err2) {
		c.String(400, "Get parameters failed")
		return
	}
	userid, _ := strconv.Atoi(userId)
	videoid, _ := strconv.Atoi(videoId)
	resp, err := rpcClient.UserFavo.FavoStatus(ctx, &userFavoPb.FavoStatusReq{
		UserId:  int64(userid),
		VideoId: int64(videoid),
	})
	if err != nil {
		c.String(400, err.Error())
	}
	c.JSON(200, &resp)
	return
}
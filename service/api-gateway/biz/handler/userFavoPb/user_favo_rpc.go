// Code generated by hertz generator.

package userFavo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
)

type FavoActionReq struct {
	UserId     int64 `query:"userId"`
	VideId     int64 `query:"videoId"`
	ActionType int32 `query:"type" vd:"$==1 || $==2"`
}

type FavoListReq struct {
	UserId int64 `query:"userId"`
}

type FavoStatusReq struct {
	UserId int64 `query:"userId"`
	VideId int64 `query:"videoId"`
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
		Type:    req.ActionType,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, &resp)
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
	c.JSON(200, &resp)
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
	c.JSON(200, &resp)
	return
}

// Code generated by hertz generator.

package UserCommPb

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
)

type CommentActionReq struct {
	UserId      int64 `query:"user_id"`
	CommentId   int64
	VideId      int64  `query:"video_id"`
	ActionType  int32  `query:"action_type" vd:"$==1 || $==2"`
	CommentText string `vd:"len($)>=0 && len($)<30"`
}

type CommentsInfoReq struct {
	UserId int64 `query:"user_id"`
	VideId int64 `query:"video_id"`
}

func CommentActionMethod(ctx context.Context, c *app.RequestContext) {
	req := new(CommentActionReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if req.ActionType == 2 {
		// 删除评论 comment_text = ""
		req.CommentText = ""
		CommentId, isOk := c.Get("comment_id")
		if !isOk {
			c.JSON(400, "get current comment_id error in delete comment")
			return
		}
		req.CommentId = CommentId.(int64)
	} else {
		// 添加评论 comment_id = 0
		req.CommentId = 0
		CommentText, isOk := c.Get("comment_text")
		if !isOk {
			c.JSON(400, "get current comment_text error in add comment")
			return
		}
		if len(CommentText.(string)) == 0 {
			c.JSON(400, "add_op: the length of comment should not be 0")
			return
		}
		req.CommentText = CommentText.(string)
	}

	resp, err := rpcClient.UserComm.CommentAction(ctx, &UserCommPb.DouyinCommentActionRequest{
		UserId:      req.UserId,
		VideoId:     req.VideId,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, &resp)
	return
}

func CommentGetListMethod(ctx context.Context, c *app.RequestContext) {
	req := new(CommentsInfoReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := rpcClient.UserComm.GetCommentsByVideo(ctx, &UserCommPb.DouyinCommentListRequest{
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

func CommentNumberMethod(ctx context.Context, c *app.RequestContext) {
	req := new(CommentsInfoReq)
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := rpcClient.UserComm.GetCommentNumberByVideo(ctx, &UserCommPb.DouyinCommentNumberRequest{
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

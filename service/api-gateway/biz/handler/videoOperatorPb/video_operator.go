// Code generated by hertz generator.

package videoOperator

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"io"
	"log"
	"mime/multipart"
	dyUtils "paigu1902/douyin/common/utils"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
	"strconv"
)

type VideoReq struct {
	Data  *multipart.FileHeader `form:"data"`
	Token string                `form:"token"`
	Title string                `form:"title"`
}

// PublishListReq get 方法需要标注 query 参数
type PublishListReq struct {
	UserId   uint64
	Token    string `query:"token"`
	AuthorId uint64 `query:"user_id"`
}

type FeedReq struct {
	Token      string `query:"token"`
	LatestTime int64  `query:"latest_time"`
}

func (req *PublishListReq) getGrpcReq() *videoOperatorPb.PublishListReq {
	return &videoOperatorPb.PublishListReq{UserId: req.UserId, AuthorId: req.AuthorId, Token: req.Token}
}

func file2Byte(file *multipart.FileHeader) ([]byte, error) {
	filepoint, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(filepoint multipart.File) {
		err := filepoint.Close()
		if err != nil {
			hlog.Warnf("文件关闭失败%s", err)
		}
	}(filepoint)
	var content []byte
	buf := make([]byte, 1024)
	for {
		n, err := filepoint.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		content = append(content, buf[:n]...)
	}
	return content, nil
}

func PublishActionMethod(ctx context.Context, c *app.RequestContext) {
	var req VideoReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	data, err := file2Byte(req.Data)
	if err != nil {
		panic(err)
	}
	resp, err := rpcClient.VideoOperatorClient.Upload(ctx, &videoOperatorPb.VideoUploadReq{
		Token: req.Token,
		Data:  data,
		Title: req.Title,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp", resp)
	c.JSON(200, utils.H{
		"status_code": resp.GetStatus(),
		"status_msg":  resp.GetStatusMsg(),
	})
}

func FeedMethod(ctx context.Context, c *app.RequestContext) {
	var req FeedReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	user, err := dyUtils.AnalyseToken(req.Token)
	if err != nil {
		req.Token = ""
	} else {
		req.Token = strconv.Itoa(int(user.ID))
	}

	resp, err := rpcClient.VideoOperatorClient.Feed(ctx, &videoOperatorPb.FeedReq{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})
	if err != nil {
		hlog.Error(err)
	}
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"next_time":   resp.GetNextTime(),
		"video_list":  resp.GetVideoList(),
	})
}
func PublishListMethod(ctx context.Context, c *app.RequestContext) {
	req := new(PublishListReq)
	// 1. 绑定校验参数
	if err := c.BindAndValidate(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	// 获取用户id
	fromId, isOK := c.Get("from_id")
	if !isOK {
		//c.JSON(400, "get current user_id error")
		//未登录状态
		fromId = req.AuthorId
	}
	req.UserId = uint64(fromId.(uint64))
	// 2.调用rpc
	resp, err := rpcClient.VideoOpClient.PublishList(ctx, req.getGrpcReq())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, resp)
	return
}

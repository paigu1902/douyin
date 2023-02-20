// Code generated by hertz generator.

package userInfo

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	"paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"strconv"
)

type LoginReq struct {
	UserName string `query:"username"`               //userName
	Password string `query:"password" vd:"login($)"` //password
}

type UserHttp struct {
	UserId          int64  `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count" default:"0"`
	FollowerCount   int64  `json:"follower_count" default:"0"`
	IsFollow        bool   `json:"is_follow" default:"false"`
	Avatar          string `json:"avatar" default:""`
	BackgroundImage string `json:"background_image" default:""`
	Signature       string `json:"signature" default:""`
	TotalFavorited  string `json:"total_favorited" default:""`
	WorkCount       int64  `json:"work_count" default:"0"`
	FavoriteCount   int64  `json:"favorite_count" default:"0"`
}

func getUserHttp(user *userInfoPb.User) *UserHttp {
	return &UserHttp{
		UserId:        int64(user.GetUserId()),
		UserName:      user.GetUserName(),
		FollowCount:   user.GetFollowCount(),
		FollowerCount: user.GetFollowerCount(),
		IsFollow:      user.GetIsFollow(),
	}
}

func init() {
	binding.MustRegValidateFunc("login", func(args ...interface{}) error {
		s, _ := args[0].(string)
		if len(s) < 6 || len(s) > 20 {
			return fmt.Errorf("登陆信息输入格式有误")
		}
		return nil
	})
}
func LoginMethod(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	// 1.绑定参数
	err := c.BindAndValidate(&req)
	if err != nil {
		respErr := &userInfoPb.LoginResp{StatusMsg: err.Error(), StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	// 2.调用rpc
	resp, err := rpcClient.UserInfo.Login(ctx, &userInfoPb.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	// 3.异常处理
	if err != nil {
		respErr := &userInfoPb.LoginResp{StatusMsg: err.Error(), StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	// 4.正常返回
	log.Println("resp", resp)
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user_id":     resp.GetUserId(),
		"token":       resp.GetToken()},
	)
	return

}

func RegisterMethod(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	// 1.绑定参数
	err := c.BindAndValidate(&req)
	if err != nil {
		respErr := &userInfoPb.LoginResp{StatusMsg: err.Error(), StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	// 2.调用rpc
	resp, err := rpcClient.UserInfo.Register(ctx, &userInfoPb.RegisterReq{UserName: req.UserName, Password: req.Password})
	// 3.异常处理
	if err != nil {
		respErr := &userInfoPb.LoginResp{StatusMsg: err.Error(), StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	// 4.正常返回
	log.Println("resp", resp)
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		"user_id":     resp.GetUserId(),
		"token":       resp.GetToken()},
	)
	return
}

func InfoMethod(ctx context.Context, c *app.RequestContext) {
	var req userInfoPb.UserInfoReq
	var isok bool
	fromId, _ := c.Get("from_id")
	userId, isok := c.GetQuery("user_id")
	uid, _ := strconv.Atoi(userId)
	req.ToId = uint64(uid)
	req.FromId = uint64(fromId.(uint))
	if !isok {
		respErr := &userInfoPb.UserInfoResp{StatusMsg: "获取用户参数失败", StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	resp, err := rpcClient.UserInfo.Info(ctx, &req)
	if err != nil {
		respErr := &userInfoPb.UserInfoResp{StatusMsg: err.Error(), StatusCode: 1}
		c.JSON(200, respErr)
		return
	}
	log.Println("resp", resp)
	c.JSON(200, utils.H{
		"status_code": resp.GetStatusCode(),
		"status_msg":  resp.GetStatusMsg(),
		//"user":        getUserHttp(resp.GetUser()),
		"user": utils.H{
			"id":             resp.GetUser().GetUserId(),
			"name":           resp.GetUser().GetUserName(),
			"follow_count":   resp.GetUser().GetFollowCount(),
			"follower_count": resp.GetUser().GetFollowerCount(),
		},
	},
	)
	return
}

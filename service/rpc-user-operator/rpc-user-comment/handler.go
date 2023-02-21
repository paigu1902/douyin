package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"paigu1902/douyin/common/cache"
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/service/api-gateway/biz/rpcClient"
	UserInfo "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
	"strconv"
	"time"
)

// UserCommRpcImpl implements the last service interface defined in the IDL.
type UserCommRpcImpl struct{}

// GetCommentNumberByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentNumberByVideo(ctx context.Context, req *UserCommPb.DouyinCommentNumberRequest) (resp *UserCommPb.DouyinCommentNumberResponse, err error) {
	videoId := req.VideoId
	count, err := models.GetCommentsNumByVideo(videoId)
	if err != nil {
		return &UserCommPb.DouyinCommentNumberResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
			Count:      -1,
		}, err
	}
	return &UserCommPb.DouyinCommentNumberResponse{
		StatusCode: 0,
		StatusMsg:  "SUCCESS",
		Count:      count,
	}, nil
}

// CommentAction implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) CommentAction(ctx context.Context, req *UserCommPb.DouyinCommentActionRequest) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	videoId := req.VideoId
	var videos []models.VideoInfo

	err = models.GetVideosByIds([]uint64{uint64(videoId)}, &videos)
	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 3,
			StatusMsg:  "GET VIDEO AUTHOR ERROR",
		}, errors.New("VIDEO FIND Error")
	}

	userInfo, err := rpcClient.UserInfo.Info(ctx, &UserInfo.UserInfoReq{FromId: uint64(req.UserId), ToId: uint64(req.UserId)})
	user := userInfo.GetUser()
	if err != nil {
		klog.Error("远程调用UserInfo.BatchInfo失败")
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 3,
			StatusMsg:  "GET VIDEO AUTHOR ERROR",
		}, errors.New("VIDEO FIND Error")
	}
	commentTxt := req.CommentText
	commentId := req.CommentId // del用

	comment := UserCommPb.Comment{
		Id: commentId,
		User: &UserCommPb.User{
			Id:            int64(user.UserId),
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
		Content:    commentTxt,
		CreateDate: time.Now().Format("1-2"),
	}
	if req.ActionType == 1 {
		// 发表评论
		return AddComment(ctx, &comment, videoId, commentTxt)
	} else {
		// 删除评论
		return DelComment(ctx, &comment, commentId, videoId)
	}
}

// GetCommentsByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentsByVideo(ctx context.Context, req *UserCommPb.DouyinCommentListRequest) (resp *UserCommPb.DouyinCommentListResponse, err error) {
	videoId := req.VideoId
	var commentList []models.UserComm
	var respCommentList []*UserCommPb.Comment

	var LimitCommentIds []uint
	var comm models.UserComm
	values, err := cache.RdbUserOp.SMembers(ctx, "VideoIdToComments:"+strconv.FormatInt(videoId, 10)).Result()
	if err == nil {
		for _, valI := range values {
			err := json.Unmarshal([]byte(valI), &comm)
			if err == nil {
				klog.Info(comm)
				commentList = append(commentList, comm)
				LimitCommentIds = append(LimitCommentIds, comm.ID)
			}
		}
	}
	var commentList1 []models.UserComm
	//log.Println(LimitCommentIds)
	err = models.GetCommentsByVideo(videoId, &commentList1, &LimitCommentIds)
	commentList = append(commentList, commentList1...)
	if err != nil {
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	}
	respCommentList, err = FillCommentListFields(ctx, commentList, uint64(req.UserId))
	if err != nil {
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode:  1,
			StatusMsg:   "ERROR",
			CommentList: respCommentList,
		}, err
	}
	// redis 更新评论id
	go func() {
		cnt, err := cache.RdbUserOp.SCard(ctx, "VideoIdToComments:"+strconv.FormatInt(videoId, 10)).Result()
		if err != nil {
			klog.Info("get cnt from VC error", err)
		}
		if cnt == int64(len(respCommentList)+1) {
			// 正确，无需更新
			return
		}
		cache.RdbUserOp.FlushAll(ctx)
		_, err = cache.RdbUserOp.SAdd(ctx, "VideoIdToComments:"+strconv.Itoa(int(videoId)), -1).Result()
		if err != nil {
			klog.Error("redis save -1 error")
			return
		}
		//设置key值过期时间
		_, err2 := cache.RdbUserOp.Expire(ctx, "VideoIdToComments:"+strconv.Itoa(int(videoId)),
			time.Duration(60*60*24*30)*time.Second).Result()
		if err2 != nil {
			klog.Info("redis save one vId - c expire failed")
		}
		for _, comment := range commentList {
			InsertRedisComment(ctx, videoId, comment.CommText)
		}
	}()

	return &UserCommPb.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "SUCCESS",
		CommentList: respCommentList,
	}, nil
}

func InsertRedisComment(ctx context.Context, videoId int64, CommentText string) {
	// 在VideoId下存储Comments
	// Redis update
	_, err := cache.RdbUserOp.SAdd(ctx, "VideoIdToComments:"+strconv.FormatInt(videoId, 10), CommentText).Result()
	if err != nil {
		klog.Info("redis save send: vId - cId failed, key deleted")
		cache.RdbUserOp.Del(ctx, "VideoIdToComments:"+strconv.FormatInt(videoId, 10))
		return
	}
}

func FillCommentListFields(ctx context.Context, comments []models.UserComm, userId uint64) ([]*UserCommPb.Comment, error) {
	size := len(comments)
	var commentListPb []*UserCommPb.Comment
	if comments == nil || size == 0 {
		return commentListPb, nil
	}
	var UserIds []uint64
	for _, com := range comments {
		UserIds = append(UserIds, com.UserId)
	}
	myReq := UserInfo.BatchUserReq{
		Batchids: UserIds,
		Fromid:   userId,
	}
	myRes, err := rpcClient.UserInfo.BatchInfo(ctx, &myReq)
	if err != nil {
		return commentListPb, err
	}
	for i, v := range comments {
		user := myRes.GetBatchusers()[i]
		commentListPb = append(commentListPb, &UserCommPb.Comment{
			Id: int64(v.ID),
			User: &UserCommPb.User{
				Id:            int64(user.UserId),
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			Content:    v.CommText,
			CreateDate: v.CreatedAt.Format("1-2"),
		})
	}
	return commentListPb, nil
}

func AddComment(ctx context.Context, comment *UserCommPb.Comment, videoId int64, commentTxt string) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	commentTmp := models.UserComm{
		UserName: comment.User.Name,
		VideoId:  videoId,
		CommText: commentTxt,
		UserId:   uint64(comment.User.Id),
		Status:   1,
	}
	err = models.InsertComment(&commentTmp)
	// 删除VideoInfo:id数据
	_, err = cache.RDB.Del(ctx, "VideoInfo:"+strconv.Itoa(int(videoId))).Result()
	if err != nil {
		klog.Error("del video redis error", videoId)
	}
	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	}
	//将此发表的评论id存入redis
	go func() {
		InsertRedisComment(ctx, videoId, strconv.Itoa(int(commentTmp.ID)))
		klog.Info("send comment save in redis")
	}()
	comment.Id = int64(commentTmp.ID)
	return &UserCommPb.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "SUCCESS",
		Comment:    comment,
	}, nil
}

func DelComment(ctx context.Context, comment *UserCommPb.Comment, commentId int64, videoId int64) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	// 先对redis中去删除
	n, err := cache.RdbUserOp.Exists(ctx, "CommentIdToVideoId:"+strconv.FormatInt(commentId, 10)).Result()
	if err != nil {
		klog.Info(err)
	}
	if n > 0 { // redis 有数据
		del1, err := cache.RdbUserOp.Del(ctx, "CommentIdToVideoId:"+strconv.FormatInt(commentId, 10)).Result()
		if err != nil {
			klog.Info("Del in CV table err", err)
		}
		klog.Info("del comment in Redis success:", del1)
	}
	err = models.DeleteComment(commentId, videoId)
	_, err = cache.RDB.Del(ctx, "VideoInfo:"+strconv.Itoa(int(videoId))).Result()
	if err != nil {
		klog.Error("del video redis error", videoId)
	}
	if err != nil {
		if err.Error() == "del comment is not exist" {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "NOT_EXIST_ERROR",
			}, err
		} else {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 2,
				StatusMsg:  "OTHER_ERROR",
			}, err
		}
	}
	return &UserCommPb.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "SUCCESS",
		Comment:    comment,
	}, nil
}

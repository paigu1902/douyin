package main

import (
	"context"
	"errors"
	"log"
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
	//先在缓存中查
	cnt, err := cache.RdbUserOp.SCard(ctx, "VideoIdToCommentIds:"+strconv.FormatInt(videoId, 10)).Result()
	if err != nil { //若查询缓存出错，则打印log
		log.Println("count from redis error:", err)
	}
	log.Println("count from redis is:", cnt)
	if cnt != 0 {
		return &UserCommPb.DouyinCommentNumberResponse{
			StatusCode: 0,
			StatusMsg:  "SUCCESS",
			Count:      cnt - 1,
		}, nil
	}
	// 缓存查不到就到model里调用函数去查询
	count, err := models.GetCommentsNumByVideo(videoId)
	if err != nil {
		return &UserCommPb.DouyinCommentNumberResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
			Count:      -1,
		}, err
	}
	go func() {
		getList, _ := models.GetCommentList(videoId)
		_, err2 := cache.RdbUserOp.SAdd(ctx, "VideoIdToCommentIds:"+strconv.Itoa(int(videoId)), -1).Result()
		if err2 != nil {
			log.Println("redis save one vId - cId -1 failed")
			return
		}
		// 设置过期时间
		_, err := cache.RdbUserOp.Expire(ctx, "VideoIdToCommentIds:"+strconv.Itoa(int(videoId)),
			time.Duration(60*60*24*30)*time.Second).Result()
		if err != nil {
			log.Println("redis save one vId - cId expire failed")
		}
		// 存入redis
		for _, CommentId := range getList {
			InsertRedisComment(ctx, videoId, CommentId)
		}
		log.Println("save in redis success")
	}()

	return &UserCommPb.DouyinCommentNumberResponse{
		StatusCode: 0,
		StatusMsg:  "SUCCESS",
		Count:      count,
	}, nil
}

// CommentAction implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) CommentAction(ctx context.Context, req *UserCommPb.DouyinCommentActionRequest) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	IDs := []uint64{uint64(req.UserId)}
	videoId := req.VideoId
	var videos []models.VideoInfo

	err = models.GetVideosByIds([]uint64{uint64(videoId)}, &videos)
	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 3,
			StatusMsg:  "GET VIDEO AUTHOR ERROR",
		}, errors.New("VIDEO FIND Error")
	}
	myReq := UserInfo.BatchUserReq{
		Batchids: IDs,
		Fromid:   videos[0].AuthorId,
	}
	getResult, _ := rpcClient.UserInfo.BatchInfo(ctx, &myReq)
	user := getResult.Batchusers[0] // get user
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
		return DelComment(ctx, &comment, commentId)
	}
}

// GetCommentsByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentsByVideo(ctx context.Context, req *UserCommPb.DouyinCommentListRequest) (resp *UserCommPb.DouyinCommentListResponse, err error) {
	videoId := req.VideoId
	commentList, err := models.GetCommentsByVideo(videoId)
	if err != nil {
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	}
	respCommentList, err := FillCommentListFields(commentList, videoId)
	if err != nil {
		// 评论为空，此时应该只是提示，不报错
		if err.Error() == "find list is empty" {
			return &UserCommPb.DouyinCommentListResponse{
				StatusCode:  0,
				StatusMsg:   "SUCCESS BUT NOT_EXIST_LIST",
				CommentList: respCommentList,
			}, nil
		} else {
			return &UserCommPb.DouyinCommentListResponse{
				StatusCode:  1,
				StatusMsg:   "ERROR",
				CommentList: respCommentList,
			}, nil
		}
	}
	// redis 更新评论id
	go func() {
		cnt, err := cache.RdbUserOp.SCard(ctx, "VideoIdToCommentIds:"+strconv.FormatInt(videoId, 10)).Result()
		if err != nil {
			log.Println("get cnt from VC error", err)
		}
		if cnt > 0 {
			// 已经有正确的值在里面了，无需更新
			return
		}
		_, err = cache.RdbUserOp.SAdd(ctx, "VideoIdToCommentIds:"+strconv.Itoa(int(videoId)), -1).Result()
		if err != nil {
			log.Println("redis save -1 error")
			return
		}
		//设置key值过期时间
		_, err2 := cache.RdbUserOp.Expire(ctx, "VideoIdToCommentIds:"+strconv.Itoa(int(videoId)),
			time.Duration(60*60*24*30)*time.Second).Result()
		if err2 != nil {
			log.Println("redis save one vId - cId expire failed")
		}
		for _, _comment := range commentList {
			InsertRedisComment(ctx, videoId, strconv.Itoa(int(_comment.ID)))
		}
	}()
	return &UserCommPb.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "SUCCESS",
		CommentList: respCommentList,
	}, nil
}

func InsertRedisComment(ctx context.Context, videoId int64, CommentId string) {
	// 在VideoId下存储CommentId
	// Redis update
	_, err := cache.RdbUserOp.SAdd(ctx, "VideoIdToCommentIds:"+strconv.FormatInt(videoId, 10), CommentId).Result()
	if err != nil {
		log.Println("redis save send: vId - cId failed, key deleted")
		cache.RdbUserOp.Del(ctx, "VideoIdToCommentIds:"+strconv.FormatInt(videoId, 10))
		return
	}
	// 在CommentId 存储 VideoId
	_, err = cache.RdbUserOp.Set(ctx, "CommentIdToVideoId:"+CommentId, videoId, 0).Result()
	if err != nil {
		log.Println("redis save one cId - vId failed")
	}
}

func FillCommentListFields(comments []models.UserComm, videoId int64) ([]*UserCommPb.Comment, error) {
	size := len(comments)
	var commentListPb []*UserCommPb.Comment
	if comments == nil || size == 0 {
		return commentListPb, errors.New("find list is empty")
	}
	var UserIds []uint64
	for _, com := range comments {
		UserIds = append(UserIds, com.UserId)
	}
	var videos []models.VideoInfo
	err := models.GetVideosByIds([]uint64{uint64(videoId)}, &videos)
	if err != nil {
		return commentListPb, errors.New("get video info by ids error")
	}
	myReq := UserInfo.BatchUserReq{
		Batchids: UserIds,
		Fromid:   videos[0].AuthorId,
	}
	myRes, _ := rpcClient.UserInfo.BatchInfo(context.Background(), &myReq)
	for i, v := range comments {
		user := myRes.Batchusers[i]
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
	err = models.InsertComment(commentTmp)
	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	}
	//将此发表的评论id存入redis
	go func() {
		InsertRedisComment(ctx, videoId, strconv.Itoa(int(commentTmp.ID)))
		log.Println("send comment save in redis")
	}()
	comment.Id = int64(commentTmp.ID)
	return &UserCommPb.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "SUCCESS",
		Comment:    comment,
	}, nil
}

func DelComment(ctx context.Context, comment *UserCommPb.Comment, commentId int64) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	// 先对redis中去删除
	n, err := cache.RdbUserOp.Exists(ctx, "CommentIdToVideoId:"+strconv.FormatInt(commentId, 10)).Result()
	if err != nil {
		log.Println(err)
	}
	if n > 0 { // redis 有数据
		vid, _ := cache.RdbUserOp.Get(ctx, "CommentIdToVideoId:"+strconv.FormatInt(commentId, 10)).Result()
		del1, err := cache.RdbUserOp.Del(ctx, "CommentIdToVideoId:"+strconv.FormatInt(commentId, 10)).Result()
		if err != nil {
			log.Println("Del in CV table err", err)
		}
		del2, err := cache.RdbUserOp.SRem(ctx, "VideoIdToCommentIds:"+vid, strconv.FormatInt(commentId, 10)).Result()
		if err != nil {
			log.Println("Del in VC table err", err)
		}
		log.Println("del comment in Redis success:", del1, del2)
	}
	err = models.DeleteComment(commentId)
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

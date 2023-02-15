package main

import (
	"context"
	"errors"
	"log"
	"paigu1902/douyin/common/utils"
	UserInfo "paigu1902/douyin/service/rpc-user-info/kitex_gen/userInfoPb"
	"paigu1902/douyin/service/rpc-user-info/logic"
	"paigu1902/douyin/service/rpc-user-operator/models"
	"paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/cache"
	UserCommPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-comment/kitex_gen/UserCommPb"
	"strconv"
	"time"
)

// UserCommRpcImpl implements the last service interface defined in the IDL.
type UserCommRpcImpl struct{}

// GetCommentNumberByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentNumberByVideo(ctx context.Context, req *UserCommPb.DouyinCommentNumberRequest) (resp *UserCommPb.DouyinCommentNumberResponse, err error) {
	// TODO: Your code here...
	videoId := req.VideoId
	//先在缓存中查
	cnt, err := cache.RdbVCid.SCard(cache.Ctx, strconv.FormatInt(videoId, 10)).Result()
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
	} else {
		go func() {
			get_list, _ := models.GetCommentList(videoId)
			_, err_2 := cache.RdbVCid.SAdd(cache.Ctx, strconv.Itoa(int(videoId)), -1).Result()
			if err_2 != nil {
				log.Println("redis save one vId - cId -1 failed")
				return
			}
			// 设置过期时间
			_, err := cache.RdbVCid.Expire(cache.Ctx, strconv.Itoa(int(videoId)),
				time.Duration(60*60*24*30)*time.Second).Result()
			if err != nil {
				log.Println("redis save one vId - cId expire failed")
			}
			// 存入redis
			for _, commentid := range get_list {
				InsertRedisComment(videoId, commentid)
			}
			log.Println("save in redis success")
		}()

		return &UserCommPb.DouyinCommentNumberResponse{
			StatusCode: 0,
			StatusMsg:  "SUCCESS",
			Count:      count,
		}, nil
	}
}

// CommentAction implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) CommentAction(ctx context.Context, req *UserCommPb.DouyinCommentActionRequest) (resp *UserCommPb.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	IDs := []uint64{uint64(req.UserId)}
	videoId := req.VideoId
	myReq := UserInfo.BatchUserReq{
		Batchids: IDs,
		Fromid:   uint64((videoId)), // TODO: Find AuthorId
		// 查找videoid的author
	}
	get_result, _ := logic.BatchInfo(ctx, &myReq)
	user := get_result.Batchusers[0] // get user

	// get_result.Batchusers
	//user := UserInfo.FindUserByID(ID)
	commentTxt := req.CommentText
	commentId := req.CommentId // del用

	if err != nil {
		return &UserCommPb.DouyinCommentActionResponse{
			StatusCode: 3,
			StatusMsg:  "USER TOKEN ERROR",
		}, errors.New("Analyse Token Error")
	}
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
		comment_tmp := models.UserComm{
			UserName: user.UserName,
			VideoId:  videoId,
			CommText: commentTxt,
		}
		err := models.InsertComment(comment_tmp)
		if err != nil {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 2,
				StatusMsg:  "OTHER_ERROR",
			}, err
		} else {
			//将此发表的评论id存入redis
			go func() {
				InsertRedisComment(videoId, strconv.Itoa(int(comment_tmp.ID)))
				log.Println("send comment save in redis")
			}()
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 0,
				StatusMsg:  "SUCCESS",
				Comment:    &comment,
			}, nil
		}
	} else {
		// 删除评论
		// 先对redis中去删除
		n, err := cache.RdbCVid.Exists(cache.Ctx, strconv.FormatInt(commentId, 10)).Result()
		if err != nil {
			log.Println(err)
		}
		if n > 0 { // redis 有数据
			vid, _ := cache.RdbCVid.Get(cache.Ctx, strconv.FormatInt(commentId, 10)).Result()
			del1, err := cache.RdbCVid.Del(cache.Ctx, strconv.FormatInt(commentId, 10)).Result()
			if err != nil {
				log.Println("Del in CV table err", err)
			}
			del2, err := cache.RdbVCid.SRem(cache.Ctx, vid, strconv.FormatInt(commentId, 10)).Result()
			if err != nil {
				log.Println("Del in VC table err", err)
			}
			log.Println("del comment in Redis success:", del1, del2)
		}
		err = models.DeleteComment(req.CommentId)
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
		} else {
			return &UserCommPb.DouyinCommentActionResponse{
				StatusCode: 0,
				StatusMsg:  "SUCCESS",
				Comment:    &comment,
			}, nil
		}
	}
}

// GetCommentsByVideo implements the UserCommRpcImpl interface.
func (s *UserCommRpcImpl) GetCommentsByVideo(ctx context.Context, req *UserCommPb.DouyinCommentListRequest) (resp *UserCommPb.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	videoId := req.VideoId
	commentList, err := models.GetCommentsByVideo(videoId)
	if err != nil {
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode: 2,
			StatusMsg:  "OTHER_ERROR",
		}, err
	} else {
		respCommentList, err := utils.FillCommentListFields(commentList)
		if err != nil {
			// 评论为空，此时应该只是提示，不报错
			return &UserCommPb.DouyinCommentListResponse{
				StatusCode: 1,
				StatusMsg:  "NOT_EXIST_LIST",
				//CommentList: UserCommPb.Comment{
				//	Id: comment_list
				//},
				CommentList: respCommentList,
			}, nil
		}
		// redis 更新评论id
		go func() {
			cnt, err := cache.RdbVCid.SCard(cache.Ctx, strconv.FormatInt(videoId, 10)).Result()
			if err != nil {
				log.Println("get cnt from VC error", err)
			}
			if cnt > 0 {
				return
			}
			_, err = cache.RdbVCid.SAdd(cache.Ctx, strconv.Itoa(int(videoId)), -1).Result()
			if err != nil {
				log.Println("redis save -1 error")
				return
			}
			//设置key值过期时间
			_, err2 := cache.RdbVCid.Expire(cache.Ctx, strconv.Itoa(int(videoId)),
				time.Duration(60*60*24*30)*time.Second).Result()
			if err2 != nil {
				log.Println("redis save one vId - cId expire failed")
			}
			for _, _comment := range commentList {
				InsertRedisComment(videoId, strconv.Itoa(int(_comment.ID)))
			}
		}()
		return &UserCommPb.DouyinCommentListResponse{
			StatusCode:  0,
			StatusMsg:   "SUCCESS",
			CommentList: respCommentList,
		}, nil
	}
}

func InsertRedisComment(VideoId int64, CommentId string) {
	// 在VideoId下存储CommentId
	_, err := cache.RdbVCid.SAdd(cache.Ctx, strconv.FormatInt(VideoId, 10), CommentId).Result()
	if err != nil {
		log.Println("redis save send: vId - cId failed, key deleted")
		cache.RdbVCid.Del(cache.Ctx, strconv.FormatInt(VideoId, 10))
		return
	}
	// 在CommentId 存储 VideoId
	_, err = cache.RdbCVid.Set(cache.Ctx, CommentId, VideoId, 0).Result()
	if err != nil {
		log.Println("redis save one cId - vId failed")
	}
}

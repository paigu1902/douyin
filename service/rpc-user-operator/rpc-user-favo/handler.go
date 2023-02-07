package main

import (
	"context"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
)

// UserFavoRpcImpl implements the last service interface defined in the IDL.
type UserFavoRpcImpl struct{}

// FavoAction implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoAction(ctx context.Context, req *userFavoPb.FavoActionReq) (resp *userFavoPb.FavoActionResp, err error) {
	// TODO: Your code here...
    if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
        return nil, errors.New("Parameter Error")
    }
	likeRecord，err := models.GetLikeRecord(req.UserId，req.VideoId)
	if err != nil {
	    err := models. CreateLikeRecord(req.UserId, req.VideoId)
	    if err != nil {
	        return nil,errors.new("Create LikeRecord Failed")
        } 
    }else {
	    if likeRecord.Status == 0 {
	        sts := 1
	    } else {
	        sts := 0
	    }
	    err := models.UpdateLikeStatus(req.UserId, req.VideoId，sts)
	    if err != nil {
	        return nil,new("Update LikeRecord Failed")
	    }
    }
    return &userFavoPb.FavoActionResp{StatusCode: 1, StatusMsg: "Succeed"}, nil
}

// FavoList implements the UserFavoRpcImpl interface.
func (s *UserFavoRpcImpl) FavoList(ctx context.Context, req *userFavoPb.FavoListReq) (resp *userFavoPb.FavoListResp, err error) {
	// TODO: Your code here...
    if req.UserId == 0 {
        return nil, errors.New("Parameter Error")
    }
    videoList, err := GetLikeVideoId(req.UserId)
    if err != nil {
        if err.Error() == "record not found" {
            return nil, err
        } else {
             return nil, errors.New("Get VideoList Failed")
        }
    }
    return &userFavoPb.FavoListResp{StatusCode: 1, StatusMsg: "Succeed", VideoList: videoList}, nil
}

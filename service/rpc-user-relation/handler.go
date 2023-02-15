package main

import (
	"context"
	userRelationPb "paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"paigu1902/douyin/service/rpc-user-relation/logic"
)

// UserRelationImpl implements the last service interface defined in the IDL.
type UserRelationImpl struct{}

// FollowAction implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowAction(ctx context.Context, req *userRelationPb.FollowActionReq) (resp *userRelationPb.FollowActionResp, err error) {
	// TODO: Your code here...
	return logic.FollowAction(ctx, req)
}

// FollowList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowList(ctx context.Context, req *userRelationPb.FollowListReq) (resp *userRelationPb.FollowListResp, err error) {
	// TODO: Your code here...
	return logic.FollowList(req)
}

// FollowerList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowerList(ctx context.Context, req *userRelationPb.FollowerListReq) (resp *userRelationPb.FollowerListResp, err error) {
	// TODO: Your code here...
	return logic.FollowerList(req)
}

// FriendList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FriendList(ctx context.Context, req *userRelationPb.FriendListReq) (resp *userRelationPb.FriendListResp, err error) {
	// TODO: Your code here...
	return logic.FriendList(req)
}

// SendMessage implements the UserRelationImpl interface.
func (s *UserRelationImpl) SendMessage(ctx context.Context, req *userRelationPb.SendMessageReq) (resp *userRelationPb.SendMessageResp, err error) {
	// TODO: Your code here...
	return logic.SendMessage(req)
}

// HistoryMessage implements the UserRelationImpl interface.
func (s *UserRelationImpl) HistoryMessage(ctx context.Context, req *userRelationPb.HistoryMessageReq) (resp *userRelationPb.HistoryMessageResp, err error) {
	// TODO: Your code here...
	return logic.HistoryMessage(req)
}

// IsFollow implements the UserRelationImpl interface.
func (s *UserRelationImpl) IsFollow(ctx context.Context, req *userRelationPb.IsFollowReq) (resp *userRelationPb.IsFollowResp, err error) {
	// TODO: Your code here...
	return logic.IsFollow(ctx, req)
}

// IsFollowList implements the UserRelationImpl interface.
func (s *UserRelationImpl) IsFollowList(ctx context.Context, req *userRelationPb.IsFollowListReq) (resp *userRelationPb.IsFollowListResp, err error) {
	// TODO: Your code here...
	return logic.IsFollowList(ctx, req)
}

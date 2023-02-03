package main

import (
	"context"
	userRelationPb "paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
)

// UserRelationImpl implements the last service interface defined in the IDL.
type UserRelationImpl struct{}

// FollowAction implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowAction(ctx context.Context, req *userRelationPb.FollowActionReq) (resp *userRelationPb.FollowActionResp, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowList(ctx context.Context, req *userRelationPb.FollowListReq) (resp *userRelationPb.FollowListResp, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FollowerList(ctx context.Context, req *userRelationPb.FollowerListReq) (resp *userRelationPb.FollowerListResp, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserRelationImpl interface.
func (s *UserRelationImpl) FriendList(ctx context.Context, req *userRelationPb.FriendListReq) (resp *userRelationPb.FriendListResp, err error) {
	// TODO: Your code here...
	return
}

// SendMessage implements the UserRelationImpl interface.
func (s *UserRelationImpl) SendMessage(ctx context.Context, req *userRelationPb.SendMessageReq) (resp *userRelationPb.SendMessageResp, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the UserRelationImpl interface.
func (s *UserRelationImpl) MessageChat(ctx context.Context, req *userRelationPb.MessageChatReq) (resp *userRelationPb.MessageChatResp, err error) {
	// TODO: Your code here...
	return
}

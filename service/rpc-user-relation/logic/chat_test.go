package logic

import (
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
	"reflect"
	"testing"
)

func TestFollowAction(t *testing.T) {
	type args struct {
		req *userRelationPb.FollowActionReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.FollowActionResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := FollowAction(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FollowAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFollowList(t *testing.T) {
	type args struct {
		req *userRelationPb.FollowListReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.FollowListResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := FollowList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FollowList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFollowerList(t *testing.T) {
	type args struct {
		req *userRelationPb.FollowerListReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.FollowerListResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := FollowerList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FollowerList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFriendList(t *testing.T) {
	type args struct {
		req *userRelationPb.FriendListReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.FriendListResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := FriendList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FriendList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHistoryMessage(t *testing.T) {
	type args struct {
		req *userRelationPb.HistoryMessageReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.HistoryMessageResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := HistoryMessage(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoryMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HistoryMessage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestSendMessage(t *testing.T) {
	type args struct {
		req *userRelationPb.SendMessageReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *userRelationPb.SendMessageResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := SendMessage(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendMessage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_followIds(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		wantIds []uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := followIds(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("followIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIds, tt.wantIds) {
				t.Errorf("followIds() gotIds = %v, want %v", gotIds, tt.wantIds)
			}
		})
	}
}

func Test_followerIds(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		wantIds []uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := followerIds(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("followerIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIds, tt.wantIds) {
				t.Errorf("followerIds() gotIds = %v, want %v", gotIds, tt.wantIds)
			}
		})
	}
}

func Test_idsToMap(t *testing.T) {
	type args struct {
		ids []uint64
	}
	tests := []struct {
		name string
		args args
		want map[uint64]struct{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idsToMap(tt.args.ids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idsToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFollow(t *testing.T) {
	type args struct {
		followMap map[uint64]struct{}
		id        uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFollow(tt.args.followMap, tt.args.id); got != tt.want {
				t.Errorf("isFollow() = %v, want %v", got, tt.want)
			}
		})
	}
}

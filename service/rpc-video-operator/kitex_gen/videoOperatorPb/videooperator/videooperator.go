// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videooperator

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	videoOperatorPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoOperatorServiceInfo
}

var videoOperatorServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoOperator"
	handlerType := (*videoOperatorPb.VideoOperator)(nil)
	methods := map[string]kitex.MethodInfo{
		"Upload":      kitex.NewMethodInfo(uploadHandler, newUploadArgs, newUploadResult, false),
		"Feed":        kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"PublishList": kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
		"VideoList":   kitex.NewMethodInfo(videoListHandler, newVideoListArgs, newVideoListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "videoOperatorPb",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func uploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoOperatorPb.VideoUploadReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoOperatorPb.VideoOperator).Upload(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UploadArgs:
		success, err := handler.(videoOperatorPb.VideoOperator).Upload(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UploadResult)
		realResult.Success = success
	}
	return nil
}
func newUploadArgs() interface{} {
	return &UploadArgs{}
}

func newUploadResult() interface{} {
	return &UploadResult{}
}

type UploadArgs struct {
	Req *videoOperatorPb.VideoUploadReq
}

func (p *UploadArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videoOperatorPb.VideoUploadReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UploadArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UploadArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UploadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UploadArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UploadArgs) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.VideoUploadReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UploadArgs_Req_DEFAULT *videoOperatorPb.VideoUploadReq

func (p *UploadArgs) GetReq() *videoOperatorPb.VideoUploadReq {
	if !p.IsSetReq() {
		return UploadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadArgs) IsSetReq() bool {
	return p.Req != nil
}

type UploadResult struct {
	Success *videoOperatorPb.VideoUploadResp
}

var UploadResult_Success_DEFAULT *videoOperatorPb.VideoUploadResp

func (p *UploadResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videoOperatorPb.VideoUploadResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UploadResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UploadResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UploadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UploadResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UploadResult) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.VideoUploadResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadResult) GetSuccess() *videoOperatorPb.VideoUploadResp {
	if !p.IsSetSuccess() {
		return UploadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoOperatorPb.VideoUploadResp)
}

func (p *UploadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoOperatorPb.FeedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoOperatorPb.VideoOperator).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(videoOperatorPb.VideoOperator).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *videoOperatorPb.FeedReq
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videoOperatorPb.FeedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.FeedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *videoOperatorPb.FeedReq

func (p *FeedArgs) GetReq() *videoOperatorPb.FeedReq {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *videoOperatorPb.FeedResp
}

var FeedResult_Success_DEFAULT *videoOperatorPb.FeedResp

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videoOperatorPb.FeedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.FeedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *videoOperatorPb.FeedResp {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoOperatorPb.FeedResp)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoOperatorPb.PublishListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoOperatorPb.VideoOperator).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(videoOperatorPb.VideoOperator).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *videoOperatorPb.PublishListReq
}

func (p *PublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videoOperatorPb.PublishListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.PublishListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *videoOperatorPb.PublishListReq

func (p *PublishListArgs) GetReq() *videoOperatorPb.PublishListReq {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *videoOperatorPb.PublishListResp
}

var PublishListResult_Success_DEFAULT *videoOperatorPb.PublishListResp

func (p *PublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videoOperatorPb.PublishListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.PublishListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *videoOperatorPb.PublishListResp {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoOperatorPb.PublishListResp)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func videoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoOperatorPb.VideoListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoOperatorPb.VideoOperator).VideoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *VideoListArgs:
		success, err := handler.(videoOperatorPb.VideoOperator).VideoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VideoListResult)
		realResult.Success = success
	}
	return nil
}
func newVideoListArgs() interface{} {
	return &VideoListArgs{}
}

func newVideoListResult() interface{} {
	return &VideoListResult{}
}

type VideoListArgs struct {
	Req *videoOperatorPb.VideoListReq
}

func (p *VideoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videoOperatorPb.VideoListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VideoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VideoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VideoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in VideoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *VideoListArgs) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.VideoListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VideoListArgs_Req_DEFAULT *videoOperatorPb.VideoListReq

func (p *VideoListArgs) GetReq() *videoOperatorPb.VideoListReq {
	if !p.IsSetReq() {
		return VideoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VideoListArgs) IsSetReq() bool {
	return p.Req != nil
}

type VideoListResult struct {
	Success *videoOperatorPb.VideoListResp
}

var VideoListResult_Success_DEFAULT *videoOperatorPb.VideoListResp

func (p *VideoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videoOperatorPb.VideoListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VideoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VideoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VideoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in VideoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *VideoListResult) Unmarshal(in []byte) error {
	msg := new(videoOperatorPb.VideoListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VideoListResult) GetSuccess() *videoOperatorPb.VideoListResp {
	if !p.IsSetSuccess() {
		return VideoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VideoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoOperatorPb.VideoListResp)
}

func (p *VideoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Upload(ctx context.Context, Req *videoOperatorPb.VideoUploadReq) (r *videoOperatorPb.VideoUploadResp, err error) {
	var _args UploadArgs
	_args.Req = Req
	var _result UploadResult
	if err = p.c.Call(ctx, "Upload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Feed(ctx context.Context, Req *videoOperatorPb.FeedReq) (r *videoOperatorPb.FeedResp, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *videoOperatorPb.PublishListReq) (r *videoOperatorPb.PublishListResp, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoList(ctx context.Context, Req *videoOperatorPb.VideoListReq) (r *videoOperatorPb.VideoListResp, err error) {
	var _args VideoListArgs
	_args.Req = Req
	var _result VideoListResult
	if err = p.c.Call(ctx, "VideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

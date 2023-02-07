// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userfavorpc

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	userFavoPb "paigu1902/douyin/service/rpc-user-operator/rpc-user-favo/kitex_gen/userFavoPb"
)

func serviceInfo() *kitex.ServiceInfo {
	return userFavoRpcServiceInfo
}

var userFavoRpcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserFavoRpc"
	handlerType := (*userFavoPb.UserFavoRpc)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoAction": kitex.NewMethodInfo(favoActionHandler, newFavoActionArgs, newFavoActionResult, false),
		"FavoList":   kitex.NewMethodInfo(favoListHandler, newFavoListArgs, newFavoListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "userFavoPb",
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

func favoActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userFavoPb.FavoActionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userFavoPb.UserFavoRpc).FavoAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoActionArgs:
		success, err := handler.(userFavoPb.UserFavoRpc).FavoAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoActionArgs() interface{} {
	return &FavoActionArgs{}
}

func newFavoActionResult() interface{} {
	return &FavoActionResult{}
}

type FavoActionArgs struct {
	Req *userFavoPb.FavoActionReq
}

func (p *FavoActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userFavoPb.FavoActionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoActionArgs) Unmarshal(in []byte) error {
	msg := new(userFavoPb.FavoActionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoActionArgs_Req_DEFAULT *userFavoPb.FavoActionReq

func (p *FavoActionArgs) GetReq() *userFavoPb.FavoActionReq {
	if !p.IsSetReq() {
		return FavoActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoActionResult struct {
	Success *userFavoPb.FavoActionResp
}

var FavoActionResult_Success_DEFAULT *userFavoPb.FavoActionResp

func (p *FavoActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userFavoPb.FavoActionResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoActionResult) Unmarshal(in []byte) error {
	msg := new(userFavoPb.FavoActionResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoActionResult) GetSuccess() *userFavoPb.FavoActionResp {
	if !p.IsSetSuccess() {
		return FavoActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*userFavoPb.FavoActionResp)
}

func (p *FavoActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userFavoPb.FavoListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userFavoPb.UserFavoRpc).FavoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoListArgs:
		success, err := handler.(userFavoPb.UserFavoRpc).FavoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoListArgs() interface{} {
	return &FavoListArgs{}
}

func newFavoListResult() interface{} {
	return &FavoListResult{}
}

type FavoListArgs struct {
	Req *userFavoPb.FavoListReq
}

func (p *FavoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userFavoPb.FavoListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoListArgs) Unmarshal(in []byte) error {
	msg := new(userFavoPb.FavoListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoListArgs_Req_DEFAULT *userFavoPb.FavoListReq

func (p *FavoListArgs) GetReq() *userFavoPb.FavoListReq {
	if !p.IsSetReq() {
		return FavoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoListResult struct {
	Success *userFavoPb.FavoListResp
}

var FavoListResult_Success_DEFAULT *userFavoPb.FavoListResp

func (p *FavoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userFavoPb.FavoListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoListResult) Unmarshal(in []byte) error {
	msg := new(userFavoPb.FavoListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoListResult) GetSuccess() *userFavoPb.FavoListResp {
	if !p.IsSetSuccess() {
		return FavoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*userFavoPb.FavoListResp)
}

func (p *FavoListResult) IsSetSuccess() bool {
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

func (p *kClient) FavoAction(ctx context.Context, Req *userFavoPb.FavoActionReq) (r *userFavoPb.FavoActionResp, err error) {
	var _args FavoActionArgs
	_args.Req = Req
	var _result FavoActionResult
	if err = p.c.Call(ctx, "FavoAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoList(ctx context.Context, Req *userFavoPb.FavoListReq) (r *userFavoPb.FavoListResp, err error) {
	var _args FavoListArgs
	_args.Req = Req
	var _result FavoListResult
	if err = p.c.Call(ctx, "FavoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

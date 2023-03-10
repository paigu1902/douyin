// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videooperator

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	videoOperatorPb "paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Upload(ctx context.Context, Req *videoOperatorPb.VideoUploadReq, callOptions ...callopt.Option) (r *videoOperatorPb.VideoUploadResp, err error)
	Feed(ctx context.Context, Req *videoOperatorPb.FeedReq, callOptions ...callopt.Option) (r *videoOperatorPb.FeedResp, err error)
	PublishList(ctx context.Context, Req *videoOperatorPb.PublishListReq, callOptions ...callopt.Option) (r *videoOperatorPb.PublishListResp, err error)
	VideoList(ctx context.Context, Req *videoOperatorPb.VideoListReq, callOptions ...callopt.Option) (r *videoOperatorPb.VideoListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoOperatorClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoOperatorClient struct {
	*kClient
}

func (p *kVideoOperatorClient) Upload(ctx context.Context, Req *videoOperatorPb.VideoUploadReq, callOptions ...callopt.Option) (r *videoOperatorPb.VideoUploadResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Upload(ctx, Req)
}

func (p *kVideoOperatorClient) Feed(ctx context.Context, Req *videoOperatorPb.FeedReq, callOptions ...callopt.Option) (r *videoOperatorPb.FeedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kVideoOperatorClient) PublishList(ctx context.Context, Req *videoOperatorPb.PublishListReq, callOptions ...callopt.Option) (r *videoOperatorPb.PublishListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, Req)
}

func (p *kVideoOperatorClient) VideoList(ctx context.Context, Req *videoOperatorPb.VideoListReq, callOptions ...callopt.Option) (r *videoOperatorPb.VideoListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoList(ctx, Req)
}

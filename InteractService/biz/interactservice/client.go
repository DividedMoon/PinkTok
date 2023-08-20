// Code generated by Kitex v0.7.0. DO NOT EDIT.

package interactservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"interact_service/biz"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, Req *biz.FavoriteActionReq, callOptions ...callopt.Option) (r *biz.FavoriteActionResp, err error)
	FavoriteList(ctx context.Context, Req *biz.FavoriteListReq, callOptions ...callopt.Option) (r *biz.FavoriteListResp, err error)
	CommentAction(ctx context.Context, Req *biz.CommentActionReq, callOptions ...callopt.Option) (r *biz.CommentActionResp, err error)
	CommentList(ctx context.Context, Req *biz.CommentListReq, callOptions ...callopt.Option) (r *biz.CommentListResp, err error)
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
	return &kInteractServiceClient{
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

type kInteractServiceClient struct {
	*kClient
}

func (p *kInteractServiceClient) FavoriteAction(ctx context.Context, Req *biz.FavoriteActionReq, callOptions ...callopt.Option) (r *biz.FavoriteActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, Req)
}

func (p *kInteractServiceClient) FavoriteList(ctx context.Context, Req *biz.FavoriteListReq, callOptions ...callopt.Option) (r *biz.FavoriteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}

func (p *kInteractServiceClient) CommentAction(ctx context.Context, Req *biz.CommentActionReq, callOptions ...callopt.Option) (r *biz.CommentActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kInteractServiceClient) CommentList(ctx context.Context, Req *biz.CommentListReq, callOptions ...callopt.Option) (r *biz.CommentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}

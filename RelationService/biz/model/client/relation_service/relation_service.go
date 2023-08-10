// Code generated by hertz generator.

package relation_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"relation_service/biz/model/client"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	SendRelationAction(context context.Context, req *client.RelationActionReq, reqOpt ...config.RequestOption) (resp *client.RelationActionResp, rawResponse *protocol.Response, err error)

	GetFollowList(context context.Context, req *client.RelationFollowListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowListResp, rawResponse *protocol.Response, err error)

	GetFollowerList(context context.Context, req *client.RelationFollowerListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowerListResp, rawResponse *protocol.Response, err error)

	GetFriendList(context context.Context, req *client.RelationFriendListReq, reqOpt ...config.RequestOption) (resp *client.RelationFriendListResp, rawResponse *protocol.Response, err error)
}

type RelationServiceClient struct {
	client *cli
}

func NewRelationServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &RelationServiceClient{
		client: cli,
	}, nil
}

func (s *RelationServiceClient) SendRelationAction(context context.Context, req *client.RelationActionReq, reqOpt ...config.RequestOption) (resp *client.RelationActionResp, rawResponse *protocol.Response, err error) {
	httpResp := &client.RelationActionResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/internal/relation/action")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *RelationServiceClient) GetFollowList(context context.Context, req *client.RelationFollowListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowListResp, rawResponse *protocol.Response, err error) {
	httpResp := &client.RelationFollowListResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/internal/relation/follow/list")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *RelationServiceClient) GetFollowerList(context context.Context, req *client.RelationFollowerListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowerListResp, rawResponse *protocol.Response, err error) {
	httpResp := &client.RelationFollowerListResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/internal/relation/follower/list")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *RelationServiceClient) GetFriendList(context context.Context, req *client.RelationFriendListReq, reqOpt ...config.RequestOption) (resp *client.RelationFriendListResp, rawResponse *protocol.Response, err error) {
	httpResp := &client.RelationFriendListResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/internal/relation/friend/list")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewRelationServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewRelationServiceClient("", ops...)
	return
}

func SendRelationAction(context context.Context, req *client.RelationActionReq, reqOpt ...config.RequestOption) (resp *client.RelationActionResp, rawResponse *protocol.Response, err error) {
	return defaultClient.SendRelationAction(context, req, reqOpt...)
}

func GetFollowList(context context.Context, req *client.RelationFollowListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowListResp, rawResponse *protocol.Response, err error) {
	return defaultClient.GetFollowList(context, req, reqOpt...)
}

func GetFollowerList(context context.Context, req *client.RelationFollowerListReq, reqOpt ...config.RequestOption) (resp *client.RelationFollowerListResp, rawResponse *protocol.Response, err error) {
	return defaultClient.GetFollowerList(context, req, reqOpt...)
}

func GetFriendList(context context.Context, req *client.RelationFriendListReq, reqOpt ...config.RequestOption) (resp *client.RelationFriendListResp, rawResponse *protocol.Response, err error) {
	return defaultClient.GetFriendList(context, req, reqOpt...)
}

// Code generated by hertz generator.

package feed_service

import (
	"context"
	"fmt"

	dto "feed_service/biz/model/dto"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	Feed(context context.Context, req *dto.FeedReq, reqOpt ...config.RequestOption) (resp *dto.FeedResp, rawResponse *protocol.Response, err error)
}

type FeedServiceClient struct {
	client *cli
}

func NewFeedServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &FeedServiceClient{
		client: cli,
	}, nil
}

func (s *FeedServiceClient) Feed(context context.Context, req *dto.FeedReq, reqOpt ...config.RequestOption) (resp *dto.FeedResp, rawResponse *protocol.Response, err error) {
	httpResp := &dto.FeedResp{}
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
		execute("GET", "/internal/feed")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewFeedServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewFeedServiceClient("", ops...)
	return
}

func Feed(context context.Context, req *dto.FeedReq, reqOpt ...config.RequestOption) (resp *dto.FeedResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Feed(context, req, reqOpt...)
}

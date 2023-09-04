// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	"video_service/biz"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*biz.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":                 kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"PublishVideo":         kitex.NewMethodInfo(publishVideoHandler, newPublishVideoArgs, newPublishVideoResult, false),
		"GetPublishList":       kitex.NewMethodInfo(getPublishListHandler, newGetPublishListArgs, newGetPublishListResult, false),
		"GetFavoriteVideoList": kitex.NewMethodInfo(getFavoriteVideoListHandler, newGetFavoriteVideoListArgs, newGetFavoriteVideoListResult, false),
		"FavoriteAction":       kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video_service",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.FeedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.VideoService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(biz.VideoService).Feed(ctx, s.Req)
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
	Req *biz.FeedReq
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.FeedReq)
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
	msg := new(biz.FeedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *biz.FeedReq

func (p *FeedArgs) GetReq() *biz.FeedReq {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FeedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FeedResult struct {
	Success *biz.FeedResp
}

var FeedResult_Success_DEFAULT *biz.FeedResp

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.FeedResp)
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
	msg := new(biz.FeedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *biz.FeedResp {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.FeedResp)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FeedResult) GetResult() interface{} {
	return p.Success
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.PublishReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.VideoService).PublishVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishVideoArgs:
		success, err := handler.(biz.VideoService).PublishVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishVideoResult)
		realResult.Success = success
	}
	return nil
}
func newPublishVideoArgs() interface{} {
	return &PublishVideoArgs{}
}

func newPublishVideoResult() interface{} {
	return &PublishVideoResult{}
}

type PublishVideoArgs struct {
	Req *biz.PublishReq
}

func (p *PublishVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.PublishReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishVideoArgs) Unmarshal(in []byte) error {
	msg := new(biz.PublishReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishVideoArgs_Req_DEFAULT *biz.PublishReq

func (p *PublishVideoArgs) GetReq() *biz.PublishReq {
	if !p.IsSetReq() {
		return PublishVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PublishVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PublishVideoResult struct {
	Success *biz.PublishResp
}

var PublishVideoResult_Success_DEFAULT *biz.PublishResp

func (p *PublishVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.PublishResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishVideoResult) Unmarshal(in []byte) error {
	msg := new(biz.PublishResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishVideoResult) GetSuccess() *biz.PublishResp {
	if !p.IsSetSuccess() {
		return PublishVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.PublishResp)
}

func (p *PublishVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PublishVideoResult) GetResult() interface{} {
	return p.Success
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.GetPublishListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.VideoService).GetPublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPublishListArgs:
		success, err := handler.(biz.VideoService).GetPublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPublishListResult)
		realResult.Success = success
	}
	return nil
}
func newGetPublishListArgs() interface{} {
	return &GetPublishListArgs{}
}

func newGetPublishListResult() interface{} {
	return &GetPublishListResult{}
}

type GetPublishListArgs struct {
	Req *biz.GetPublishListReq
}

func (p *GetPublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.GetPublishListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPublishListArgs) Unmarshal(in []byte) error {
	msg := new(biz.GetPublishListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPublishListArgs_Req_DEFAULT *biz.GetPublishListReq

func (p *GetPublishListArgs) GetReq() *biz.GetPublishListReq {
	if !p.IsSetReq() {
		return GetPublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPublishListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPublishListResult struct {
	Success *biz.GetPublishListResp
}

var GetPublishListResult_Success_DEFAULT *biz.GetPublishListResp

func (p *GetPublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.GetPublishListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPublishListResult) Unmarshal(in []byte) error {
	msg := new(biz.GetPublishListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPublishListResult) GetSuccess() *biz.GetPublishListResp {
	if !p.IsSetSuccess() {
		return GetPublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.GetPublishListResp)
}

func (p *GetPublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPublishListResult) GetResult() interface{} {
	return p.Success
}

func getFavoriteVideoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.GetFavoriteVideoListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.VideoService).GetFavoriteVideoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteVideoListArgs:
		success, err := handler.(biz.VideoService).GetFavoriteVideoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteVideoListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteVideoListArgs() interface{} {
	return &GetFavoriteVideoListArgs{}
}

func newGetFavoriteVideoListResult() interface{} {
	return &GetFavoriteVideoListResult{}
}

type GetFavoriteVideoListArgs struct {
	Req *biz.GetFavoriteVideoListReq
}

func (p *GetFavoriteVideoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.GetFavoriteVideoListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteVideoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteVideoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteVideoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteVideoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteVideoListArgs) Unmarshal(in []byte) error {
	msg := new(biz.GetFavoriteVideoListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteVideoListArgs_Req_DEFAULT *biz.GetFavoriteVideoListReq

func (p *GetFavoriteVideoListArgs) GetReq() *biz.GetFavoriteVideoListReq {
	if !p.IsSetReq() {
		return GetFavoriteVideoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteVideoListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFavoriteVideoListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFavoriteVideoListResult struct {
	Success *biz.GetFavoriteVideoListResp
}

var GetFavoriteVideoListResult_Success_DEFAULT *biz.GetFavoriteVideoListResp

func (p *GetFavoriteVideoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.GetFavoriteVideoListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteVideoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteVideoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteVideoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteVideoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteVideoListResult) Unmarshal(in []byte) error {
	msg := new(biz.GetFavoriteVideoListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteVideoListResult) GetSuccess() *biz.GetFavoriteVideoListResp {
	if !p.IsSetSuccess() {
		return GetFavoriteVideoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteVideoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.GetFavoriteVideoListResp)
}

func (p *GetFavoriteVideoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFavoriteVideoListResult) GetResult() interface{} {
	return p.Success
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.FavoriteActionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.VideoService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(biz.VideoService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *biz.FavoriteActionReq
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.FavoriteActionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(biz.FavoriteActionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *biz.FavoriteActionReq

func (p *FavoriteActionArgs) GetReq() *biz.FavoriteActionReq {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *biz.FavoriteActionResp
}

var FavoriteActionResult_Success_DEFAULT *biz.FavoriteActionResp

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.FavoriteActionResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(biz.FavoriteActionResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *biz.FavoriteActionResp {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.FavoriteActionResp)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, Req *biz.FeedReq) (r *biz.FeedResp, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishVideo(ctx context.Context, Req *biz.PublishReq) (r *biz.PublishResp, err error) {
	var _args PublishVideoArgs
	_args.Req = Req
	var _result PublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishList(ctx context.Context, Req *biz.GetPublishListReq) (r *biz.GetPublishListResp, err error) {
	var _args GetPublishListArgs
	_args.Req = Req
	var _result GetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteVideoList(ctx context.Context, Req *biz.GetFavoriteVideoListReq) (r *biz.GetFavoriteVideoListResp, err error) {
	var _args GetFavoriteVideoListArgs
	_args.Req = Req
	var _result GetFavoriteVideoListResult
	if err = p.c.Call(ctx, "GetFavoriteVideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *biz.FavoriteActionReq) (r *biz.FavoriteActionResp, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

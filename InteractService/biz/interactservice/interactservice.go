// Code generated by Kitex v0.6.2. DO NOT EDIT.

package interactservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	"interact_service/biz"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactServiceServiceInfo
}

var interactServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractService"
	handlerType := (*biz.InteractService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":     kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"QueryFavoriteExist": kitex.NewMethodInfo(queryFavoriteExistHandler, newQueryFavoriteExistArgs, newQueryFavoriteExistResult, false),
		"CommentAction":      kitex.NewMethodInfo(commentActionHandler, newCommentActionArgs, newCommentActionResult, false),
		"CommentList":        kitex.NewMethodInfo(commentListHandler, newCommentListArgs, newCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interact_service",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.FavoriteActionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.InteractService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(biz.InteractService).FavoriteAction(ctx, s.Req)
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

func queryFavoriteExistHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.QueryFavoriteExistReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.InteractService).QueryFavoriteExist(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryFavoriteExistArgs:
		success, err := handler.(biz.InteractService).QueryFavoriteExist(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryFavoriteExistResult)
		realResult.Success = success
	}
	return nil
}
func newQueryFavoriteExistArgs() interface{} {
	return &QueryFavoriteExistArgs{}
}

func newQueryFavoriteExistResult() interface{} {
	return &QueryFavoriteExistResult{}
}

type QueryFavoriteExistArgs struct {
	Req *biz.QueryFavoriteExistReq
}

func (p *QueryFavoriteExistArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.QueryFavoriteExistReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryFavoriteExistArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryFavoriteExistArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryFavoriteExistArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryFavoriteExistArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryFavoriteExistArgs) Unmarshal(in []byte) error {
	msg := new(biz.QueryFavoriteExistReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryFavoriteExistArgs_Req_DEFAULT *biz.QueryFavoriteExistReq

func (p *QueryFavoriteExistArgs) GetReq() *biz.QueryFavoriteExistReq {
	if !p.IsSetReq() {
		return QueryFavoriteExistArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryFavoriteExistArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryFavoriteExistArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryFavoriteExistResult struct {
	Success *biz.QueryFavoriteExistResp
}

var QueryFavoriteExistResult_Success_DEFAULT *biz.QueryFavoriteExistResp

func (p *QueryFavoriteExistResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.QueryFavoriteExistResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryFavoriteExistResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryFavoriteExistResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryFavoriteExistResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryFavoriteExistResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryFavoriteExistResult) Unmarshal(in []byte) error {
	msg := new(biz.QueryFavoriteExistResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryFavoriteExistResult) GetSuccess() *biz.QueryFavoriteExistResp {
	if !p.IsSetSuccess() {
		return QueryFavoriteExistResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryFavoriteExistResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.QueryFavoriteExistResp)
}

func (p *QueryFavoriteExistResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryFavoriteExistResult) GetResult() interface{} {
	return p.Success
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.CommentActionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.InteractService).CommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentActionArgs:
		success, err := handler.(biz.InteractService).CommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newCommentActionArgs() interface{} {
	return &CommentActionArgs{}
}

func newCommentActionResult() interface{} {
	return &CommentActionResult{}
}

type CommentActionArgs struct {
	Req *biz.CommentActionReq
}

func (p *CommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.CommentActionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentActionArgs) Unmarshal(in []byte) error {
	msg := new(biz.CommentActionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentActionArgs_Req_DEFAULT *biz.CommentActionReq

func (p *CommentActionArgs) GetReq() *biz.CommentActionReq {
	if !p.IsSetReq() {
		return CommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentActionResult struct {
	Success *biz.CommentActionResp
}

var CommentActionResult_Success_DEFAULT *biz.CommentActionResp

func (p *CommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.CommentActionResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentActionResult) Unmarshal(in []byte) error {
	msg := new(biz.CommentActionResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentActionResult) GetSuccess() *biz.CommentActionResp {
	if !p.IsSetSuccess() {
		return CommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.CommentActionResp)
}

func (p *CommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentActionResult) GetResult() interface{} {
	return p.Success
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.CommentListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.InteractService).CommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentListArgs:
		success, err := handler.(biz.InteractService).CommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentListResult)
		realResult.Success = success
	}
	return nil
}
func newCommentListArgs() interface{} {
	return &CommentListArgs{}
}

func newCommentListResult() interface{} {
	return &CommentListResult{}
}

type CommentListArgs struct {
	Req *biz.CommentListReq
}

func (p *CommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.CommentListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentListArgs) Unmarshal(in []byte) error {
	msg := new(biz.CommentListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentListArgs_Req_DEFAULT *biz.CommentListReq

func (p *CommentListArgs) GetReq() *biz.CommentListReq {
	if !p.IsSetReq() {
		return CommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentListResult struct {
	Success *biz.CommentListResp
}

var CommentListResult_Success_DEFAULT *biz.CommentListResp

func (p *CommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.CommentListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentListResult) Unmarshal(in []byte) error {
	msg := new(biz.CommentListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentListResult) GetSuccess() *biz.CommentListResp {
	if !p.IsSetSuccess() {
		return CommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.CommentListResp)
}

func (p *CommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentListResult) GetResult() interface{} {
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

func (p *kClient) FavoriteAction(ctx context.Context, Req *biz.FavoriteActionReq) (r *biz.FavoriteActionResp, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFavoriteExist(ctx context.Context, Req *biz.QueryFavoriteExistReq) (r *biz.QueryFavoriteExistResp, err error) {
	var _args QueryFavoriteExistArgs
	_args.Req = Req
	var _result QueryFavoriteExistResult
	if err = p.c.Call(ctx, "QueryFavoriteExist", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, Req *biz.CommentActionReq) (r *biz.CommentActionResp, err error) {
	var _args CommentActionArgs
	_args.Req = Req
	var _result CommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, Req *biz.CommentListReq) (r *biz.CommentListResp, err error) {
	var _args CommentListArgs
	_args.Req = Req
	var _result CommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

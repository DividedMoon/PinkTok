// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	"user_service/biz"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*biz.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":   kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"Login":      kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"UserInfo":   kitex.NewMethodInfo(userInfoHandler, newUserInfoArgs, newUserInfoResult, false),
		"UserUpdate": kitex.NewMethodInfo(userUpdateHandler, newUserUpdateArgs, newUserUpdateResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user_service",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.UserRegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(biz.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *biz.UserRegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.UserRegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(biz.UserRegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *biz.UserRegisterReq

func (p *RegisterArgs) GetReq() *biz.UserRegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *biz.UserRegisterResp
}

var RegisterResult_Success_DEFAULT *biz.UserRegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.UserRegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(biz.UserRegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *biz.UserRegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.UserRegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.UserLoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(biz.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *biz.UserLoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.UserLoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(biz.UserLoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *biz.UserLoginReq

func (p *LoginArgs) GetReq() *biz.UserLoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *biz.UserLoginResp
}

var LoginResult_Success_DEFAULT *biz.UserLoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.UserLoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(biz.UserLoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *biz.UserLoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.UserLoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.UserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.UserService).UserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserInfoArgs:
		success, err := handler.(biz.UserService).UserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUserInfoArgs() interface{} {
	return &UserInfoArgs{}
}

func newUserInfoResult() interface{} {
	return &UserInfoResult{}
}

type UserInfoArgs struct {
	Req *biz.UserInfoReq
}

func (p *UserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.UserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserInfoArgs) Unmarshal(in []byte) error {
	msg := new(biz.UserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserInfoArgs_Req_DEFAULT *biz.UserInfoReq

func (p *UserInfoArgs) GetReq() *biz.UserInfoReq {
	if !p.IsSetReq() {
		return UserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserInfoResult struct {
	Success *biz.UserInfoResp
}

var UserInfoResult_Success_DEFAULT *biz.UserInfoResp

func (p *UserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.UserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserInfoResult) Unmarshal(in []byte) error {
	msg := new(biz.UserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserInfoResult) GetSuccess() *biz.UserInfoResp {
	if !p.IsSetSuccess() {
		return UserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.UserInfoResp)
}

func (p *UserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserInfoResult) GetResult() interface{} {
	return p.Success
}

func userUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(biz.UserUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(biz.UserService).UserUpdate(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserUpdateArgs:
		success, err := handler.(biz.UserService).UserUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserUpdateResult)
		realResult.Success = success
	}
	return nil
}
func newUserUpdateArgs() interface{} {
	return &UserUpdateArgs{}
}

func newUserUpdateResult() interface{} {
	return &UserUpdateResult{}
}

type UserUpdateArgs struct {
	Req *biz.UserUpdateReq
}

func (p *UserUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(biz.UserUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserUpdateArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserUpdateArgs) Unmarshal(in []byte) error {
	msg := new(biz.UserUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserUpdateArgs_Req_DEFAULT *biz.UserUpdateReq

func (p *UserUpdateArgs) GetReq() *biz.UserUpdateReq {
	if !p.IsSetReq() {
		return UserUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserUpdateResult struct {
	Success *biz.UserUpdateResp
}

var UserUpdateResult_Success_DEFAULT *biz.UserUpdateResp

func (p *UserUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(biz.UserUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserUpdateResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserUpdateResult) Unmarshal(in []byte) error {
	msg := new(biz.UserUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserUpdateResult) GetSuccess() *biz.UserUpdateResp {
	if !p.IsSetSuccess() {
		return UserUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*biz.UserUpdateResp)
}

func (p *UserUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserUpdateResult) GetResult() interface{} {
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

func (p *kClient) Register(ctx context.Context, Req *biz.UserRegisterReq) (r *biz.UserRegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *biz.UserLoginReq) (r *biz.UserLoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, Req *biz.UserInfoReq) (r *biz.UserInfoResp, err error) {
	var _args UserInfoArgs
	_args.Req = Req
	var _result UserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserUpdate(ctx context.Context, Req *biz.UserUpdateReq) (r *biz.UserUpdateResp, err error) {
	var _args UserUpdateArgs
	_args.Req = Req
	var _result UserUpdateResult
	if err = p.c.Call(ctx, "UserUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

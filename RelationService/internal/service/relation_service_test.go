package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/stretchr/testify/mock"
	"relation_service/biz"
	internalClient "relation_service/internal/client"
	"relation_service/internal/model"
	"relation_service/internal/mw/redis"
	"testing"
	userBiz "user_service/biz"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, Req *userBiz.UserRegisterReq, callOptions ...callopt.Option) (r *userBiz.UserRegisterResp, err error) {
	args := m.Called(ctx, Req)
	return args.Get(0).(*userBiz.UserRegisterResp), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, Req *userBiz.UserLoginReq, callOptions ...callopt.Option) (r *userBiz.UserLoginResp, err error) {
	args := m.Called(ctx, Req)
	return args.Get(0).(*userBiz.UserLoginResp), args.Error(1)
}

func (m *MockUserService) UserInfo(ctx context.Context, Req *userBiz.UserInfoReq, callOptions ...callopt.Option) (r *userBiz.UserInfoResp, err error) {
	args := m.Called(ctx, Req)
	return args.Get(0).(*userBiz.UserInfoResp), args.Error(1)
}

func (m *MockUserService) UserUpdate(ctx context.Context, Req *userBiz.UserUpdateReq, callOptions ...callopt.Option) (r *userBiz.UserUpdateResp, err error) {
	args := m.Called(ctx, Req)
	return args.Get(0).(*userBiz.UserUpdateResp), args.Error(1)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func setup() {
	redis.InitRedis()
	model.InitDB()
}

func teardown() {
	redis.CloseRedis()
}

func TestSubmitFollowRelationAction(t *testing.T) {
	var (
		ctx = context.Background()
		req = &biz.RelationActionReq{
			UserId:     1,
			ToUserId:   1001,
			ActionType: 1,
		}
	)

	err := SubmitFollowRelationAction(ctx, req)
	assert.Assert(t, err == nil)
}

func TestGetFollowListById(t *testing.T) {
	mock := new(MockUserService)
	var (
		req1  = &userBiz.UserInfoReq{UserId: 2}
		req2  = &userBiz.UserInfoReq{UserId: 1001}
		resp1 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   2,
				Name: "test1",
			},
		}
		reps2 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   1001,
				Name: "test2",
			},
		}
	)
	mock.On("UserInfo", context.TODO(), req1).Return(resp1, nil)
	mock.On("UserInfo", context.TODO(), req2).Return(reps2, nil)
	internalClient.UserServiceClient = mock
	infos, err := GetFollowListById(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, len(infos) == 2)
	assert.Assert(t, infos[0].Name == "test1")
	assert.Assert(t, infos[1].Name == "test2")
}

func TestGetFollowerListById(t *testing.T) {
	f1 := &model.Follower{
		UserIdA: 1,
		UserIdB: 2,
	}
	f2 := &model.Follower{
		UserIdA: 1,
		UserIdB: 1001,
	}
	_ = f1.Create(model.DB)
	_ = f2.Create(model.DB)
	defer func() {
		f1.HardDelete()
		f2.HardDelete()
	}()
	mock := new(MockUserService)
	var (
		req1  = &userBiz.UserInfoReq{UserId: 2}
		req2  = &userBiz.UserInfoReq{UserId: 1001}
		resp1 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   2,
				Name: "test1",
			},
		}
		reps2 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   1001,
				Name: "test2",
			},
		}
	)
	mock.On("UserInfo", context.TODO(), req1).Return(resp1, nil)
	mock.On("UserInfo", context.TODO(), req2).Return(reps2, nil)
	internalClient.UserServiceClient = mock
	infos, err := GetFollowerListById(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, len(infos) == 2)
	assert.Assert(t, infos[0].Name == "test1")
	assert.Assert(t, infos[1].Name == "test2")
}

func TestGetFriendListById(t *testing.T) {
	dbData := []interface{}{
		&model.Follow{
			UserIdA: 1,
			UserIdB: 2,
		},
		&model.Follow{
			UserIdA: 1,
			UserIdB: 1001,
		},
		&model.Follower{
			UserIdA: 2,
			UserIdB: 1,
		},
		&model.Follower{
			UserIdA: 1001,
			UserIdB: 1,
		},
		&model.Follow{
			UserIdA: 1001,
			UserIdB: 1,
		},
		&model.Follower{
			UserIdA: 1,
			UserIdB: 1001,
		},
	}
	for _, do := range dbData {
		if follow, ok := do.(*model.Follow); ok {
			_ = follow.Create(model.DB)
		} else if follower, ok := do.(*model.Follower); ok {
			_ = follower.Create(model.DB)
		}
	}
	defer func() {
		for _, do := range dbData {
			if follow, ok := do.(*model.Follow); ok {
				follow.HardDelete()
			} else if follower, ok := do.(*model.Follower); ok {
				follower.HardDelete()
			}
		}
	}()
	mock := new(MockUserService)
	var (
		req2  = &userBiz.UserInfoReq{UserId: 1001}
		reps2 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   1001,
				Name: "test2",
			},
		}
	)
	mock.On("UserInfo", context.TODO(), req2).Return(reps2, nil)
	internalClient.UserServiceClient = mock
	infos, err := GetFriendListById(context.Background(), 1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Assert(t, len(infos) == 1)
	assert.Assert(t, infos[0].Name == "test2")
}

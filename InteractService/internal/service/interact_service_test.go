package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/stretchr/testify/mock"
	internalClient "interact_service/internal/client"
	"interact_service/internal/model"
	"testing"
	userBiz "user_service/biz"
)

var (
	mockUser *MockUserService
)

func TestMain(m *testing.M) {
	model.InitDB()
	mockUser = new(MockUserService)
	m.Run()
}

func TestPublishComment(t *testing.T) {
	var (
		req1  = &userBiz.UserInfoReq{UserId: 1001}
		resp1 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   1001,
				Name: "test1",
			},
		}
	)
	mockUser.On("UserInfo", context.TODO(), req1).Return(resp1, nil)
	internalClient.UserServiceClient = mockUser
	resp, err := PublishComment(context.TODO(), &model.Comment{
		UserId:  1001,
		VideoId: 10,
		Content: "好看的2",
	})
	assert.Nil(t, err)
	assert.Assert(t, req1.UserId == resp.User.Id)
	assert.NotEqual(t, 0, resp.Id)
}

func TestPublishCommentWithSame(t *testing.T) {
	var (
		req1  = &userBiz.UserInfoReq{UserId: 1001}
		resp1 = &userBiz.UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: &userBiz.UserInfo{
				Id:   1001,
				Name: "test1",
			},
		}
	)
	mockUser.On("UserInfo", context.TODO(), req1).Return(resp1, nil)
	internalClient.UserServiceClient = mockUser
	resp, err := PublishComment(context.TODO(), &model.Comment{
		UserId:  1001,
		VideoId: 10,
		Content: "Test PublishCommentWithSame22",
	})
	assert.Nil(t, err)
	assert.Assert(t, req1.UserId == resp.User.Id)
	assert.NotEqual(t, 0, resp.Id)
	resp, err = PublishComment(context.TODO(), &model.Comment{
		UserId:  1001,
		VideoId: 10,
		Content: "Test PublishCommentWithSame22",
	})
	assert.Nil(t, err)
	assert.Assert(t, req1.UserId == resp.User.Id)
	assert.NotEqual(t, 0, resp.Id)
}

func TestDeleteComment(t *testing.T) {
	c := &model.Comment{
		ID:      10,
		UserId:  1001,
		VideoId: 10,
	}
	err := DeleteComment(c)
	assert.Nil(t, err)
	model.DB.Model(&model.Comment{}).Where("id = ?", c.ID).First(c)
	assert.True(t, c.Deleted == 1)
}

func TestGetCommentByUserAndVideo(t *testing.T) {
	video, err := GetCommentByUserAndVideo(context.Background(), 1001, 10)
	assert.Nil(t, err)
	assert.Assert(t, len(video) == 3)
	hlog.Info(video)
}

// TestQueryFavoriteExist 查询喜欢是否存在 Done
func TestQueryFavoriteExist(t *testing.T) {
	exist, err := QueryFavoriteExist(1, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(exist)

}

// TestQueryUserFavoriteVideoIds 测试查询用户收藏视频id列表 Done
func TestQueryUserFavoriteVideoIds(t *testing.T) {
	ids, err := QueryUserFavoriteVideoIds(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(ids)

}

// TestAddFavoriteRecord 测试点赞记录 Done
func TestAddFavoriteRecord(t *testing.T) {
	//_ = AddFavoriteRecord(1, 1, 1) //对已存在记录点赞
	//_ = AddFavoriteRecord(1, 1, 2) //对已存在记录取消赞
	_ = AddFavoriteRecord(8, 11, 1) //对未存在记录点赞
	//_ = AddFavoriteRecord(8, 11, 2) //对未存在记录取消赞
}

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

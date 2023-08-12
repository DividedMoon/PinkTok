package client

import "user_service/biz/model/client/user_service"

var (
	UserServiceClient, _ = user_service.NewUserServiceClient("//127.0.0.1:8889")
)

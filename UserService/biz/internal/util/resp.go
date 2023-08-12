package utils

import (
	"errors"
	"user_service/biz/internal/constant"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

// BuildBaseResp convert error and build BaseResp
func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(constant.Success)
	}

	e := constant.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := constant.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

// baseResp build BaseResp from error
func baseResp(err constant.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

package utils

import (
	"errors"
	"video_service/internal/constants"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

// BuildBaseResp convert error and build BaseResp
func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(constants.Success)
	}

	e := constants.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := constants.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

// baseResp build BaseResp from error
func baseResp(err constants.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

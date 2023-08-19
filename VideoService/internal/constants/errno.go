package constants

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 20000
	InnerServiceErrCode
	DBErrCode
	VideoCopyErrCode
	GetExtensionErrCode
	UploadImgErrCode
)

const (
	SuccessMsg         = "Success"
	ServiceErrMsg      = "Service call failed"
	DBErrMsg           = "DB call failed"
	VideoCopyErrMsg    = "Video copy failed"
	GetExtensionErrMsg = "Get extension failed"
	UploadImgErrMsg    = "Upload image failed"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr = NewErrNo(ServiceErrCode, ServiceErrMsg)
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

package constant

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 30000
	ParamErrCode
	AuthorizationFailedErrCode
	InvalidActionTypeCode
	UserNotExistErrCode
	UpdateNotEqualOneErrCode
)

const (
	SuccessMsg           = "Success"
	ServerErrMsg         = "Service is unable to start successfully"
	ParamErrMsg          = "Wrong Parameter has been given"
	InvalidActionTypeMsg = "action type is invalid"
	UserNotExistErrMsg   = "user not exist"
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
	Success                = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr             = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr               = NewErrNo(ParamErrCode, ParamErrMsg)
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	InvalidActionTypeErr   = NewErrNo(InvalidActionTypeCode, InvalidActionTypeMsg)
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, UserNotExistErrMsg)
	UpdateNotEqualOneErr   = NewErrNo(UpdateNotEqualOneErrCode, "update not equal one")
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

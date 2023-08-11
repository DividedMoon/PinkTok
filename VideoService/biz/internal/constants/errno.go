package constants

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 20000
	InnerServiceErrCode
)

const (
	SuccessMsg    = "Success"
	ServiceErrMsg = "Service call failed"
)

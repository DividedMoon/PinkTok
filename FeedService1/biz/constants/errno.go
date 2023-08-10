package constants

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10000

	ReturnErrCode
)

const (
	SuccessMsg    = "Success"
	ServiceErrMsg = "Service call failed"
	ReturnErrMsg  = "Return"
)

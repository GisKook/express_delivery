package base

import ()

type ExpressDeliveryError struct {
	Err      error
	Code     int
	Describe string
}

func (e *ExpressDeliveryError) Error() string {
	return e.Err.Error()
}

func (e *ExpressDeliveryError) Desc() string {
	return e.Describe
}

func NewErr(err error, code int, desc string) *ExpressDeliveryError {
	return &ExpressDeliveryError{
		Err:      err,
		Code:     code,
		Describe: desc,
	}
}

const (
	ERR_COMMON_NOT_CAPTURE_CODE int    = 999
	ERR_COMMON_NOT_CAPTURE_DESC string = "未捕获的错误"

	ERR_NONE_CODE int    = 0
	ERR_NONE_DESC string = "成功"

	ERR_HTTP_LACK_PARAMTERS_CODE int    = 1
	ERR_HTTP_LACK_PARAMTERS_DESC string = "缺少参数"

	ERR_HTTP_INNER_PANIC_CODE int    = 2
	ERR_HTTP_INNER_PANIC_DESC string = "内部错误"

	ERR_HTTP_TIMEOUT_CODE int    = 3
	ERR_HTTP_TIMEOUT_DESC string = "超时"

	ERR_HTTP_WECHAT_LOGIN_CODE int    = 4
	ERR_HTTP_WECHAT_LOGIN_DESC string = "微信登录失败"

	ERR_HTTP_WECHAT_GEN_SESSION_CODE int    = 5
	ERR_HTTP_WECHAT_GEN_SESSION_DESC string = "登录微信生成session失败"
)

var (
	ERROR_HTTP_LACK_PARAMTERS *ExpressDeliveryError = NewErr(nil, ERR_HTTP_LACK_PARAMTERS_CODE, ERR_HTTP_LACK_PARAMTERS_DESC)
	ERROR_HTTP_INNER_PANIC    *ExpressDeliveryError = NewErr(nil, ERR_HTTP_INNER_PANIC_CODE, ERR_HTTP_INNER_PANIC_DESC)
	ERROR_HTTP_TIMEOUT        *ExpressDeliveryError = NewErr(nil, ERR_HTTP_TIMEOUT_CODE, ERR_HTTP_TIMEOUT_DESC)
	ERROR_NONE                *ExpressDeliveryError = NewErr(nil, ERR_NONE_CODE, ERR_NONE_DESC)
	ERROR_NOT_CAPTURE         *ExpressDeliveryError = NewErr(nil, ERR_COMMON_NOT_CAPTURE_CODE, ERR_COMMON_NOT_CAPTURE_DESC)
)

package resp

import "net/http"

var (
	// errcode枚举
	Code_Succ = 0
	Code_Fail = 1
	Code_Err  = 999

	Code_Cal = 31000

	Code_ParamInValid = 10000
	Code_NoLogin      = 10004
	Code_LoginTimeout = 10005
	Code_IllegalToken = 10006
	Code_Repeat       = 10010
	Code_Illegal      = 20003

	Code_NotFund = 40001
)

var (
	// 参数错误
	ParamInValid = func(msg ...string) *Exception {
		return NewException(http.StatusBadRequest, Code_ParamInValid, msg...)
	}
	Repeat = NewException(http.StatusBadRequest, Code_Repeat, "资源重复")

	NoLogin      = NewException(http.StatusUnauthorized, Code_NoLogin, "未登录")
	LoginTimeOut = NewException(http.StatusUnauthorized, Code_LoginTimeout, "登录超时")
	IllegalToken = NewException(http.StatusUnauthorized, Code_IllegalToken, "token非法")
	Illegal      = NewException(http.StatusForbidden, Code_Illegal, "非法操作")

	NotFund = NewException(http.StatusNotFound, Code_NotFund, "资源未找到")
)

package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code = 1000xxx 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONT      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code = 2000xxx 文章模块错误
	ERROR_CATENAME_USED = 2001
	// code = 3000xxx 分类模块错误
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已经存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONT:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_CATENAME_USED:    "分类已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
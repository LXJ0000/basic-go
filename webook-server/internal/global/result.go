package g

import "fmt"

const (
	SUCCESS = 0
	FAIL    = 1
)

type Result struct {
	code int
	msg  string
}

func (e Result) Code() int {
	return e.code
}

func (e Result) Msg() string {
	return e.msg
}

var (
	_codes    = map[int]struct{}{}
	_messages = make(map[int]string)
)

func RegisterResult(code int, msg string) Result {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("Error code %d already exists, please change it.", code))
	}
	if msg == "" {
		panic("Error code message cannot be empty")
	}
	_codes[code] = struct{}{}
	_messages[code] = msg
	return Result{
		code: code,
		msg:  msg,
	}
}

func GetMsg(code int) string {
	return _messages[code]
}

var (
	SuccessResult = RegisterResult(SUCCESS, "ok")
	FailResult    = RegisterResult(FAIL, "fail")
)

var (
	ErrRequest  = RegisterResult(1001, "请求参数有误")
	ErrDbOp     = RegisterResult(1002, "MySQL 数据库操作异常")
	ErrRedisOp  = RegisterResult(1003, "Redis 操作异常")
	ErrUserAuth = RegisterResult(1004, "用户认证异常")

	ErrPassword              = RegisterResult(2001, "密码有误")
	ErrUserNotExist          = RegisterResult(2002, "该用户不存在")
	ErrOldPassword           = RegisterResult(2003, "旧密码不正确")
	ErrPasswordsInconsistent = RegisterResult(2004, "密码不一致")
	ErrEmailFormatWrong      = RegisterResult(2005, "非法邮箱格式")
	ErrPasswordStrength      = RegisterResult(2006, "密码必须包含字母、数字、特殊字符，并且不少于八位")
	ErrBcryptFail            = RegisterResult(2007, "加密失败")
	ErrUserExist             = RegisterResult(2008, "邮箱或用户名已存在")

	ErrCodeSendFrequently   = RegisterResult(3001, "验证码发送过于频繁")
	ErrCodeVerifyFrequently = RegisterResult(3002, "验证过于频繁")
	ErrCodeWrong            = RegisterResult(3003, "验证码有误")

	ErrTokenNotExist = RegisterResult(4001, "TOKEN 不存在，请重新登陆")
	ErrTokenRuntime  = RegisterResult(4002, "TOKEN 已过期，请重新登陆")
	ErrTokenWrong    = RegisterResult(4003, "TOKEN 不正确，请重新登陆")
	ErrTokenType     = RegisterResult(4004, "TOKEN 格式错误，请重新登陆")
	ErrTokenCreate   = RegisterResult(4005, "TOKEN 生成失败")
)

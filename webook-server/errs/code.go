package errs

//40 客户端错误 50 服务端错误

// User
const (
	CodeUserInvalidInput = 40100 + iota
	CodeUserNameOrEmailDuplicate
	CodeUserNotAuthorization

	CodeUserInternalServerError = 50100
)

package web

import (
	"errors"
	"net/http"
	"webook-server/errs"
	"webook-server/internal/domain"
	"webook-server/internal/service"
	"webook-server/pkg/jwt"
	"webook-server/pkg/snowflake"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const (
	RegexpPassword = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	RegexpEmail    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
)

type UserHandler struct {
	svc            *service.UserService
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"'msg'"`
	Data interface{} `json:"data,omitempty"`
}

type UserToken struct {
	Token string `json:"token"`
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:            svc,
		emailRegexp:    regexp.MustCompile(RegexpEmail, regexp.None),
		passwordRegexp: regexp.MustCompile(RegexpPassword, regexp.None),
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	type Req struct {
		UserName        string `json:"user_name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "请求参数有误",
		})
		return
	}

	if isEmail, _ := h.emailRegexp.MatchString(req.Email); !isEmail {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "非法邮箱格式",
		})
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "密码不一致",
		})
		return
	}
	if isPassword, _ := h.passwordRegexp.MatchString(req.Password); !isPassword {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "密码必须包含字母、数字、特殊字符，并且不少于八位",
		})
		return
	}

	err := h.svc.Register(ctx, domain.User{
		UserId:   snowflake.GenID(),
		Email:    req.Email,
		UserName: req.UserName,
		Password: req.Password,
	})
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, Response{
			Code: 0,
			Msg:  "注册成功",
		})
	case errors.Is(err, service.ErrDuplicateEmail):
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserNameOrEmailDuplicate,
			Msg:  "邮箱或用户名已存在",
		})
	default:
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "请求参数有误",
		})
		return
	}

	user, err := h.svc.Login(ctx, req.Email, req.Password)

	switch {
	case err == nil:
		token, _ := jwt.GenToken(ctx, user.UserId, user.UserName)
		ctx.JSON(http.StatusOK, Response{
			Code: 0,
			Msg:  "登录成功",
			Data: UserToken{token},
		})
	case errors.Is(err, service.ErrInvalidUserOrPassword):
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "用户名或密码错误",
		})
	default:
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
	}
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	userIdRaw, exist := ctx.Get("user_id")
	userId, ok := userIdRaw.(int64)
	if !exist || !ok {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserNotAuthorization,
			Msg:  "用户登录状态有误",
		})
		return
	}
	user, err := h.svc.Profile(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: 1,
			Msg:  "fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: user,
	})
}

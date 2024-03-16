package web

import (
	"errors"
	"net/http"
	"webook-server/errs"
	"webook-server/internal/domain"
	"webook-server/internal/service"
	"webook-server/internal/web/middleware"
	"webook-server/pkg/jwt"
	"webook-server/pkg/snowflake"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/login", h.Login)
	userGroup.POST("/register", h.Register)

	userGroup.POST("login/sms/code", h.SendLoginSMSCode)
	userGroup.POST("login/sms/verify", h.VerifyLoginSMSCode)

	authUserGroup := userGroup.Use(middleware.JwtAuthMiddleware())
	authUserGroup.GET("/", h.Profile)
}

const (
	biz = "login"

	RegexpPassword = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	RegexpEmail    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
)

type UserHandler struct {
	svc            *service.UserService
	codeSvc        *service.CodeService
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService, codeSvc *service.CodeService) *UserHandler {
	return &UserHandler{
		svc:            svc,
		codeSvc:        codeSvc,
		emailRegexp:    regexp.MustCompile(RegexpEmail, regexp.None),
		passwordRegexp: regexp.MustCompile(RegexpPassword, regexp.None),
	}
}

func (h *UserHandler) VerifyLoginSMSCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "请求参数有误",
		})
		return
	}
	ok, err := h.codeSvc.Verify(ctx, biz, req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "验证码有误",
		})
		return
	}

	user, err := h.svc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
		return
	}

	token, _ := jwt.GenToken(ctx, user.UserId, "")
	ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "验证通过",
		Data: gin.H{
			"token": token,
		},
	})
}

func (h *UserHandler) SendLoginSMSCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInvalidInput,
			Msg:  "请求参数有误",
		})
		return
	}
	err := h.codeSvc.Send(ctx, biz, req.Phone)
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, Response{
			Code: 0,
			Msg:  "发送成功",
		})
	case errors.Is(err, service.ErrCodeSendFrequently):
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError, // todo
			Msg:  "发送太频繁，请稍后尝试",
		})
	default:
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
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
	case errors.Is(err, service.ErrDuplicate):
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
			Data: gin.H{
				"token": token,
			},
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
	if err != nil { // todo 正常前端是不会错的，不处理了
		ctx.JSON(http.StatusOK, Response{
			Code: errs.CodeUserInternalServerError,
			Msg:  "系统错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: user,
	})
}

package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
	"webook-server/internal/domain"
	"webook-server/internal/service"
	"webook-server/pkg/jwt"
	"webook-server/pkg/snowflake"
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
		// todo error
		return
	}
	if isEmail, _ := h.emailRegexp.MatchString(req.Email); !isEmail {
		ctx.String(http.StatusOK, "非法邮箱格式")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入密码不对")
		return
	}
	if isPassword, _ := h.passwordRegexp.MatchString(req.Password); !isPassword {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符，并且不少于八位")
		return
	}

	err := h.svc.Register(ctx, domain.User{
		UserId:   snowflake.GenID(),
		Email:    req.Email,
		UserName: req.UserName,
		Password: req.Password,
	})
	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱或用户名已存在")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	user, err := h.svc.Login(ctx, req.Email, req.Password)

	switch err {
	case nil:
		token, _ := jwt.GenToken(ctx, user.UserId, user.UserName)
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
			"code":  0,
			"msg":   "登陆成功",
		})
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码错误")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	user, err := h.svc.Profile(ctx, userId.(int64)) // todo .() error
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

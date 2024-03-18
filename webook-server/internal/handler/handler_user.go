package handler

import (
	"errors"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"webook-server/internal/domain"
	"webook-server/internal/global"
	"webook-server/internal/middleware"
	"webook-server/internal/repository"
	"webook-server/internal/utils/jwt"
	"webook-server/internal/utils/snowflake"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) InitRouter(r *gin.Engine) {
	base := r.Group("/api")

	base.POST("/login", h.Login)
	base.POST("/register", h.Register)

	auth := base.Group("/user").Use(middleware.JwtAuthMiddleware())
	auth.GET("/info", h.Info)

}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnFail(c, g.ErrRequest, err.Error())
		return
	}
	u, err := h.repo.FindByEmail(c, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ReturnFail(c, g.ErrUserNotExist, err.Error())
			return
		}
		ReturnFail(c, g.ErrDbOp, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		ReturnFail(c, g.ErrPassword, err.Error())
		return
	}
	token, err := jwt.GenToken(c, u.UserId, u.UserName)
	if err != nil {
		ReturnFail(c, g.ErrTokenCreate, err.Error())
		return
	}
	ReturnSuccess(c, LoginResp{
		u, token,
	})
}
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnFail(c, g.ErrRequest, err.Error())
		return
	}

	const (
		RegexpPassword = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
		RegexpEmail    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	)
	var (
		emailRegexp    = regexp.MustCompile(RegexpEmail, regexp.None)
		passwordRegexp = regexp.MustCompile(RegexpPassword, regexp.None)
	)
	if isEmail, _ := emailRegexp.MatchString(req.Email); !isEmail {
		ReturnFail(c, g.ErrEmailFormatWrong, "")
		return
	}
	if req.Password != req.ConfirmPassword {
		ReturnFail(c, g.ErrPasswordsInconsistent, "")
		return
	}
	if isPassword, _ := passwordRegexp.MatchString(req.Password); !isPassword {
		// todo 密码强度
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ReturnFail(c, g.ErrBcryptFail, err.Error())
		return
	}
	u := domain.User{
		UserId:   snowflake.GenID(),
		Email:    req.Email,
		Password: string(encrypted),
	}
	if err := h.repo.Create(c, u); err != nil {
		//todo 判断数据库异常 还是 用户已存在 Create 异常
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ReturnFail(c, g.ErrUserExist, err.Error())
			return
		}
		ReturnFail(c, g.ErrDbOp, err.Error())
		return
	}
	ReturnSuccess(c, nil)
}

func (h *UserHandler) Info(c *gin.Context) {
	userIdRaw, exist := c.Get("user_id")
	userId, ok := userIdRaw.(int64)
	if !exist || !ok {
		ReturnFail(c, g.ErrUserAuth, "")
		return
	}
	user, err := h.repo.FindByUserId(c, userId)
	if err != nil { // todo 正常前端是不会错的，不处理了
		ReturnFail(c, g.ErrDbOp, err.Error())
		return
	}

	ReturnSuccess(c, user)
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterReq struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
type LoginResp struct {
	domain.User
	Token string `json:"token"`
}

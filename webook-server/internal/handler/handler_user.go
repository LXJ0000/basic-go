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
	"webook-server/internal/utils/sms"
	"webook-server/internal/utils/snowflake"
)

const (
	biz            = "login"
	codeTemplateId = "xxx"
)

type UserHandler struct {
	userRepo repository.UserRepository
	codeRepo repository.CodeRepository
	sms      sms.Service
}

func NewUserHandler(userRepo repository.UserRepository, codeRepo repository.CodeRepository, sms sms.Service) *UserHandler {
	return &UserHandler{userRepo: userRepo, codeRepo: codeRepo, sms: sms}
}

func (h *UserHandler) InitRouter(r *gin.Engine) {
	base := r.Group("/api")

	base.POST("/login", h.Login)
	base.POST("/register", h.Register)
	base.POST("login/sms/sent", h.SendLoginSMSCode)
	base.POST("login/sms/verify", h.VerifyLoginSMSCode)

	auth := base.Group("/user").Use(middleware.JwtAuthMiddleware())
	auth.GET("/info", h.Info)

}

func (h *UserHandler) Login(c *gin.Context) {
	type Req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	type Resp struct {
		domain.User
		Token string `json:"token"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnFail(c, g.ErrRequest, err.Error())
		return
	}
	u, err := h.userRepo.FindByEmail(c, req.Email)
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
	ReturnSuccess(c, Resp{
		u, token,
	})
}
func (h *UserHandler) Register(c *gin.Context) {
	type RegisterReq struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}
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
	if err := h.userRepo.Create(c, u); err != nil {
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
	type Resp struct {
		domain.User
		//	todo 看前端需求
	}
	userIdRaw, exist := c.Get("user_id")
	userId, ok := userIdRaw.(int64)
	if !exist || !ok {
		ReturnFail(c, g.ErrUserAuth, "")
		return
	}
	user, err := h.userRepo.FindByUserId(c, userId)
	if err != nil { // todo 正常前端是不会错的，不处理了
		ReturnFail(c, g.ErrDbOp, err.Error())
		return
	}

	ReturnSuccess(c, Resp{
		user,
	})
}
func (h *UserHandler) VerifyLoginSMSCode(c *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	type Resp struct {
		domain.User
		Token string `json:"token"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnFail(c, g.ErrRequest, err.Error())
		return
	}

	ok, err := h.codeRepo.Verify(c, biz, req.Phone, req.Code)
	if err != nil { // 不可以重试
		if errors.Is(err, repository.ErrCodeVerifyFrequently) {
			ReturnFail(c, g.ErrCodeVerifyFrequently, err.Error()) // 超次数
			return
		}
		ReturnFail(c, g.ErrRedisOp, err.Error()) // 系统错误
		return
	}
	if !ok { // 重试
		ReturnFail(c, g.ErrCodeWrong, "验证码有误，请重新下输入")
		return
	}

	//FindOrCreate user
	u, err := h.userRepo.FindByPhone(c, req.Phone)
	if err != nil { // 用户不存在
		u = domain.User{
			UserId: snowflake.GenID(),
			Phone:  req.Phone,
		}
		if err = h.userRepo.Create(c, u); err != nil {
			ReturnFail(c, g.ErrDbOp, err.Error())
			return
		}
	}
	token, err := jwt.GenToken(c, u.UserId, u.UserName)
	if err != nil {
		ReturnFail(c, g.ErrTokenCreate, err.Error())
		return
	}
	ReturnSuccess(c, Resp{
		u, token,
	})
}
func (h *UserHandler) SendLoginSMSCode(c *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnFail(c, g.ErrRequest, err.Error())
		return
	}

	//发送验证码
	//1. 生成code
	//2. store
	//3. send
	code := sms.GenerateCode()
	if err := h.codeRepo.Store(c, biz, req.Phone, code); err != nil {
		if errors.Is(err, repository.ErrCodeSendFrequently) {
			ReturnFail(c, g.ErrCodeSendFrequently, err.Error())
		}
		ReturnFail(c, g.ErrRedisOp, err.Error())
		return
	}
	if err := h.sms.Send(c, codeTemplateId, []string{code}, req.Phone); err != nil {
		ReturnFail(c, g.FailResult, err.Error())
		return
	}
	ReturnSuccess(c, nil)
}

package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"webook-server/internal/domain"
	"webook-server/internal/global"
	"webook-server/internal/middleware"
	"webook-server/internal/repository"
	"webook-server/internal/utils/jwt"
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
type LoginResp struct {
	domain.User
	Token string `json:"token"`
}

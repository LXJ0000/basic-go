package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

const (
// todo regex
)

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (u *UserHandler) Profile(ctx *gin.Context) {

}

func (u *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// todo error
		return
	}
	//todo 登陆逻辑
	ctx.JSON(http.StatusOK, Resp{
		Code: 0,
		Msg:  "登录成功",
	})
}
func (u *UserHandler) Register(ctx *gin.Context) {
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
	//todo 参数校验 - 正则匹配
	//todo 注册逻辑
	ctx.JSON(http.StatusOK, Resp{
		Code: 0,
		Msg:  "注册成功",
	})
}

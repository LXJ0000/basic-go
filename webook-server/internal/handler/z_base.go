package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"webook-server/internal/global"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ReturnSuccess(c *gin.Context, data interface{}) {
	r := g.SuccessResult
	c.JSON(http.StatusOK, Response{
		Code:    r.Code(),
		Message: r.Msg(),
		Data:    data,
	})
}

func ReturnFail(c *gin.Context, r g.Result, data string) {
	slog.Info("[Func-ReturnError] " + r.Msg())
	slog.Error(data)
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    r.Code(),
		Message: r.Msg(),
		Data:    nil,
	})
}

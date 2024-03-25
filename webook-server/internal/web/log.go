package web

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func WrapReq[T any](fn func(ctx *gin.Context, req T) (Response, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		if err := ctx.ShouldBindJSON(&req); err != nil {
			slog.Error("请求参数有误", err.Error())
			return
		}
		res, err := fn(ctx, req)
		if err != nil {
			slog.Error("", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

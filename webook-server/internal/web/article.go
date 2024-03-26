package web

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"webook-server/internal/domain"
	"webook-server/internal/service"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func NewArticleHandler(svc service.ArticleService) ArticleHandler {
	return ArticleHandler{svc: svc}
}

func (h *ArticleHandler) InitRouter(r *gin.Engine) {
	base := r.Group("/api/article")
	base.POST("", h.CreateOrUpdate)
	base.DELETE("", h.Delete)
}

func (h *ArticleHandler) CreateOrUpdate(c *gin.Context) {

}

func (h *ArticleHandler) Delete(c *gin.Context) {
	type Req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("请求参数有误", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	authorId, ok := c.Get("user_id")
	if !ok {
		slog.Error("author fail", "")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := h.svc.CreateOrUpdate(c, domain.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: authorId.(int64),
	})
	if err != nil {
		slog.Error("CreateOrUpdate fail", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Data: 1,
	})
}

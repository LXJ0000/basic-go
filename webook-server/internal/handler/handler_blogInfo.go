package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	g "webook-server/internal/global"
	"webook-server/internal/middleware"
	"webook-server/internal/repository"
)

type BlogInfoHandler struct {
	repo repository.Repository
}

func NewBlogInfoHandler(repo repository.Repository) BlogInfoHandler {
	return BlogInfoHandler{repo: repo}
}

func (h *BlogInfoHandler) InitRouter(r *gin.Engine) {
	base := r.Group("/api")

	auth := base.Use(middleware.JwtAuthMiddleware())
	auth.GET("/home", h.HomeInfo)
}

func (h *BlogInfoHandler) HomeInfo(c *gin.Context) {
	type Resp struct {
		ArticleCount int `json:"article_count"` // 文章数量
		UserCount    int `json:"user_count"`    // 用户数量
		MessageCount int `json:"message_count"` // 留言数量
		ViewCount    int `json:"view_count"`    // 访问量
	}
	articleCount, _ := strconv.Atoi(h.repo.Find(c, g.KeyArticleCnt).(string))
	userCount, _ := strconv.Atoi(h.repo.Find(c, g.KeyUserCnt).(string))
	messageCount, _ := strconv.Atoi(h.repo.Find(c, g.KeyMessageCnt).(string))
	viewCount, _ := strconv.Atoi(h.repo.Find(c, g.KeyViewCnt).(string))
	ReturnSuccess(c, Resp{
		ArticleCount: articleCount,
		UserCount:    userCount,
		MessageCount: messageCount,
		ViewCount:    viewCount,
	})
}

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"webook-server/internal/repository"
	"webook-server/internal/repository/cache"
	dao2 "webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/web"
	"webook-server/ioc"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// ArticleTestSuit 测试套件
type ArticleTestSuit struct {
	suite.Suite
	r  *gin.Engine
	db *gorm.DB
}

func (s *ArticleTestSuit) SetupSuite() {

	//注册路由
	cmd := ioc.InitRedis()
	db := ioc.InitDB()
	cache := cache.NewArticleCache(cmd)
	dao := dao2.NewArticleDao(db)
	repo := repository.NewArticleRepository(dao, cache)
	svc := service.NewArticleService(repo)
	h := web.NewArticleHandler(svc)

	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("user_id", int64(1))
	})
	h.InitRouter(r, nil)
	s.r = r
	s.db = db
}

func (s *ArticleTestSuit) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE articles") // 清空数据 自增主键恢复0
}

// TestCreateOrUpdate 集成测试 TDD
func (s *ArticleTestSuit) TestCreateOrUpdate() {
	t := s.T()
	tcs := []struct {
		name string

		before func(t *testing.T) // 集成测试准备数据
		after  func(t *testing.T) // 集成测试验证数据

		req     Req
		gotCode int
		gotResp Response[int64]
	}{
		{
			name: "ok",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				//... 检查数据库
				// var article dao2.Article
				// err := s.db.Model(&dao2.Article{}).Where("").First(&article).Error // ...
				// assert.NoError(t, err)
				// assert.True(t, article.CreateAt > 0)
				// article.CreateAt = 0                     // ...
				// assert.Equal(t, dao2.Article{}, article) // ...
			},
			req: Req{
				Title:   "123",
				Content: "123",
			},
			gotCode: http.StatusOK,
			gotResp: Response[int64]{
				Code: 0,
			},
		},
	}
	// 更新 todo 创建 更新非自己
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			defer tc.after(t)

			reqBody, err := json.Marshal(tc.req)
			assert.NoError(t, err)
			// 准备Req和记录的 recorder
			req, err := http.NewRequest(http.MethodPost,
				"/api/article",
				bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			recorder := httptest.NewRecorder()

			// 执行
			s.r.ServeHTTP(recorder, req)
			// 断言结果
			assert.Equal(t, tc.gotCode, recorder.Code)
			if tc.gotCode != http.StatusOK {
				return
			}
			var res Response[int64]
			err = json.NewDecoder(recorder.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.gotResp, res)
		})
	}
}

type Req struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func TestArticle(t *testing.T) {
	suite.Run(t, &ArticleTestSuit{})
}

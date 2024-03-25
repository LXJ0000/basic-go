package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"webook-server/internal/service"
	"webook-server/internal/web"
)

// ArticleTestSuit 测试套件
type ArticleTestSuit struct {
	suite.Suite
	r *gin.Engine
}

func (s *ArticleTestSuit) SetupSuite() {
	r := gin.Default()
	//注册路由
	svc := service.NewArticleService()
	h := web.NewArticleHandler(svc)
	h.InitRouter(r)
	s.r = r
}

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

			},
			req: Req{
				Title:   "123",
				Content: "123",
			},
			gotCode: http.StatusOK,
			gotResp: Response[int64]{
				Code: 0,
				Data: 1,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			//	1. 构造请求
			tc.before(t)
			reqBody, err := json.Marshal(tc.req)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/api/article", bytes.NewBuffer([]byte(reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			//	2. 执行
			s.r.ServeHTTP(resp, req)
			//	3. 验证结果
			require.Equal(t, tc.gotCode, resp.Code)
			if resp.Code != http.StatusOK {
				return
			}
			var res Response[int64]
			err = json.NewDecoder(resp.Body).Decode(&res)
			require.NoError(t, err)
			require.Equal(t, tc.gotResp, res)
			tc.after(t)
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

package web

import (
	"testing"
	"webook-server/pkg/snowflake"

	"go.uber.org/mock/gomock"
)

func TestArticleHandler_Publish(t *testing.T) {
	snowflake.Init("2023-01-01", 1)
	tests := []struct {
		name string
		// mock    func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler)
		req     string
		gotCode int
		gotMsg  string
	}{
		{
			name: "success",
			// mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
			// 	// userSvc := svcmock.NewMockUserService(c)
			// 	// codeSvc := svcmock.NewMockCodeService(c)
			// 	// jwtHandler := pkgmock.NewMockJWTHandler(c)
			// 	// userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
			// 	// return userSvc, codeSvc, jwtHandler
			// },
			req: `

`,
			gotCode: 200,
			gotMsg:  "注册成功",
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// r := gin.Default()
			// h := NewUserHandler(tt.mock(c))
			// h.InitRouter(r, nil)

			// req, err := http.NewRequest(http.MethodPost,
			// 	"/user/register",
			// 	bytes.NewBuffer([]byte(tt.req)))
			// require.NoError(t, err)
			// req.Header.Set("Content-Type", "application/json")

			// resp := httptest.NewRecorder()
			// r.ServeHTTP(resp, req)

			// assert.Equal(t, tt.gotCode, resp.Code)
			// if resp.Code == http.StatusBadRequest {
			// 	return // 请求参数有误 则无法 json 解析
			// }

			var respBody Response
			err = json.Unmarshal(resp.Body.Bytes(), &respBody)
			// require.NoError(t, err)
			// assert.Equal(t, tt.gotMsg, respBody.Msg)

		})
	}
}

package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"webook-server/internal/service"
	svcmock "webook-server/internal/service/mocks"
	"webook-server/pkg/jwt"
	pkgmock "webook-server/pkg/jwt/mock"
	"webook-server/pkg/snowflake"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_Login(t *testing.T) {
	snowflake.Init("2023-01-01", 1)
	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler)
		req     string
		gotCode int
		gotMsg  string
	}{
		{
			name: "success",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)
				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq.com",
    "password":"hello#world123",
    "confirm_password":"hello#world123"
}
`,
			gotCode: 200,
			gotMsg:  "注册成功",
		},
		{
			name: "非法邮箱格式",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)
				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq",
    "password":"hello#world123",
    "confirm_password":"hello#world123"
}
`,
			gotCode: 200,
			gotMsg:  "非法邮箱格式",
		},
		{
			name: "密码不一致",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)
				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq.com",
    "password":"hello#world12",
    "confirm_password":"hello#world123"
}
`,
			gotCode: 200,
			gotMsg:  "密码不一致",
		},
		{
			name: "密码必须包含字母、数字、特殊字符，并且不少于八位",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)
				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq.com",
    "password":"123",
    "confirm_password":"123"
}
`,
			gotCode: 200,
			gotMsg:  "密码必须包含字母、数字、特殊字符，并且不少于八位",
		},
		{
			name: "请求参数有误",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)
				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
}
`,
			gotCode: http.StatusBadRequest,
			gotMsg:  "请求参数有误",
		},
		{
			name: "邮箱或用户名已存在",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)

				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(service.ErrDuplicate)
				return userSvc, codeSvc, jwtHandler
			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq.com",
    "password":"hello#world123",
    "confirm_password":"hello#world123"
}
`,
			gotCode: 200,
			gotMsg:  "邮箱或用户名已存在",
		},
		{
			name: "系统错误",
			mock: func(c *gomock.Controller) (service.UserService, service.CodeService, jwt.JWTHandler) {
				userSvc := svcmock.NewMockUserService(c)
				codeSvc := svcmock.NewMockCodeService(c)
				jwtHandler := pkgmock.NewMockJWTHandler(c)

				userSvc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(errors.New(""))
				return userSvc, codeSvc, jwtHandler

			},
			req: `
{
    "user_name":"root",
    "email":"1234@qq.com",
    "password":"hello#world123",
    "confirm_password":"hello#world123"
}
`,
			gotCode: 200,
			gotMsg:  "系统错误",
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			r := gin.Default()
			h := NewUserHandler(tt.mock(c))
			h.InitRouter(r, nil)

			req, err := http.NewRequest(http.MethodPost,
				"/user/register",
				bytes.NewBuffer([]byte(tt.req)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.gotCode, resp.Code)
			if resp.Code == http.StatusBadRequest {
				return // 请求参数有误 则无法 json 解析
			}

			var respBody Response
			err = json.Unmarshal(resp.Body.Bytes(), &respBody)
			require.NoError(t, err)
			assert.Equal(t, tt.gotMsg, respBody.Msg)

		})
	}
}

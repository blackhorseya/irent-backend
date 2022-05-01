package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/user/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http/middlewares"
	"github.com/blackhorseya/irent/pb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = "A111768050"

	pwd1 = "password"

	shaPWD1 = "0x5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

	user1 = &pb.Profile{
		Id: id1,
	}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.ResponseMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Login() {
	s.r.POST("/api/v1/auth/login", s.handler.Login)

	type args struct {
		id       string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing id then error",
			args:     args{id: "", password: pwd1},
			wantCode: 400,
		},
		{
			name:     "missing password then error",
			args:     args{id: id1, password: ""},
			wantCode: 400,
		},
		{
			name: "login then error",
			args: args{id: id1, password: pwd1, mock: func() {
				s.mock.On("Login", mock.Anything, id1, pwd1).Return(nil, er.ErrLogin).Once()
			}},
			wantCode: 500,
		},
		{
			name: "login then success",
			args: args{id: id1, password: pwd1, mock: func() {
				s.mock.On("Login", mock.Anything, id1, pwd1).Return(user1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/auth/login")
			values := url.Values{
				"id":       []string{tt.args.id},
				"password": []string{tt.args.password},
			}
			req := httptest.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Login() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

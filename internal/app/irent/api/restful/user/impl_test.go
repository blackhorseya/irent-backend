package user

import (
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/pb"
	"github.com/blackhorseya/irent/test/testdata"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/user/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = "A111768050"

	pwd1 = "password"
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.r.Use(ginhttp.AddContextx())
	s.r.Use(ginhttp.HandleError())

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
		wantBody *response.Response
	}{
		{
			name:     "missing id then error",
			args:     args{id: "", password: pwd1},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "missing password then error",
			args:     args{id: id1, password: ""},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "login then error",
			args: args{id: id1, password: pwd1, mock: func() {
				s.mock.On("Login", mock.Anything, id1, pwd1).Return(nil, er.ErrLogin).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "login then success",
			args: args{id: id1, password: pwd1, mock: func() {
				s.mock.On("Login", mock.Anything, id1, pwd1).Return(testdata.User1, nil).Once()
			}},
			wantCode: 201,
			wantBody: response.OK.WithData(user.NewProfileResponse(testdata.User1)),
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

			type resp struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data *pb.Profile `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Login() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data.(*pb.Profile)) {
					t.Errorf("Login() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

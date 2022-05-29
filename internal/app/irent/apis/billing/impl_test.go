package billing

import (
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http/middlewares"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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

func (s *handlerSuite) Test_impl_GetArrears() {
	s.r.GET("/api/v1/billing/:id/arrears", middlewares.AuthMiddleware(), s.handler.GetArrears)

	type args struct {
		id    string
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing id then error",
			args:     args{id: "", token: testdata.User1.AccessToken},
			wantCode: 400,
		},
		{
			name:     "missing token then error",
			args:     args{id: testdata.User1.ID, token: ""},
			wantCode: 401,
		},
		{
			name: "get arrears then error",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetArrears", mock.Anything, testdata.User1).Return(nil, er.ErrQueryArrears).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get arrears then success",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetArrears", mock.Anything, testdata.User1).Return(testdata.Arrears1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/billing/%s/arrears", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetArrears() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

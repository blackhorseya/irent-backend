package cars

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/car/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http/middlewares"
	"github.com/blackhorseya/irent/pb"
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

func (s *handlerSuite) Test_impl_NearTopN() {
	s.r.GET("/api/v1/car/near", s.handler.NearTopN)

	type args struct {
		n         string
		latitude  string
		longitude string
		mock      func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "n is string then error",
			args:     args{n: "string", latitude: "0", longitude: "0"},
			wantCode: 400,
		},
		{
			name:     "latitude is string then error",
			args:     args{n: "10", latitude: "string", longitude: "0"},
			wantCode: 400,
		},
		{
			name:     "longitude is string then error",
			args:     args{n: "10", latitude: "0", longitude: "string"},
			wantCode: 400,
		},
		{
			name: "near then error",
			args: args{n: "10", latitude: "0", longitude: "0", mock: func() {
				s.mock.On("NearTopN", mock.Anything, 10, float64(0), float64(0)).Return(nil, 0, er.ErrListCars).Once()
			}},
			wantCode: 500,
		},
		{
			name: "near then success",
			args: args{n: "10", latitude: "0", longitude: "0", mock: func() {
				s.mock.On("NearTopN", mock.Anything, 10, float64(0), float64(0)).Return([]*pb.Car{testdata.Car1}, 5, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/car/near?n=%v&latitude=%v&longitude=%v", tt.args.n, tt.args.latitude, tt.args.longitude)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "NearTopN() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

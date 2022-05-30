package cars

import (
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/pb"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type suiteHandler struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *suiteHandler) SetupTest() {
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

func (s *suiteHandler) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestSuiteHandler(t *testing.T) {
	suite.Run(t, new(suiteHandler))
}

func (s *suiteHandler) Test_impl_NearTopN() {
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
		wantBody *response.Response
	}{
		{
			name:     "n is string then error",
			args:     args{n: "string", latitude: "0", longitude: "0"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "latitude is string then error",
			args:     args{n: "10", latitude: "string", longitude: "0"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "longitude is string then error",
			args:     args{n: "10", latitude: "0", longitude: "string"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "near then error",
			args: args{n: "10", latitude: "0", longitude: "0", mock: func() {
				s.mock.On("NearTopN", mock.Anything, 10, float64(0), float64(0)).Return(nil, 0, er.ErrListCars).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "near then success",
			args: args{n: "10", latitude: "0", longitude: "0", mock: func() {
				s.mock.On("NearTopN", mock.Anything, 10, float64(0), float64(0)).Return([]*car.Info{testdata.Car1}, 5, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData([]*pb.Car{car.NewCarResponse(testdata.Car1)}),
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

			type resp struct {
				Code int       `json:"code"`
				Msg  string    `json:"msg"`
				Data []*pb.Car `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "NearTopN() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data) {
					t.Errorf("NearTopN() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

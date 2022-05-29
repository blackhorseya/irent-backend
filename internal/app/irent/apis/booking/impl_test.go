package booking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/order/mocks"
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

func (s *handlerSuite) Test_impl_ListBookings() {
	s.r.GET("/api/v1/bookings", middlewares.AuthMiddleware(), s.handler.ListBookings)

	type args struct {
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing token then error",
			args:     args{token: ""},
			wantCode: 401,
		},
		{
			name: "list then error",
			args: args{token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("List", mock.Anything, 0, 0, mock.Anything).Return(nil, er.ErrListBooking).Once()
			}},
			wantCode: 500,
		},
		{
			name: "list then success",
			args: args{token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("List", mock.Anything, 0, 0, mock.Anything).Return([]*pb.OrderInfo{testdata.Order1}, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/bookings")
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ListBookings() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_GetBookingByID() {
	s.r.GET("/api/v1/bookings/:id", middlewares.AuthMiddleware(), s.handler.GetBookingByID)

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
			wantCode: 404,
		},
		{
			name:     "missing token then error",
			args:     args{id: testdata.User1.ID, token: ""},
			wantCode: 401,
		},
		{
			name: "get by id then error",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.User1.ID, testdata.User1).Return(nil, er.ErrGetBookingByID).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.User1.ID, testdata.User1).Return(testdata.Order1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/bookings/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetBookingByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Book() {
	s.r.POST("/api/v1/bookings", middlewares.AuthMiddleware(), s.handler.Book)

	type args struct {
		id     string
		projID string
		token  string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "book then error",
			args: args{id: testdata.User1.ID, projID: testdata.ProjID1, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("BookCar", mock.Anything, testdata.User1.ID, testdata.ProjID1, testdata.User1).Return(nil, er.ErrBook).Once()
			}},
			wantCode: 500,
		},
		{
			name: "book then success",
			args: args{id: testdata.User1.ID, projID: testdata.ProjID1, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("BookCar", mock.Anything, testdata.User1.ID, testdata.ProjID1, testdata.User1).Return(testdata.Booking1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/bookings")
			data, _ := json.Marshal(&bookRequest{ID: tt.args.id, ProjectID: tt.args.projID})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Book() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_CancelBooking() {
	s.r.DELETE("/api/v1/bookings/:id", middlewares.AuthMiddleware(), s.handler.CancelBooking)

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
			wantCode: 404,
		},
		{
			name:     "missing token then error",
			args:     args{id: testdata.User1.ID, token: ""},
			wantCode: 401,
		},
		{
			name: "cancel booking then error",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.User1.ID, testdata.User1).Return(er.ErrCancelBooking).Once()
			}},
			wantCode: 500,
		},
		{
			name: "cancel booking then success",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.User1.ID, testdata.User1).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/bookings/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "CancelBooking() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

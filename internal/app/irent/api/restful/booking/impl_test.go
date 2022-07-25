package booking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/pb"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/irent/biz/order/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/http/middlewares"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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

func (s *suiteHandler) Test_impl_ListBookings() {
	s.r.GET("/api/v1/bookings", middlewares.AuthMiddleware(), s.handler.ListBookings)

	type args struct {
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name:     "missing token then error",
			args:     args{token: ""},
			wantCode: 401,
			wantBody: nil,
		},
		{
			name: "list then error",
			args: args{token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("List", mock.Anything, 0, 0, mock.Anything).Return(nil, er.ErrListBooking).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "list then success",
			args: args{token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("List", mock.Anything, 0, 0, mock.Anything).Return([]*order.Info{testdata.Order1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData([]*pb.OrderInfo{order.NewOrderInfoResponse(testdata.Order1)}),
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

			type resp struct {
				Code int             `json:"code"`
				Msg  string          `json:"msg"`
				Data []*pb.OrderInfo `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ListBookings() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data) {
					t.Errorf("ListBookings() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteHandler) Test_impl_GetBookingByID() {
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
		wantBody *response.Response
	}{
		{
			name:     "missing token then error",
			args:     args{id: testdata.User1.ID, token: ""},
			wantCode: 401,
			wantBody: nil,
		},
		{
			name: "get by id then error",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.User1.ID, testdata.User1).Return(nil, er.ErrGetBookingByID).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.User1.ID, testdata.User1).Return(testdata.Order1, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData(order.NewOrderInfoResponse(testdata.Order1)),
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

			type resp struct {
				Code int           `json:"code"`
				Msg  string        `json:"msg"`
				Data *pb.OrderInfo `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetBookingByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data) {
					t.Errorf("GetBookingByID() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteHandler) Test_impl_Book() {
	s.r.POST("/api/v1/bookings", middlewares.AuthMiddleware(), s.handler.Book)

	type args struct {
		id     string
		projID string
		token  string
		userID string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name: "book then error",
			args: args{id: testdata.User1.ID, projID: testdata.ProjID1, token: testdata.User1.AccessToken, userID: testdata.User1.ID, mock: func() {
				s.mock.On("BookCar", mock.Anything, testdata.User1.ID, testdata.ProjID1, testdata.User1, false).Return(nil, er.ErrBook).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "book then success",
			args: args{id: testdata.User1.ID, projID: testdata.ProjID1, token: testdata.User1.AccessToken, userID: testdata.User1.ID, mock: func() {
				s.mock.On("BookCar", mock.Anything, testdata.User1.ID, testdata.ProjID1, testdata.User1, false).Return(testdata.Booking1, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData(order.NewBookingResponse(testdata.Booking1)),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/bookings")
			data, _ := json.Marshal(&bookRequest{ID: tt.args.id, ProjectID: tt.args.projID, UserID: tt.args.userID})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.args.token))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			type resp struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data *pb.Booking `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Book() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data.(*pb.Booking)) {
					t.Errorf("Book() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteHandler) Test_impl_CancelBooking() {
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
		wantBody *response.Response
	}{
		{
			name:     "missing token then error",
			args:     args{id: testdata.User1.ID, token: ""},
			wantCode: 401,
			wantBody: nil,
		},
		{
			name: "cancel booking then error",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.User1.ID, testdata.User1).Return(er.ErrCancelBooking).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "cancel booking then success",
			args: args{id: testdata.User1.ID, token: testdata.User1.AccessToken, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.User1.ID, testdata.User1).Return(nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData(testdata.User1.ID),
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

			type resp struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
				Data string `json:"data"`
			}

			var gotBody *resp
			err := json.NewDecoder(got.Body).Decode(&gotBody)
			if err != nil {
				t.Errorf("Decode response body error = %v", err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "CancelBooking() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			if 200 <= tt.wantCode && tt.wantCode < 300 {
				if !reflect.DeepEqual(gotBody.Data, tt.wantBody.Data) {
					t.Errorf("CancelBooking() gotBody = %v, want %v", gotBody, tt.wantBody)
				}
			}

			s.TearDownTest()
		})
	}
}

package order

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo/mocks"
	"github.com/blackhorseya/irent/pb"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_List() {
	type args struct {
		start int
		end   int
		user  *user.Profile
		mock  func()
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []*pb.OrderInfo
		wantErr    bool
	}{
		{
			name:       "start < 0 then error",
			args:       args{start: -1, end: 3, user: testdata.User1},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name:       "end < 0 then error",
			args:       args{start: 0, end: -1, user: testdata.User1},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name:       "missing token then error",
			args:       args{start: 0, end: 2, user: &user.Profile{}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "list then error",
			args: args{start: 0, end: 2, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return(nil, errors.New("error")).Once()
			}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "list then not found error",
			args: args{start: 0, end: 2, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return(nil, nil).Once()
			}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "list then success",
			args: args{start: 0, end: 2, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return([]*pb.OrderInfo{testdata.Order1}, nil).Once()
			}},
			wantOrders: []*pb.OrderInfo{testdata.Order1},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotOrders, err := s.biz.List(contextx.Background(), tt.args.start, tt.args.end, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("List() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_GetByID() {
	type args struct {
		id   string
		user *user.Profile
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *pb.OrderInfo
		wantErr  bool
	}{
		{
			name:     "missing token then error",
			args:     args{id: testdata.Order1.No, user: &user.Profile{}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing id then error",
			args:     args{id: "", user: testdata.User1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: testdata.Order1.No, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found error",
			args: args{id: testdata.Order1.No, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.Order1.No, user: testdata.User1, mock: func() {
				s.mock.On("QueryBookings", mock.Anything, testdata.User1).Return([]*pb.OrderInfo{testdata.Order1}, nil).Once()
			}},
			wantInfo: testdata.Order1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetByID(contextx.Background(), tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_BookCar() {
	type args struct {
		id     string
		projID string
		user   *user.Profile
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *pb.Booking
		wantErr  bool
	}{
		{
			name:     "missing id then error",
			args:     args{id: "", projID: testdata.ProjID1, user: testdata.User1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing project id then error",
			args:     args{id: testdata.Car1.ID, projID: "", user: testdata.User1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing token then error",
			args:     args{id: testdata.Car1.ID, projID: testdata.ProjID1, user: &user.Profile{}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "book then error",
			args: args{id: testdata.Car1.ID, projID: testdata.ProjID1, user: testdata.User1, mock: func() {
				s.mock.On("Book", mock.Anything, testdata.Car1.ID, testdata.ProjID1, testdata.User1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "book then success",
			args: args{id: testdata.Car1.ID, projID: testdata.ProjID1, user: testdata.User1, mock: func() {
				s.mock.On("Book", mock.Anything, testdata.Car1.ID, testdata.ProjID1, testdata.User1).Return(testdata.Booking1, nil).Once()
			}},
			wantInfo: testdata.Booking1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.BookCar(contextx.Background(), tt.args.id, tt.args.projID, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("BookCar() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_CancelBooking() {
	type args struct {
		id   string
		user *user.Profile
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "missing id then error",
			args:    args{id: "", user: testdata.User1},
			wantErr: true,
		},
		{
			name:    "missing token then error",
			args:    args{id: testdata.Car1.ID, user: &user.Profile{}},
			wantErr: true,
		},
		{
			name: "cancel booking then error",
			args: args{id: testdata.Car1.ID, user: testdata.User1, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.Car1.ID, testdata.User1).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "cancel booking then success",
			args: args{id: testdata.Car1.ID, user: testdata.User1, mock: func() {
				s.mock.On("CancelBooking", mock.Anything, testdata.Car1.ID, testdata.User1).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.CancelBooking(contextx.Background(), tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CancelBooking() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

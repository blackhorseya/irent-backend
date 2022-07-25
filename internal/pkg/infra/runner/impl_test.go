package runner

import (
	"github.com/blackhorseya/irent/internal/app/irent/biz/order/mocks"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
	"time"
)

type SuiteTest struct {
	suite.Suite
	orderBiz *mocks.IBiz
	runner   Runner
}

func (s *SuiteTest) SetupTest() {
	logger := zap.NewNop()
	s.orderBiz = new(mocks.IBiz)

	runner, err := CreateRunner(&Options{Shift: 5 * time.Minute}, logger, s.orderBiz)
	if err != nil {
		panic(err)
	}
	s.runner = runner
}

func (s *SuiteTest) TearDownTest() {
	s.orderBiz.AssertExpectations(s.T())
}

func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

func (s *SuiteTest) Test_impl_Execute() {
	type args struct {
		t    time.Time
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "list premium bookings then error",
			args: args{mock: func() {
				s.orderBiz.On("ListPremiumBookings", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "list premium bookings not found then nil",
			args: args{mock: func() {
				s.orderBiz.On("ListPremiumBookings", mock.Anything).Return(nil, nil).Once()
			}},
			wantErr: false,
		},
		{
			name: "rebook then error",
			args: args{t: time.Now(), mock: func() {
				s.orderBiz.On("ListPremiumBookings", mock.Anything).Return(map[*user.Profile]*order.Booking{
					testdata.User1: {
						No:         "no1",
						LastPickAt: time.Now().Add(3 * time.Minute),
						CarID:      "id1",
						ProjID:     "proj1",
					},
				}, nil).Once()

				s.orderBiz.On("ReBookCar", mock.Anything, "no1", "id1", "proj1", testdata.User1).Return(nil, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "rebook then update success",
			args: args{t: time.Now(), mock: func() {
				s.orderBiz.On("ListPremiumBookings", mock.Anything).Return(map[*user.Profile]*order.Booking{
					testdata.User1: {
						No:         "no1",
						LastPickAt: time.Now().Add(3 * time.Minute),
						CarID:      "id1",
						ProjID:     "proj1",
					},
				}, nil).Once()

				s.orderBiz.On("ReBookCar", mock.Anything, "no1", "id1", "proj1", testdata.User1).Return(testdata.Booking1, nil).Once()
				s.orderBiz.On("UpdatePremiumBooking", mock.Anything, testdata.User1, testdata.Booking1).Return(testdata.Booking1, nil).Once()
			}},
			wantErr: false,
		},
		{
			name: "no execute rebook car then nil",
			args: args{t: time.Now(), mock: func() {
				s.orderBiz.On("ListPremiumBookings", mock.Anything).Return(map[*user.Profile]*order.Booking{
					testdata.User1: {
						No:         "no1",
						LastPickAt: time.Now().Add(15 * time.Minute),
						CarID:      "id1",
						ProjID:     "proj1",
					},
				}, nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.runner.(*impl).Execute(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

package car

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/car/repo/mocks"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteBiz struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *suiteBiz) SetupTest() {
	logger, _ := zap.NewDevelopment()

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *suiteBiz) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestSuiteBiz(t *testing.T) {
	suite.Run(t, new(suiteBiz))
}

func (s *suiteBiz) Test_impl_NearTopN() {
	type args struct {
		top       int
		latitude  float64
		longitude float64
		mock      func()
	}
	tests := []struct {
		name      string
		args      args
		wantCars  []*car.Info
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "top <= 0 then error",
			args:      args{top: 0},
			wantCars:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list then error",
			args: args{top: 10, mock: func() {
				s.mock.On("List", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantCars:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list then not found error",
			args: args{top: 10, mock: func() {
				s.mock.On("List", mock.Anything).Return(nil, nil).Once()
			}},
			wantCars:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list then success",
			args: args{top: 1, latitude: 0, longitude: 0, mock: func() {
				s.mock.On("List", mock.Anything).Return([]*car.Info{testdata.Car1}, nil).Once()
			}},
			wantCars:  []*car.Info{testdata.Car1},
			wantTotal: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotCars, gotTotal, err := s.biz.NearTopN(contextx.Background(), tt.args.top, tt.args.latitude, tt.args.longitude)
			if (err != nil) != tt.wantErr {
				t.Errorf("NearTopN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCars, tt.wantCars) {
				t.Errorf("NearTopN() gotCars = %v, want %v", gotCars, tt.wantCars)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("NearTopN() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			s.TearDownTest()
		})
	}
}

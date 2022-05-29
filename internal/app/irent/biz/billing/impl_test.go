package billing

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo/mocks"
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
	logger := zap.NewNop()

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

func (s *bizSuite) Test_impl_GetArrears() {
	type args struct {
		from *user.Profile
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Arrears
		wantErr  bool
	}{
		{
			name:     "missing id then error",
			args:     args{from: &user.Profile{ID: "", AccessToken: "token"}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing token then error",
			args:     args{from: &user.Profile{ID: "id", AccessToken: ""}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get arrears then error",
			args: args{from: testdata.User1, mock: func() {
				s.mock.On("QueryArrears", mock.Anything, testdata.User1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get arrears then success",
			args: args{from: testdata.User1, mock: func() {
				s.mock.On("QueryArrears", mock.Anything, testdata.User1).Return(testdata.Arrears1, nil).Once()
			}},
			wantInfo: testdata.Arrears1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetArrears(contextx.Background(), tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetArrears() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

package user

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo/mocks"
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

func (s *bizSuite) Test_impl_Login() {
	type args struct {
		id       string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "missing id then error",
			args:     args{id: "", password: "password"},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing password then error",
			args:     args{id: testdata.User1.ID, password: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then error",
			args: args{id: testdata.User1.ID, password: "password", mock: func() {
				s.mock.On("Login", mock.Anything, testdata.User1.ID, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then success",
			args: args{id: testdata.User1.ID, password: "password", mock: func() {
				s.mock.On("Login", mock.Anything, testdata.User1.ID, mock.Anything).Return(testdata.User1, nil).Once()
			}},
			wantInfo: testdata.User1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Login(contextx.Background(), tt.args.id, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

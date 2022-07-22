package repo

import (
	"bytes"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient/mocks"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type repoSuite struct {
	suite.Suite
	client *mocks.HTTPClient
	repo   IRepo
}

func (s *repoSuite) SetupTest() {
	s.client = new(mocks.HTTPClient)
	repo, err := CreateIRepo(&Options{
		Endpoint:   "https://irentcar-app.azurefd.net/api",
		AppVersion: "5.8.1",
	}, s.client)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func (s *repoSuite) TearDownTest() {
	s.client.AssertExpectations(s.T())
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Login() {
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
			name: "call http then error",
			args: args{id: testdata.User1.ID, password: "password", mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then failed",
			args: args{id: testdata.User1.ID, password: "password", mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage": "failed"}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then success",
			args: args{id: testdata.User1.ID, password: "password", mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage":"Success","Data":{"UserData":{"MEMIDNO":"1"}}}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantInfo: &user.Profile{ID: "1", Name: "", AccessToken: ""},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Login(contextx.Background(), tt.args.id, tt.args.password)
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

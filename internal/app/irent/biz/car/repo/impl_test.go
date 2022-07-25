package repo

import (
	"bytes"
	"encoding/json"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/stretchr/testify/suite"
)

type repoSuite struct {
	suite.Suite
	client *mocks.HTTPClient
	repo   IRepo
}

func (s *repoSuite) SetupTest() {
	s.client = new(mocks.HTTPClient)
	repo, err := CreateIRepo(&Options{
		Endpoint: "https://irentcar-app.azurefd.net/api",
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

func (s *repoSuite) Test_impl_List() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCars []*car.Info
		wantErr  bool
	}{
		{
			name: "call http then error",
			args: args{mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantCars: nil,
			wantErr:  true,
		},
		{
			name: "call http then failed",
			args: args{mock: func() {
				data, _ := json.Marshal(&listResp{Errormessage: "failed"})
				body := ioutil.NopCloser(bytes.NewReader(data))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil)
			}},
			wantCars: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotCars, err := s.repo.List(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCars, tt.wantCars) {
				t.Errorf("List() gotCars = %v, want %v", gotCars, tt.wantCars)
			}

			s.TearDownTest()
		})
	}
}

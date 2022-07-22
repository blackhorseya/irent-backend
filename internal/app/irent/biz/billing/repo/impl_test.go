package repo

import (
	"bytes"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/test/testdata"
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

func (s *repoSuite) Test_impl_QueryArrears() {
	type args struct {
		from *user.Profile
		mock func()
	}
	tests := []struct {
		name        string
		args        args
		wantArrears *user.Arrears
		wantErr     bool
	}{
		{
			name: "call http then error",
			args: args{from: testdata.User1, mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantArrears: nil,
			wantErr:     true,
		},
		{
			name: "query arrears then return failed",
			args: args{from: testdata.User1, mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage": "failed"}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantArrears: nil,
			wantErr:     true,
		},
		{
			name: "query arrears then return success",
			args: args{from: testdata.User1, mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage":"Success","Data":{"ArrearsInfos":[{"OrderNo":"1","Total_Amount":1}],"TotalAmount":1}}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantArrears: &user.Arrears{
				Records:     []*user.ArrearsRecord{{OrderNo: "1", TotalAmount: 1}},
				TotalAmount: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotArrears, err := s.repo.QueryArrears(contextx.Background(), tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryArrears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArrears, tt.wantArrears) {
				t.Errorf("QueryArrears() gotArrears = %v, want %v", gotArrears, tt.wantArrears)
			}

			s.TearDownTest()
		})
	}
}

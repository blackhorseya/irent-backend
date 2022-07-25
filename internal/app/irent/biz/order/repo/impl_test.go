package repo

import (
	"bytes"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient/mocks"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteRepo struct {
	suite.Suite
	client *mocks.HTTPClient
	repo   IRepo
}

func (s *suiteRepo) SetupTest() {
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

func (s *suiteRepo) TearDownTest() {
	s.client.AssertExpectations(s.T())
}

func TestSuiteRepo(t *testing.T) {
	suite.Run(t, new(suiteRepo))
}

func (s *suiteRepo) Test_impl_QueryBookings() {
	type args struct {
		from *user.Profile
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []*order.Info
		wantErr    bool
	}{
		{
			name: "call http then error",
			args: args{from: testdata.User1, mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "call http then failed",
			args: args{from: testdata.User1, mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage": "failed"}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantOrders: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotOrders, err := s.repo.QueryBookings(contextx.Background(), tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBookings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("QueryBookings() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteRepo) Test_impl_Book() {
	type args struct {
		id     string
		projID string
		from   *user.Profile
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *order.Booking
		wantErr  bool
	}{
		{
			name: "call http then error",
			args: args{from: testdata.User1, mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "call http then failed",
			args: args{from: testdata.User1, mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage": "failed"}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Book(contextx.Background(), tt.args.id, tt.args.projID, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("Book() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Book() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *suiteRepo) Test_impl_CancelBooking() {
	type args struct {
		id   string
		from *user.Profile
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "call http then error",
			args: args{from: testdata.User1, mock: func() {
				s.client.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "call http then error",
			args: args{from: testdata.User1, mock: func() {
				body := ioutil.NopCloser(bytes.NewReader([]byte(`{"ErrorMessage": "failed"}`)))
				s.client.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CancelBooking(contextx.Background(), tt.args.id, tt.args.from); (err != nil) != tt.wantErr {
				t.Errorf("CancelBooking() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

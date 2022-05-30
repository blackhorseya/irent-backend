//go:build integration

package repo

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/stretchr/testify/suite"
)

type suiteRepo struct {
	suite.Suite
	repo IRepo
}

func (s *suiteRepo) SetupTest() {
	repo, err := CreateIRepo(&Options{
		Endpoint:   "https://irentcar-app.azurefd.net/api",
		AppVersion: "5.8.1",
	})
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestSuiteRepo(t *testing.T) {
	suite.Run(t, new(suiteRepo))
}

func (s *suiteRepo) Test_impl_QueryBookings() {
	type args struct {
		from *user.Profile
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []*order.Info
		wantErr    bool
	}{
		{
			name:       "query bookings then success",
			args:       args{from: testdata.User1},
			wantOrders: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotOrders, err := s.repo.QueryBookings(contextx.Background(), tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBookings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("QueryBookings() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func (s *suiteRepo) Test_impl_CancelBooking() {
	type args struct {
		id   string
		from *user.Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "cancel booking then success",
			args:    args{id: testdata.Order1.No, from: testdata.User1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.repo.CancelBooking(contextx.Background(), tt.args.id, tt.args.from); (err != nil) != tt.wantErr {
				t.Errorf("CancelBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *suiteRepo) Test_impl_Book() {
	type args struct {
		id     string
		projID string
		from   *user.Profile
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *order.Booking
		wantErr  bool
	}{
		{
			name:     "book car then success",
			args:     args{id: testdata.Car1.ID, projID: testdata.ProjID1, from: testdata.User1},
			wantInfo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.Book(contextx.Background(), tt.args.id, tt.args.projID, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("Book() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Book() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

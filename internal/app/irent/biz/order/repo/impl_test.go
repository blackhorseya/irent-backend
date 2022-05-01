//go:build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/stretchr/testify/suite"
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	repo, err := CreateIRepo(&Options{
		Endpoint:   "https://irentcar-app.azurefd.net/api",
		AppVersion: "5.8.1",
	})
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_QueryBookings() {
	type args struct {
		user *pb.Profile
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []*pb.OrderInfo
		wantErr    bool
	}{
		{
			name:       "query bookings then success",
			args:       args{user: testdata.User1},
			wantOrders: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotOrders, err := s.repo.QueryBookings(contextx.Background(), tt.args.user)
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

func (s *repoSuite) Test_impl_CancelBooking() {
	type args struct {
		id   string
		user *pb.Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "cancel booking then success",
			args:    args{id: testdata.Order1.No, user: testdata.User1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.repo.CancelBooking(contextx.Background(), tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CancelBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Book() {
	type args struct {
		id     string
		projID string
		user   *pb.Profile
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *pb.Booking
		wantErr  bool
	}{
		{
			name:     "book car then success",
			args:     args{id: testdata.Car1.Id, projID: testdata.ProjID1, user: testdata.User1},
			wantInfo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.Book(contextx.Background(), tt.args.id, tt.args.projID, tt.args.user)
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

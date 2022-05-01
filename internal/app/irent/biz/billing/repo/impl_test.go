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

func (s *repoSuite) Test_impl_QueryArrears() {
	type args struct {
		user *pb.Profile
	}
	tests := []struct {
		name        string
		args        args
		wantArrears *pb.Arrears
		wantErr     bool
	}{
		{
			name:        "query arrears then success",
			args:        args{user: testdata.User1},
			wantArrears: nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotArrears, err := s.repo.QueryArrears(contextx.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryArrears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArrears, tt.wantArrears) {
				t.Errorf("QueryArrears() gotArrears = %v, want %v", gotArrears, tt.wantArrears)
			}
		})
	}
}

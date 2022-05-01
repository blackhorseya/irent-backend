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

func (s *repoSuite) Test_impl_Login() {
	type args struct {
		id       string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *pb.Profile
		wantErr  bool
	}{
		{
			name:     "login then success",
			args:     args{id: testdata.User1.Id, password: "password"},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.Login(contextx.Background(), tt.args.id, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

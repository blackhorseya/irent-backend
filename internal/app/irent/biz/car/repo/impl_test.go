//go:build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/stretchr/testify/suite"
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	repo, err := CreateIRepo(&Options{
		Endpoint: "https://irentcar-app.azurefd.net/api",
	})
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_List() {
	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		wantCars []*pb.Car
		wantErr  bool
	}{
		{
			name:     "any rent then success",
			args:     args{},
			wantCars: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotCars, err := s.repo.List(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCars, tt.wantCars) {
				t.Errorf("List() gotCars = %v, want %v", gotCars, tt.wantCars)
			}
		})
	}
}

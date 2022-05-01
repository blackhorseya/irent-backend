package contextx

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type contextxSuite struct {
	suite.Suite
}

func TestContextxSuite(t *testing.T) {
	suite.Run(t, new(contextxSuite))
}

func (s *contextxSuite) TestBackground() {
	tests := []struct {
		name string
		want Contextx
	}{
		{
			name: "background then contextx",
			want: Background(),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := Background(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *contextxSuite) TestWithValue() {
	ctx := WithValue(Background(), "key", "value")

	type args struct {
		parent Contextx
		key    string
		val    interface{}
	}
	tests := []struct {
		name string
		args args
		want Contextx
	}{
		{
			name: "withValue then contextx",
			args: args{Background(), "key", "value"},
			want: ctx,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := WithValue(tt.args.parent, tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *contextxSuite) TestWithCancel() {
	ctx, cancel := WithCancel(Background())
	defer cancel()

	type args struct {
		parent Contextx
	}
	tests := []struct {
		name  string
		args  args
		want  Contextx
		want1 context.CancelFunc
	}{
		{
			name:  "withCancel then contextx func",
			args:  args{Background()},
			want:  ctx,
			want1: cancel,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, _ := WithCancel(tt.args.parent)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCancel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *contextxSuite) TestWithTimeout() {
	ctx, cancel := WithTimeout(Background(), 10*time.Second)
	defer cancel()

	type args struct {
		parent Contextx
		d      time.Duration
	}
	tests := []struct {
		name  string
		args  args
		want  Contextx
		want1 context.CancelFunc
	}{
		{
			name:  "withTimeout then contextx func",
			args:  args{Background(), 10 * time.Second},
			want:  ctx,
			want1: cancel,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			_, _ = WithTimeout(tt.args.parent, tt.args.d)
		})
	}
}

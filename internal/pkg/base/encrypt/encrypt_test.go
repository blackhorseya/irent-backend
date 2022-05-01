package encrypt

import (
	"testing"
)

var (
	pwd1 = "password"

	shaPWD1 = "0x5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
)

func TestEncPWD(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "encrypt then password",
			args: args{pwd: pwd1},
			want: shaPWD1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncPWD(tt.args.pwd); got != tt.want {
				t.Errorf("EncPWD() = %v, want %v", got, tt.want)
			}
		})
	}
}

package timex

import (
	"reflect"
	"testing"
	"time"
)

var (
	unixNano = int64(1610548520788105000)

	loc, _ = time.LoadLocation("Asia/Taipei")

	layout = "2006-01-02 15:04:05"

	time1, _ = time.ParseInLocation(layout, "2021-05-14 22:58:20", loc)

	layout2 = "20060102150405"

	time2, _ = time.ParseInLocation(layout2, "20210515100347", loc)
)

func TestUnix(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "1610548520788105000 then time",
			args: args{t: unixNano},
			want: time.Unix(unixNano/1e9, unixNano%1e9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unix(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseString2Time(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2021-05-14 22:58:20",
			args: args{str: "2021-05-14 22:58:20"},
			want: time1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseString2Time(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString2Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYYYYMMddHHmmss(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "20210515100347 then time",
			args: args{str: "20210515100347"},
			want: time2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseYYYYMMddHHmmss(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYYYYMMddHHmmss() = %v, want %v", got, tt.want)
			}
		})
	}
}

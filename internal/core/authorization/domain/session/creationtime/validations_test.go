package creationtime

import (
	"testing"
	"time"
)

func Test_isTooOld(t *testing.T) {
	type args struct {
		when time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "creation time too old",
			args: args{when: time.Now().Add(-42 * 30 * 24 * time.Hour)},
			want: true,
		},
		{
			name: "creation time doesn't exceed max age",
			args: args{when: time.Now().Add(-24 * time.Hour)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTooOld(tt.args.when); got != tt.want {
				t.Errorf("isTooOld() = %v, want %v", got, tt.want)
			}
		})
	}
}

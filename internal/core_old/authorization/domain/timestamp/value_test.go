package timestamp

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		when time.Time
	}
	type want struct {
		value Value
		err   error
	}

	val := time.Now()

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "initializes a new timestamp",
			args: args{when: val},
			want: want{
				value: Value{timestamp: val.UTC()},
				err:   nil,
			},
		},
		{
			name: "fails to initialize a new timestamp, because moment is empty",
			args: args{when: time.Time{}},
			want: want{
				value: Value{},
				err:   ErrEmpty,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.when)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("New() error = %v, wantErr %v", err, tt.want.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.value) {
				t.Errorf("New() got = %v, want %v", got, tt.want.value)
			}
		})
	}
}

func TestValue_Time(t *testing.T) {
	type fields struct {
		value time.Time
	}

	val := time.Now().Add(-7 * 24 * time.Hour)

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "returns the inner timestamp",
			fields: fields{value: val},
			want:   val,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				timestamp: tt.fields.value,
			}
			if got := v.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

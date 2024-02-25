package creationtime

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

	now := time.Now()
	old := time.Now().Add(-343 * 24 * time.Hour)

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "initializes a new creationtime value",
			args: args{when: now},
			want: want{
				value: Value{value: now.UTC()},
				err:   nil,
			},
		},
		{
			name: "fails to initialize a new creationtime value, because the created time is way to old",
			args: args{when: old},
			want: want{
				value: Value{},
				err:   ErrCreationTimeExceedMaxAge,
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

func TestValue_IsTooOld(t *testing.T) {
	type fields struct {
		value time.Time
	}
	type args struct {
		maxAge time.Duration
	}

	val := time.Now().UTC().Add(-7 * 24 * time.Hour)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "is not too old",
			fields: fields{value: val},
			args:   args{maxAge: 30 * 24 * time.Hour},
			want:   false,
		},
		{
			name:   "is too old",
			fields: fields{value: val},
			args:   args{maxAge: 24 * time.Hour},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				value: tt.fields.value,
			}
			if got := v.IsTooOld(tt.args.maxAge); got != tt.want {
				t.Errorf("IsTooOld() = %v, want %v", got, tt.want)
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
			name:   "returns the inner value",
			fields: fields{value: val},
			want:   val,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				value: tt.fields.value,
			}
			if got := v.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

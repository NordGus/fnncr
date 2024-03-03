package username

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		username string
	}
	type want struct {
		value Value
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "initializes a username value object",
			args: args{username: "john_wick"},
			want: want{
				value: Value{username: "john_wick"},
				err:   nil,
			},
		},
		{
			name: "can't initialize a username, is empty",
			args: args{username: ""},
			want: want{
				value: Value{username: ""},
				err:   ErrBlank,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.username)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("New() error = %v, want %v", err, tt.want.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.value) {
				t.Errorf("New() got = %v, want %v", got, tt.want.value)
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	type fields struct {
		username string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the inner username value",
			fields: fields{username: "john_wick"},
			want:   "john_wick",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				username: tt.fields.username,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

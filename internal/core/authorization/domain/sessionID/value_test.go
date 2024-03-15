package sessionID

import (
	"errors"
	"reflect"
	"testing"
)

type encoderMock struct {
	fail bool
}

func (mock encoderMock) EncodeToString(_ []byte) string {
	return "mock string"
}

func (mock encoderMock) DecodeString(s string) ([]byte, error) {
	if mock.fail {
		return []byte(s), errors.New("fail")
	}

	return make([]byte, ByteSize), nil
}

func TestNew(t *testing.T) {
	type args struct {
		id      [ByteSize]byte
		encoder Encoder
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{
			name: "initializes a new sessionID",
			args: args{
				id:      [ByteSize]byte{},
				encoder: encoderMock{fail: false},
			},
			want: Value{
				id:      [ByteSize]byte{},
				encoder: encoderMock{fail: false},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.encoder)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	type args struct {
		id      string
		encoder Encoder
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{
			name: "initializes a new sessionID",
			args: args{
				id:      "valid-id",
				encoder: encoderMock{fail: false},
			},
			want: Value{
				id:      [ByteSize]byte{},
				encoder: encoderMock{fail: false},
			},
			wantErr: false,
		},
		{
			name: "fails to initialize a new sessionID, because the string can't be parsed",
			args: args{
				id:      "invalid-id",
				encoder: encoderMock{fail: true},
			},
			want:    Value{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.args.id, tt.args.encoder)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	type fields struct {
		value   [ByteSize]byte
		encoder Encoder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the expected string representation for the given encoder",
			fields: fields{
				value:   [64]byte{},
				encoder: encoderMock{fail: false},
			},
			want: "mock string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				id:      tt.fields.value,
				encoder: tt.fields.encoder,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

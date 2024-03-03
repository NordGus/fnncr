package passworddigest

import (
	"errors"
	"reflect"
	"testing"
)

type cryptMock struct {
	compareErr bool
	costErr    bool
	cost       int
}

func (cm cryptMock) CompareHashAndPassword(_ []byte, _ []byte) error {
	if cm.compareErr {
		return errors.New("compare hash and password error")
	}

	return nil
}

func (cm cryptMock) Cost(hashedPassword []byte) (int, error) {
	if cm.costErr {
		return cm.cost, errors.New("cost error")
	}

	return cm.cost, nil
}

func TestNew(t *testing.T) {
	type args struct {
		hash  string
		crypt Crypt
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
			name: "returns a new password digest value object from a string hash",
			args: args{
				hash: "hash",
				crypt: cryptMock{
					cost: hashCost,
				},
			},
			want: want{
				value: Value{
					hash:  []byte("hash"),
					crypt: cryptMock{cost: hashCost},
				},
				err: nil,
			},
		},
		{
			name: "fails with invalid hash cost",
			args: args{
				hash:  "hash",
				crypt: cryptMock{cost: hashCost + 1},
			},
			want: want{
				value: Value{
					hash:  []byte("hash"),
					crypt: cryptMock{cost: hashCost + 1},
				},
				err: ErrHashCostInvalid,
			},
		},
		{
			name: "fails with invalid hash",
			args: args{
				hash:  "hash",
				crypt: cryptMock{costErr: true, cost: hashCost},
			},
			want: want{
				value: Value{
					hash:  []byte("hash"),
					crypt: cryptMock{costErr: true, cost: hashCost},
				},
				err: ErrHashInvalid,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.hash, tt.args.crypt)
			if !reflect.DeepEqual(got, tt.want.value) {
				t.Errorf("New() got = %v, want %v", got, tt.want.value)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("New() err = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestValue_Compare(t *testing.T) {
	type fields struct {
		hash  []byte
		crypt Crypt
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "password is valid",
			fields: fields{
				hash:  []byte("12345678"),
				crypt: cryptMock{cost: hashCost},
			},
			args: args{password: "12345678"},
		},
		{
			name: "password is invalid",
			fields: fields{
				hash:  []byte("12345678"),
				crypt: cryptMock{compareErr: true, cost: hashCost},
			},
			args: args{password: "12345678"},
			want: ErrInvalidPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				hash:  tt.fields.hash,
				crypt: tt.fields.crypt,
			}

			got := v.Compare(tt.args.password)
			if !errors.Is(got, tt.want) {
				t.Errorf("Compare() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	type fields struct {
		hash  []byte
		crypt Crypt
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the string representation of the inner hash",
			fields: fields{
				hash:  []byte("this is a test"),
				crypt: cryptMock{cost: hashCost},
			},
			want: "this is a test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				hash:  tt.fields.hash,
				crypt: tt.fields.crypt,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

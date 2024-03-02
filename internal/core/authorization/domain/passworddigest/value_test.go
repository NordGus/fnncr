package passworddigest

import (
	"errors"
	"reflect"
	"testing"
)

type cryptMock struct {
	compareErr  bool
	generateErr bool
	costErr     bool
	cost        int
}

func (cm cryptMock) CompareHashAndPassword(_ []byte, _ []byte) error {
	if cm.compareErr {
		return errors.New("compare hash and password error")
	}

	return nil
}

func (cm cryptMock) GenerateFromPassword(password []byte, _ int) ([]byte, error) {
	if cm.generateErr {
		return password, errors.New("failed to generate")
	}

	return password, nil
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
					hash:                 []byte("hash"),
					password:             "",
					passwordConfirmation: "",
					crypt:                cryptMock{cost: hashCost},
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
					hash:                 []byte("hash"),
					password:             "",
					passwordConfirmation: "",
					crypt:                cryptMock{cost: hashCost + 1},
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
					hash:                 []byte("hash"),
					password:             "",
					passwordConfirmation: "",
					crypt:                cryptMock{costErr: true, cost: hashCost},
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

func TestNewFromPassword(t *testing.T) {
	type args struct {
		password             string
		passwordConfirmation string
		crypt                Crypt
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
			name: "returns a new password digest value object from passwords",
			args: args{
				password:             "12345678",
				passwordConfirmation: "12345678",
				crypt:                cryptMock{cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte("12345678"),
					password:             "12345678",
					passwordConfirmation: "12345678",
					crypt:                cryptMock{cost: hashCost},
				},
				err: nil,
			},
		},
		{
			name: "fails with empty password",
			args: args{
				password:             "",
				passwordConfirmation: "",
				crypt:                cryptMock{cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte(""),
					password:             "",
					passwordConfirmation: "",
					crypt:                cryptMock{cost: hashCost},
				},
				err: ErrPasswordEmpty,
			},
		},
		{
			name: "fails with password too short",
			args: args{
				password:             "12345",
				passwordConfirmation: "12345",
				crypt:                cryptMock{cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte("12345"),
					password:             "12345",
					passwordConfirmation: "12345",
					crypt:                cryptMock{cost: hashCost},
				},
				err: ErrPasswordTooShort,
			},
		},
		{
			name: "fails with password too long",
			args: args{
				password:             "0123456789012345678901234567890123456789012345678901234567890123456789",
				passwordConfirmation: "0123456789012345678901234567890123456789012345678901234567890123456789",
				crypt:                cryptMock{cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte("0123456789012345678901234567890123456789012345678901234567890123456789"),
					password:             "0123456789012345678901234567890123456789012345678901234567890123456789",
					passwordConfirmation: "0123456789012345678901234567890123456789012345678901234567890123456789",
					crypt:                cryptMock{cost: hashCost},
				},
				err: ErrPasswordTooLong,
			},
		},
		{
			name: "fails with password doesn't match",
			args: args{
				password:             "012345678901234567890123456789",
				passwordConfirmation: "0123456789",
				crypt:                cryptMock{cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte("012345678901234567890123456789"),
					password:             "012345678901234567890123456789",
					passwordConfirmation: "0123456789",
					crypt:                cryptMock{cost: hashCost},
				},
				err: ErrPasswordDoesntMatch,
			},
		},
		{
			name: "fails with hash invalid",
			args: args{
				password:             "12345678",
				passwordConfirmation: "12345678",
				crypt:                cryptMock{generateErr: true, cost: hashCost},
			},
			want: want{
				value: Value{
					hash:                 []byte("12345678"),
					password:             "12345678",
					passwordConfirmation: "12345678",
					crypt:                cryptMock{generateErr: true, cost: hashCost},
				},
				err: ErrHashInvalid,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromPassword(tt.args.password, tt.args.passwordConfirmation, tt.args.crypt)
			if !reflect.DeepEqual(got, tt.want.value) {
				t.Errorf("NewFromPassword() got = %v, want %v", got, tt.want.value)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("NewFromPassword() err = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestValue_Compare(t *testing.T) {
	type fields struct {
		hash                 []byte
		password             string
		passwordConfirmation string
		crypt                Crypt
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
				hash:                 []byte("12345678"),
				password:             "",
				passwordConfirmation: "",
				crypt:                cryptMock{cost: hashCost},
			},
			args: args{password: "12345678"},
		},
		{
			name: "password is invalid",
			fields: fields{
				hash:                 []byte("12345678"),
				password:             "",
				passwordConfirmation: "",
				crypt:                cryptMock{compareErr: true, cost: hashCost},
			},
			args: args{password: "12345678"},
			want: ErrInvalidPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				hash:                 tt.fields.hash,
				password:             tt.fields.password,
				passwordConfirmation: tt.fields.passwordConfirmation,
				crypt:                tt.fields.crypt,
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
		hash                 []byte
		password             string
		passwordConfirmation string
		crypt                Crypt
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the string representation of the inner hash",
			fields: fields{
				hash:                 []byte("this is a test"),
				password:             "",
				passwordConfirmation: "",
				crypt:                cryptMock{cost: hashCost},
			},
			want: "this is a test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				hash:                 tt.fields.hash,
				password:             tt.fields.password,
				passwordConfirmation: tt.fields.passwordConfirmation,
				crypt:                tt.fields.crypt,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

package userID

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

type encoderMock struct {
	validateErr bool
	parseErr    bool
}

func (mock encoderMock) Validate(_ string) error {
	if mock.validateErr {
		return errors.New("failed to validate")
	}

	return nil
}

func (mock encoderMock) Parse(s string) (uuid.UUID, error) {
	if mock.parseErr {
		return uuid.UUID{}, errors.New("failed to parse")
	}

	return uuid.Parse(s)
}

func TestNew(t *testing.T) {
	type args struct {
		id      string
		encoder Encoder
	}
	type want struct {
		value Value
		err   error
	}

	i := uuid.New()

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "initializes a userID value object",
			args: args{
				id:      i.String(),
				encoder: encoderMock{},
			},
			want: want{
				value: Value{
					value:   i,
					encoder: encoderMock{},
				},
				err: nil,
			},
		},
		{
			name: "failed to initialize userID, id is invalid",
			args: args{
				id:      i.String(),
				encoder: encoderMock{validateErr: true},
			},
			want: want{
				value: Value{
					value:   i,
					encoder: encoderMock{validateErr: true},
				},
				err: ErrInvalid,
			},
		},
		{
			name: "failed to initialize userID, id can't be parsed",
			args: args{
				id:      i.String(),
				encoder: encoderMock{parseErr: true},
			},
			want: want{
				value: Value{
					value:   uuid.UUID{},
					encoder: encoderMock{parseErr: true},
				},
				err: ErrFailedToParsed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.encoder)
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
		value   uuid.UUID
		encoder Encoder
	}

	i := uuid.New()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the expected string representation of the inner uuid",
			fields: fields{
				value:   i,
				encoder: encoderMock{},
			},
			want: i.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				value:   tt.fields.value,
				encoder: tt.fields.encoder,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

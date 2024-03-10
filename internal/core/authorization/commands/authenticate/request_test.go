package authenticate

import (
	"context"
	"errors"
	"financo/internal/core/authorization/domain/sessionID"
	"reflect"
	"testing"
)

var decodeErr = errors.New("decode err")

type sessionIDEncoderMock struct {
	decodeErr bool
}

func (mock sessionIDEncoderMock) EncodeToString(src []byte) string {
	return string(src)
}

func (mock sessionIDEncoderMock) DecodeString(s string) ([]byte, error) {
	out := make([]byte, sessionID.ByteSize)

	copy(out, s)

	if mock.decodeErr {
		return out, decodeErr
	}

	return out, nil
}

func TestNewRequest(t *testing.T) {
	type args struct {
		ctx       context.Context
		sessionId string
		encoder   sessionID.Encoder
	}

	ctx := context.TODO()
	idValue := make([]byte, sessionID.ByteSize)
	copy(idValue, "this is a test")
	val, _ := sessionID.NewFromString("this is a test", sessionIDEncoderMock{})

	tests := []struct {
		name string
		args args
		want Request
	}{
		{
			name: "initializes the Request DTO",
			args: args{
				ctx:       ctx,
				sessionId: "this is a test",
				encoder:   sessionIDEncoderMock{},
			},
			want: Request{
				ctx:       ctx,
				sessionID: val,
				err:       nil,
			},
		},
		{
			name: "fails to initializes the Request DTO, failed to decode value",
			args: args{
				ctx:       ctx,
				sessionId: "this is a test",
				encoder:   sessionIDEncoderMock{decodeErr: true},
			},
			want: Request{
				ctx:       ctx,
				sessionID: sessionID.Value{},
				err:       errors.Join(sessionID.ErrCantBeDecodedFromString, decodeErr),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(tt.args.ctx, tt.args.sessionId, tt.args.encoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

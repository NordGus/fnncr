package authenticate

import (
	"errors"
	"financo/internal/core_old/authorization/domain/passworddigest"
	"financo/internal/core_old/authorization/domain/sessionversion"
	"financo/internal/core_old/authorization/domain/timestamp"
	"financo/internal/core_old/authorization/domain/user"
	"financo/internal/core_old/authorization/domain/userID"
	"financo/internal/core_old/authorization/domain/username"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

type passwordDigestCryptMock struct{}

func (passwordDigestCryptMock) CompareHashAndPassword(_ []byte, _ []byte) error {
	return nil
}

func (passwordDigestCryptMock) Cost(_ []byte) (int, error) {
	return 10, nil
}

type userIDEncoderMock struct{}

func (userIDEncoderMock) Validate(s string) error {
	return uuid.Validate(s)
}

func (userIDEncoderMock) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func userMock() user.Entity {
	id, _ := userID.New(uuid.NewString(), userIDEncoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", passwordDigestCryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	return user.New(id, un, pw, sv, ct, ut)
}

func TestResponse_Error(t *testing.T) {
	type fields struct {
		user user.Entity
		err  error
	}

	u := userMock()
	err := errors.New("an error")

	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "returns the expected wrap error, nil",
			fields: fields{
				user: u,
				err:  nil,
			},
			want: nil,
		},
		{
			name: "returns the expected wrap error, an error",
			fields: fields{
				user: u,
				err:  err,
			},
			want: err,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				user: tt.fields.user,
				err:  tt.fields.err,
			}
			if got := resp.Error(); !errors.Is(tt.want, got) {
				t.Errorf("Error() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_User(t *testing.T) {
	type fields struct {
		user user.Entity
		err  error
	}

	u := userMock()

	tests := []struct {
		name   string
		fields fields
		want   user.Entity
	}{
		{
			name: "returns the wrapped user entity",
			fields: fields{
				user: u,
				err:  nil,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				user: tt.fields.user,
				err:  tt.fields.err,
			}
			if got := resp.User(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User() = %v, want %v", got, tt.want)
			}
		})
	}
}

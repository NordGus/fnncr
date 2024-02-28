package session

import (
	"reflect"
	"testing"
	"time"

	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/user/passworddigest"
	"financo/internal/core/authorization/domain/user/username"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type cryptMock struct{}

func (cryptMock) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (cryptMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (cryptMock) Cost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}

func userMock(sessionVersion uint32) user.Entity {
	uid := uuid.New()
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.NewFromPassword("12345678", "12345678", cryptMock{})
	sv, _ := sessionversion.New(sessionVersion)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	return user.New(uid, un, pw, sv, ct, ut)
}

func TestEntity_Expired(t *testing.T) {
	type fields struct {
		id        sessionID.Value
		version   sessionversion.Value
		createdAt timestamp.Value
		userID    uuid.UUID
	}
	type args struct {
		user   user.Entity
		maxAge time.Duration
	}

	uid := uuid.New()
	i, _ := sessionID.New([sessionID.ByteSize]byte{}, sessionID.DefaultEncoder)
	ver, _ := sessionversion.New(42)
	createdAt, _ := timestamp.New(time.Now().Add(-7 * 24 * time.Hour))

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "is a valid session",
			fields: fields{
				id:        i,
				version:   ver,
				createdAt: createdAt,
				userID:    uid,
			},
			args: args{
				user:   userMock(42),
				maxAge: 30 * 24 * time.Hour,
			},
			want: false,
		},
		{
			name: "is an invalid session, because is too old",
			fields: fields{
				id:        i,
				version:   ver,
				createdAt: createdAt,
				userID:    uid,
			},
			args: args{
				user:   userMock(42),
				maxAge: 24 * time.Hour,
			},
			want: true,
		},
		{
			name: "is an invalid session, because the session is invalid",
			fields: fields{
				id:        i,
				version:   ver,
				createdAt: createdAt,
				userID:    uid,
			},
			args: args{
				user:   userMock(7),
				maxAge: 30 * 24 * time.Hour,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:        tt.fields.id,
				userID:    tt.fields.userID,
				version:   tt.fields.version,
				createdAt: tt.fields.createdAt,
			}

			if got := e.Expired(tt.args.user, tt.args.maxAge); got != tt.want {
				t.Errorf("Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		id        sessionID.Value
		version   sessionversion.Value
		createdAt timestamp.Value
		userID    uuid.UUID
	}

	uid := uuid.New()
	i, _ := sessionID.New([sessionID.ByteSize]byte{1}, sessionID.DefaultEncoder)
	ver, _ := sessionversion.New(42)
	createdAt, _ := timestamp.New(time.Now())

	tests := []struct {
		name string
		args args
		want Entity
	}{
		{
			name: "initializes a new session entity",
			args: args{
				id:        i,
				version:   ver,
				createdAt: createdAt,
				userID:    uid,
			},
			want: Entity{
				id:        i,
				userID:    uid,
				version:   ver,
				createdAt: createdAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.id, tt.args.version, tt.args.createdAt, tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

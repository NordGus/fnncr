package session

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"financo/internal/core/authorization/domain/session/creationtime"
	"financo/internal/core/authorization/domain/session/id"
	"financo/internal/core/authorization/domain/session/version"
	"github.com/google/uuid"
)

type userEntityMock struct {
	sessionVersion uint32
}

func (mock userEntityMock) CurrentSessionVersion() uint32 {
	return mock.sessionVersion
}

func newUserEntityMock(version uint32) userEntityMock {
	return userEntityMock{
		sessionVersion: version,
	}
}

func TestEntity_Expired(t *testing.T) {
	type fields struct {
		id        id.Value
		version   version.Value
		createdAt creationtime.Value
		userID    uuid.UUID
	}
	type args struct {
		user   UserEntity
		maxAge time.Duration
	}

	uid := uuid.New()
	i, _ := id.New([id.ByteSize]byte{}, base64.URLEncoding)
	ver, _ := version.New(42)
	createdAt, _ := creationtime.New(time.Now().Add(-7 * 24 * time.Hour))

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
				user:   newUserEntityMock(42),
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
				user:   newUserEntityMock(42),
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
				user:   newUserEntityMock(7),
				maxAge: 30 * 24 * time.Hour,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				ID:        tt.fields.id,
				UserID:    tt.fields.userID,
				Version:   tt.fields.version,
				CreatedAt: tt.fields.createdAt,
			}

			if got := e.Expired(tt.args.user, tt.args.maxAge); got != tt.want {
				t.Errorf("Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		id        id.Value
		version   version.Value
		createdAt creationtime.Value
		userID    uuid.UUID
	}

	uid := uuid.New()
	i, _ := id.New([id.ByteSize]byte{1}, base64.URLEncoding)
	ver, _ := version.New(42)
	createdAt, _ := creationtime.New(time.Now())

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
				ID:        i,
				UserID:    uid,
				Version:   ver,
				CreatedAt: createdAt,
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

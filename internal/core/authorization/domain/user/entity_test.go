package user

import (
	"reflect"
	"testing"
	"time"

	"financo/internal/core/authorization/domain/passworddigest"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
	"github.com/google/uuid"
)

type cryptMock struct{}

func (cryptMock) CompareHashAndPassword(_ []byte, _ []byte) error {
	return nil
}

func (cryptMock) Cost(_ []byte) (int, error) {
	return 10, nil
}

type encoderMock struct{}

func (encoderMock) Validate(s string) error {
	return uuid.Validate(s)
}

func (encoderMock) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func TestEntity_CreatedAt(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   timestamp.Value
	}{
		{
			name: "returns user's created at timestamp",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: ct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_ID(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   userID.Value
	}{
		{
			name: "returns user's id",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: id,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_PasswordDigest(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   passworddigest.Value
	}{
		{
			name: "returns user's password digest",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: pw,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.PasswordDigest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PasswordDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_SessionVersion(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   sessionversion.Value
	}{
		{
			name: "returns user's password digest",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: sv,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.SessionVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpdatedAt(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   timestamp.Value
	}{
		{
			name: "returns user's updated at timestamp",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: ut,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_Username(t *testing.T) {
	type fields struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createAt       timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name   string
		fields fields
		want   username.Value
	}{
		{
			name: "returns user's updated at timestamp",
			fields: fields{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createAt:       ct,
				updatedAt:      ut,
			},
			want: un,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:             tt.fields.id,
				username:       tt.fields.username,
				passwordDigest: tt.fields.passwordDigest,
				sessionVersion: tt.fields.sessionVersion,
				createdAt:      tt.fields.createAt,
				updatedAt:      tt.fields.updatedAt,
			}
			if got := e.Username(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Username() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		id             userID.Value
		username       username.Value
		passwordDigest passworddigest.Value
		sessionVersion sessionversion.Value
		createdAt      timestamp.Value
		updatedAt      timestamp.Value
	}

	id, _ := userID.New(uuid.NewString(), encoderMock{})
	un, _ := username.New("john_wick")
	pw, _ := passworddigest.New("hash", cryptMock{})
	sv, _ := sessionversion.New(42)
	ct, _ := timestamp.New(time.Now())
	ut, _ := timestamp.New(time.Now())

	tests := []struct {
		name string
		args args
		want Entity
	}{
		{
			name: "returns user's updated at timestamp",
			args: args{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createdAt:      ct,
				updatedAt:      ut,
			},
			want: Entity{
				id:             id,
				username:       un,
				passwordDigest: pw,
				sessionVersion: sv,
				createdAt:      ct,
				updatedAt:      ut,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.id, tt.args.username, tt.args.passwordDigest, tt.args.sessionVersion, tt.args.createdAt, tt.args.updatedAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

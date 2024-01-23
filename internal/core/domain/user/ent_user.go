package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Username       Username
	PasswordDigest PasswordDigest
	Password       Password
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func New(id uuid.UUID, un Username, pw PasswordDigest, ct time.Time, ut time.Time) User {
	return User{
		ID:             id,
		Username:       un,
		PasswordDigest: pw,
		CreatedAt:      ct,
		UpdatedAt:      ut,
	}
}

func NewWithPassword(id uuid.UUID, un Username, pw Password, ct time.Time, ut time.Time) User {
	return User{
		ID:        id,
		Username:  un,
		Password:  pw,
		CreatedAt: ct,
		UpdatedAt: ut,
	}
}

package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Username       Username
	PasswordDigest PasswordDigest
	SessionVersion int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func New(id uuid.UUID, un Username, pw PasswordDigest, sv int32, ct time.Time, ut time.Time) User {
	return User{
		ID:             id,
		Username:       un,
		PasswordDigest: pw,
		SessionVersion: sv,
		CreatedAt:      ct,
		UpdatedAt:      ut,
	}
}

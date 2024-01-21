package user

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID
	Username       Username
	PasswordDigest PasswordDigest
}

func New(id uuid.UUID, un Username, pw PasswordDigest) User {
	return User{
		ID:             id,
		Username:       un,
		PasswordDigest: pw,
	}
}

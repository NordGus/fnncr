package user

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID
	Username       Username
	PasswordDigest PasswordDigest
	Password       Password
}

func New(id uuid.UUID, un Username, pw PasswordDigest) User {
	return User{
		ID:             id,
		Username:       un,
		PasswordDigest: pw,
	}
}

func NewWithPassword(id uuid.UUID, un Username, pw Password) User {
	return User{
		ID:       id,
		Username: un,
		Password: pw,
	}
}

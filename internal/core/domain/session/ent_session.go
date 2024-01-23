package session

import "github.com/google/uuid"

type Session struct {
	ID     ID
	UserID uuid.UUID
}

func New(id ID, userId uuid.UUID) Session {
	return Session{
		ID:     id,
		UserID: userId,
	}
}

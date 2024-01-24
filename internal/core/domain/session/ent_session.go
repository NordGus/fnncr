package session

import "github.com/google/uuid"

type Session struct {
	ID      ID
	UserID  uuid.UUID
	Version int32
}

func New(id ID, userId uuid.UUID, version int32) Session {
	return Session{
		ID:      id,
		UserID:  userId,
		Version: version,
	}
}

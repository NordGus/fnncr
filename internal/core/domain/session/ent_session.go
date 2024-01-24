package session

import (
	"errors"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/google/uuid"
)

var (
	ErrExpired = errors.New("session: session expired")
)

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

func (s Session) Expired(usr user.User) bool {
	// TODO: implement a session lifespan
	return s.Version != usr.SessionVersion
}

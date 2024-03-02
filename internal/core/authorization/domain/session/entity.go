package session

import (
	"time"

	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
)

type Entity struct {
	id        sessionID.Value
	userID    userID.Value
	version   sessionversion.Value
	createdAt timestamp.Value
}

func New(id sessionID.Value, version sessionversion.Value, createdAt timestamp.Value, userID userID.Value) Entity {
	return Entity{
		id:        id,
		version:   version,
		createdAt: createdAt,
		userID:    userID,
	}
}

func (e *Entity) ID() sessionID.Value {
	return e.id
}

func (e *Entity) Version() sessionversion.Value {
	return e.version
}

func (e *Entity) CreatedAt() timestamp.Value {
	return e.createdAt
}

func (e *Entity) UserID() userID.Value {
	return e.userID
}

func (e *Entity) Expired(user user.Entity, maxAge time.Duration) bool {
	return e.version.IsInvalid(user.SessionVersion()) || time.Since(e.createdAt.Time()) > maxAge
}

func (e *Entity) IsTooOld(maxAge time.Duration) bool {
	return time.Since(e.createdAt.Time()) > maxAge
}

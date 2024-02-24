package session

import (
	"time"

	"financo/internal/core/authorization/domain/session/creationtime"
	"financo/internal/core/authorization/domain/session/id"
	"financo/internal/core/authorization/domain/session/version"
	"github.com/google/uuid"
)

type UserEntity interface {
	CurrentSessionVersion() uint32
}

type Entity struct {
	ID        id.Value
	UserID    uuid.UUID
	Version   version.Value
	CreatedAt creationtime.Value
}

func New(id id.Value, version version.Value, createdAt creationtime.Value, userID uuid.UUID) Entity {
	return Entity{
		ID:        id,
		Version:   version,
		CreatedAt: createdAt,
		UserID:    userID,
	}
}

func (e *Entity) Expired(user UserEntity, maxAge time.Duration) bool {
	return e.Version.IsInvalid(user.CurrentSessionVersion()) || e.CreatedAt.IsTooOld(maxAge)
}

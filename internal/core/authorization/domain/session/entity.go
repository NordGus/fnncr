package session

import (
	"time"

	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user"
	"github.com/google/uuid"
)

type Entity struct {
	id        sessionID.Value
	userID    uuid.UUID // TODO: migrate to a value object.
	version   sessionversion.Value
	createdAt timestamp.Value
}

func New(id sessionID.Value, version sessionversion.Value, createdAt timestamp.Value, userID uuid.UUID) Entity {
	return Entity{
		id:        id,
		version:   version,
		createdAt: createdAt,
		userID:    userID,
	}
}

func (e *Entity) Expired(user user.Entity, maxAge time.Duration) bool {
	return e.version.IsInvalid(user.CurrentSessionVersion()) || time.Since(e.createdAt.Time()) > maxAge
}

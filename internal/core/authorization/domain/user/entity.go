package user

import (
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user/passworddigest"
	"financo/internal/core/authorization/domain/user/username"
	"github.com/google/uuid"
)

type Entity struct {
	id             uuid.UUID // TODO: migrate to a value object.
	username       username.Value
	passwordDigest passworddigest.Value
	sessionVersion sessionversion.Value
	createAt       timestamp.Value
	updatedAt      timestamp.Value
}

func New(
	id uuid.UUID,
	username username.Value,
	passwordDigest passworddigest.Value,
	sessionVersion sessionversion.Value,
	createdAt timestamp.Value,
	updatedAt timestamp.Value,
) Entity {
	return Entity{
		id:             id,
		username:       username,
		passwordDigest: passwordDigest,
		sessionVersion: sessionVersion,
		createAt:       createdAt,
		updatedAt:      updatedAt,
	}
}

func (e *Entity) CurrentSessionVersion() sessionversion.Value {
	return e.sessionVersion
}

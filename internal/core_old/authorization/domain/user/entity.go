package user

import (
	"financo/internal/core_old/authorization/domain/passworddigest"
	"financo/internal/core_old/authorization/domain/sessionversion"
	"financo/internal/core_old/authorization/domain/timestamp"
	"financo/internal/core_old/authorization/domain/userID"
	"financo/internal/core_old/authorization/domain/username"
)

type Entity struct {
	id             userID.Value
	username       username.Value
	passwordDigest passworddigest.Value
	sessionVersion sessionversion.Value
	createdAt      timestamp.Value
	updatedAt      timestamp.Value
}

func New(
	id userID.Value,
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
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func (e *Entity) ID() userID.Value {
	return e.id
}

func (e *Entity) Username() username.Value {
	return e.username
}

func (e *Entity) PasswordDigest() passworddigest.Value {
	return e.passwordDigest
}

func (e *Entity) SessionVersion() sessionversion.Value {
	return e.sessionVersion
}

func (e *Entity) CreatedAt() timestamp.Value {
	return e.createdAt
}

func (e *Entity) UpdatedAt() timestamp.Value {
	return e.updatedAt
}

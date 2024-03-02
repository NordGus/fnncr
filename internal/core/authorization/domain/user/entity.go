package user

import (
	"financo/internal/core/authorization/domain/passworddigest"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
)

type Entity struct {
	id             userID.Value
	username       username.Value
	passwordDigest passworddigest.Value
	sessionVersion sessionversion.Value
	createAt       timestamp.Value
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
		createAt:       createdAt,
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
	return e.createAt
}

func (e *Entity) UpdatedAt() timestamp.Value {
	return e.updatedAt
}

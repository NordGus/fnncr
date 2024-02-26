package user

import (
	"financo/internal/core/authorization/domain/user/creationtime"
	"financo/internal/core/authorization/domain/user/passworddigest"
	"financo/internal/core/authorization/domain/user/sessionversion"
	"financo/internal/core/authorization/domain/user/updatetime"
	"financo/internal/core/authorization/domain/user/username"
	"github.com/google/uuid"
)

type Entity struct {
	ID             uuid.UUID
	Username       username.Value
	PasswordDigest passworddigest.Value
	SessionVersion sessionversion.Value
	CreateAt       creationtime.Value
	UpdatedAt      updatetime.Value
}

func New(
	id uuid.UUID,
	un username.Value,
	pw passworddigest.Value,
	sv sessionversion.Value,
	ct creationtime.Value,
	ut updatetime.Value,
) Entity {
	return Entity{
		ID:             id,
		Username:       un,
		PasswordDigest: pw,
		SessionVersion: sv,
		CreateAt:       ct,
		UpdatedAt:      ut,
	}
}

func (e *Entity) CurrentSessionVersion() uint32 {
	return e.SessionVersion.Uint32()
}

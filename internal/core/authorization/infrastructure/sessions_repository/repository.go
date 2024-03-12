package sessions_repository

import (
	"context"

	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
)

type Repository interface {
	Get(ctx context.Context, id sessionID.Value) (session.Entity, error)
	Create(ctx context.Context, entity session.Entity) error
	Delete(ctx context.Context, id sessionID.Value) error
}

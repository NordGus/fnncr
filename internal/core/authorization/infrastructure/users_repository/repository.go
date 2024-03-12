package users_repository

import (
	"context"

	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
)

type Repository interface {
	GetByID(ctx context.Context, id userID.Value) (user.Entity, error)
	GetByUsername(ctx context.Context, value username.Value) (user.Entity, error)
	Save(ctx context.Context, entity user.Entity) error
}

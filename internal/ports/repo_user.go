package ports

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user does not exists")
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username user.Username) (user.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (user.User, error)
}

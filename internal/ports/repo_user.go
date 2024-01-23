package ports

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user does not exists")
	ErrUserNotSaved = errors.New("user couldn't be saved")
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username user.Username) (user.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (user.User, error)
	Create(ctx context.Context, entity user.User) error
}

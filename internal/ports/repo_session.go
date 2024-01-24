package ports

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/internal/core/domain/session"
)

var (
	ErrSessionNotCreated = errors.New("session could not be created")
	ErrSessionNotFound   = errors.New("session does not exists")
)

type SessionRepository interface {
	Get(ctx context.Context, id session.ID) (session.Session, error)
	Create(ctx context.Context, session session.Session) error
	Delete(ctx context.Context, session session.Session) error
}

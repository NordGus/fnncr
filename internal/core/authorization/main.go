package authorization

import (
	"context"
	"time"

	"financo/internal/core/authorization/commands/authenticate"
	"financo/internal/core/authorization/commands/signin"
	"financo/internal/core/authorization/commands/signout"

	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
)

var (
	instance API
)

type (
	UserRepository interface {
		GetByID(ctx context.Context, id userID.Value) (user.Entity, error)
		GetByUsername(ctx context.Context, value username.Value) (user.Entity, error)
		Save(ctx context.Context, entity user.Entity) error
	}

	SessionRepository interface {
		Get(ctx context.Context, id sessionID.Value) (session.Entity, error)
		Create(ctx context.Context, entity session.Entity) error
		Delete(ctx context.Context, id sessionID.Value) error
	}

	// Opts is a configuration struct for the authentication service.
	Opts struct {
		// SessionMaxAge is the maximum age for the current user session. Default 1 hour.
		SessionMaxAge time.Duration
		// SessionStaleAge represents the maximum age for a session cached in the system. Default 24 hours.
		SessionStaleAge time.Duration
	}

	// OptFunc is a configuration function that lets you configure the service's options
	OptFunc func(opts *Opts)
)

func defaultOptions() *Opts {
	return &Opts{
		SessionMaxAge:   time.Hour,
		SessionStaleAge: 24 * time.Hour,
	}
}

func New(userRepo UserRepository, sessionRepo SessionRepository, configs ...OptFunc) API {
	if instance != nil {
		return instance
	}

	opts := defaultOptions()

	for i := 0; i < len(configs); i++ {
		configs[i](opts)
	}

	instance = newService(
		signin.New(userRepo, sessionRepo),
		signout.New(userRepo),
		authenticate.New(userRepo, sessionRepo, opts.SessionMaxAge, opts.SessionStaleAge),
	)

	return instance
}

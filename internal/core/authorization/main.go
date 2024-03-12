package authorization

import (
	"time"

	"financo/internal/core/authorization/commands/authenticate"
	"financo/internal/core/authorization/commands/signin"
	"financo/internal/core/authorization/commands/signout"
	"financo/internal/core/authorization/infrastructure/bcrypt_crypt"
	"financo/internal/core/authorization/infrastructure/sessions_repository"
	"financo/internal/core/authorization/infrastructure/users_repository"
	"financo/internal/core/authorization/infrastructure/uuid_encoder"
)

var (
	instance API
)

type (
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

func New(pgService users_repository.PostgresService, redisService sessions_repository.RedisService, configs ...OptFunc) API {
	if instance != nil {
		return instance
	}

	var (
		userIDEncoder       = uuid_encoder.New()
		passwordDigestCrypt = bcrypt_crypt.New()
		userRepo            = users_repository.NewPostgreSQLRepository(pgService, userIDEncoder, passwordDigestCrypt)
		sessionRepo         = sessions_repository.NewRedisRepository(redisService, userIDEncoder)
		opts                = defaultOptions()
	)

	for i := 0; i < len(configs); i++ {
		configs[i](opts)
	}

	instance = newService(
		signin.New(userRepo, sessionRepo),
		signout.New(userRepo),
		authenticate.New(userRepo, sessionRepo, opts.SessionMaxAge, opts.SessionStaleAge),
		userIDEncoder,
		passwordDigestCrypt,
	)

	return instance
}

package authentication

import (
	"context"
	"time"
)

type (
	UserRecord interface {
		Id() int64
		PasswordHash() []byte
	}

	UserRepository interface {
		GetByUsername(ctx context.Context, username string) (UserRecord, error)
		GetByID(ctx context.Context, id int64) (UserRecord, error)
	}

	SessionRecord interface {
		UserId() int64
	}

	SessionStore interface {
		Create(ctx context.Context, sessionID string, userID int64) error
		Get(ctx context.Context, sessionID string) (SessionRecord, error)
	}

	Service struct {
		sessionCookieName     string
		sessionCookieDuration time.Duration
		sessionCookieDomain   string
		userRepository        UserRepository
		sessionRepository     SessionStore
	}

	Opts struct {
		SessionCookieName     string
		SessionCookieDuration time.Duration
		SessionCookieDomain   string
		UserRepository        UserRepository
		SessionRepository     SessionStore
	}

	ConfigFunc func(opts *Opts)
)

var (
	defaults = Opts{
		SessionCookieName:     "_fnncr_session",
		SessionCookieDuration: 24 * time.Hour, // One Day
		SessionCookieDomain:   "localhost",
	}
)

func New(configs ...ConfigFunc) *Service {
	var (
		opts = defaults
	)

	for i := 0; i < len(configs); i++ {
		configs[i](&opts)
	}

	return &Service{
		sessionCookieName:     opts.SessionCookieName,
		sessionCookieDuration: opts.SessionCookieDuration,
		sessionCookieDomain:   opts.SessionCookieDomain,
		userRepository:        opts.UserRepository,
		sessionRepository:     opts.SessionRepository,
	}
}

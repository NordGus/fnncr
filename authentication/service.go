package authentication

import (
	"time"
)

type (
	Service struct {
		sessionCookieName     string
		sessionCookieDuration time.Duration
		sessionCookieDomain   string
	}

	Opts struct {
		sessionCookieName     string
		sessionCookieDuration time.Duration
		sessionCookieDomain   string
	}
)

var (
	defaults = Opts{
		sessionCookieName:     "_fnncr_session",
		sessionCookieDuration: 24 * time.Hour, // One Day
		sessionCookieDomain:   "localhost",
	}
)

func New() *Service {
	var (
		opts = defaults
	)

	// TODO: Implement a mechanism to configure the authentication service

	return &Service{
		sessionCookieName:     opts.sessionCookieName,
		sessionCookieDuration: opts.sessionCookieDuration,
		sessionCookieDomain:   opts.sessionCookieDomain,
	}
}

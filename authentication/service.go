package authentication

import "time"

type Service struct {
	sessionCookieName     string
	sessionCookieDuration time.Duration
	sessionCookieDomain   string
}

func New() *Service {
	return &Service{
		sessionCookieName:     "_fnncr_session",
		sessionCookieDuration: 24 * time.Hour, // One Day
		sessionCookieDomain:   "localhost",
	}
}

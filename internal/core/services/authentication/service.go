package authentication

import (
	"context"
	"time"

	"financo/internal/ports"
)

type (
	API interface {
		SignIn(ctx context.Context, req SignInUserReq) (SignInUserResp, error)
		SignOut(ctx context.Context, req SignOutUserReq) (SignOutUserResp, error)
		Authenticate(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error)
	}

	service struct {
		sessionMaxAge time.Duration
		sessionRepo   ports.SessionRepository
		userRepo      ports.UserRepository
	}
)

func NewService(sessionMaxAge time.Duration, sessionRepo ports.SessionRepository, userRepo ports.UserRepository) API {
	return &service{
		sessionMaxAge: sessionMaxAge,
		sessionRepo:   sessionRepo,
		userRepo:      userRepo,
	}
}

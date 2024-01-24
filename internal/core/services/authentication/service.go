package authentication

import (
	"context"
	"time"

	"github.com/NordGus/fnncr/internal/ports"
)

type (
	API interface {
		SignInUser(ctx context.Context, req SignInUserReq) (SignInUserResp, error)
		SignOutUser(ctx context.Context, req SignOutUserReq) (SignOutUserResp, error)
		AuthenticateSession(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error)
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

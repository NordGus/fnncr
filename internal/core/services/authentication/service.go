package authentication

import (
	"context"

	"github.com/NordGus/fnncr/internal/ports"
)

type (
	API interface {
		SignInUser(ctx context.Context, req SignInUserReq) (SignInUserResp, error)
		AuthenticateSession(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error)
	}

	Service struct {
		sessionRepo ports.SessionRepository
		userRepo    ports.UserRepository
	}
)

func NewService(sessionRepo ports.SessionRepository, userRepo ports.UserRepository) API {
	return &Service{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

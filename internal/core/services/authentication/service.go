package authentication

import (
	"context"

	"github.com/NordGus/fnncr/internal/ports"
)

type (
	API interface {
		SignInUser(ctx context.Context, req SignInUserReq) (SignInUserResp, error)
		SignOutUser(ctx context.Context, req SignOutUserReq) (SignOutUserResp, error)
		AuthenticateSession(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error)
	}

	service struct {
		sessionRepo ports.SessionRepository
		userRepo    ports.UserRepository
	}
)

func NewService(sessionRepo ports.SessionRepository, userRepo ports.UserRepository) API {
	return &service{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

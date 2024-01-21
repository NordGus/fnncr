package authentication

import (
	"context"

	"github.com/NordGus/fnncr/internal/ports"
)

type (
	API interface {
		LoginUser(ctx context.Context, req LoginUserReq) (LoginUserResp, error)
		AuthenticateUser(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error)
	}

	Service struct {
		sessionRepo ports.SessionRepository
		userRepo    ports.UserRepository
	}
)

func NewService(sessionRepo ports.SessionRepository, userRepo ports.UserRepository) *Service {
	return &Service{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

package users

import (
	"context"

	"github.com/NordGus/fnncr/internal/ports"
)

type (
	API interface {
		Create(ctx context.Context, req CreateReq) (CreateResp, error)
	}

	service struct {
		userRepo ports.UserRepository
	}
)

func NewService(userRepo ports.UserRepository) API {
	return &service{
		userRepo: userRepo,
	}
}

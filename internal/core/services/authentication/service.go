package authentication

import "github.com/NordGus/fnncr/internal/ports"

type Service struct {
	sessionRepo ports.SessionRepository
	userRepo    ports.UserRepository
}

func NewService(sessionRepo ports.SessionRepository, userRepo ports.UserRepository) *Service {
	return &Service{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

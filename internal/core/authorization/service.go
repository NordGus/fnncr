package authorization

import (
	"context"
	"time"

	"financo/internal/core/authorization/commands/authenticate"
	"financo/internal/core/authorization/commands/signin"
	"financo/internal/core/authorization/commands/signout"
	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
)

const (
	SessionMaxAge   = time.Hour
	SessionStaleAge = 24 * time.Hour
)

type (
	API interface {
		SignIn(ctx context.Context, username string, password string) SignInResponse
		SignOut(ctx context.Context, sessionID string) SignOutResponse
		AuthenticateUser(ctx context.Context, sessionID string) AuthenticateUserResponse
	}

	UserRepository interface {
		GetByID(ctx context.Context, id userID.Value) (user.Entity, error)
		GetByUsername(ctx context.Context, value username.Value) (user.Entity, error)
		Save(ctx context.Context, entity user.Entity) error
	}

	SessionRepository interface {
		Get(ctx context.Context, id sessionID.Value) (session.Entity, error)
		Create(ctx context.Context, entity session.Entity) error
		Delete(ctx context.Context, id sessionID.Value) error
	}

	service struct {
		signInCommand       signin.Command
		signOutCommand      signout.Command
		authenticateCommand authenticate.Command
	}
)

func New(userRepository UserRepository, sessionRepository SessionRepository) API {
	return &service{
		signInCommand:       signin.New(userRepository, sessionRepository),
		signOutCommand:      signout.New(userRepository),
		authenticateCommand: authenticate.New(userRepository, sessionRepository, SessionMaxAge, SessionStaleAge),
	}
}

func (s *service) SignIn(ctx context.Context, username string, password string) SignInResponse {
	return SignInResponse{}
}

func (s *service) SignOut(ctx context.Context, sessionID string) SignOutResponse {
	return SignOutResponse{}
}

func (s *service) AuthenticateUser(ctx context.Context, sessionID string) AuthenticateUserResponse {
	return AuthenticateUserResponse{}
}

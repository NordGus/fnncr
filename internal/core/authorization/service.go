package authorization

import (
	"context"

	"financo/internal/core/authorization/commands/authenticate"
	"financo/internal/core/authorization/commands/signin"
	"financo/internal/core/authorization/commands/signout"
	"github.com/google/uuid"
)

type (
	API interface {
		SignIn(ctx context.Context, username string, password string) SignInResponse
		SignOut(ctx context.Context, sessionID string) SignOutResponse
		AuthenticateUser(ctx context.Context, sessionID string) AuthenticateUserResponse
	}

	service struct {
		signInCommand       signin.Command
		signOutCommand      signout.Command
		authenticateCommand authenticate.Command
	}
)

func newService(
	signInCmd signin.Command,
	signOutCmd signout.Command,
	authenticateCmd authenticate.Command,
) API {
	return &service{
		signInCommand:       signInCmd,
		signOutCommand:      signOutCmd,
		authenticateCommand: authenticateCmd,
	}
}

func (s *service) SignIn(ctx context.Context, username string, password string) SignInResponse {
	req := signin.NewRequest(ctx, username, password)
	res := s.signInCommand.Execute(req)

	return SignInResponse{
		SessionID: res.SessionID().String(),
		Err:       res.Error(),
	}
}

func (s *service) SignOut(ctx context.Context, sessionID string) SignOutResponse {
	authRes := s.authenticateCommand.Execute(authenticate.NewRequest(ctx, sessionID))
	if authRes.Error() != nil {
		return SignOutResponse{Err: authRes.Error()}
	}

	res := s.signOutCommand.Execute(signout.NewRequest(ctx, authRes.User()))
	return SignOutResponse{Err: res.Error()}
}

func (s *service) AuthenticateUser(ctx context.Context, sessionID string) AuthenticateUserResponse {
	res := s.authenticateCommand.Execute(authenticate.NewRequest(ctx, sessionID))
	if res.Error() != nil {
		return AuthenticateUserResponse{Err: res.Error()}
	}

	user := res.User()

	return AuthenticateUserResponse{
		CurrentUser: struct {
			ID    uuid.UUID
			Roles []string
		}{
			ID:    user.ID().UUID(),
			Roles: []string{},
		},
	}
}

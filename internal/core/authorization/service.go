package authorization

import (
	"context"
	"financo/internal/core/authorization/domain/passworddigest"
	"financo/internal/core/authorization/domain/userID"

	"financo/internal/core/authorization/commands/authenticate"
	"financo/internal/core/authorization/commands/signin"
	"financo/internal/core/authorization/commands/signout"
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
		userIDEncoder       userID.Encoder
		passwordDigestCrypt passworddigest.Crypt
	}
)

func newService(
	signInCmd signin.Command,
	signOutCmd signout.Command,
	authenticateCmd authenticate.Command,
	userIDEncoder userID.Encoder,
	pwdCrypt passworddigest.Crypt,
) API {
	return &service{
		signInCommand:       signInCmd,
		signOutCommand:      signOutCmd,
		authenticateCommand: authenticateCmd,
		userIDEncoder:       userIDEncoder,
		passwordDigestCrypt: pwdCrypt,
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

package authorization

import "github.com/google/uuid"

type (
	SignInResponse struct {
		SessionID string
		Err       error
	}

	SignOutResponse struct {
		Err error
	}

	AuthenticateUserResponse struct {
		CurrentUser struct {
			ID    uuid.UUID
			Roles []string
		}
		Err error
	}
)

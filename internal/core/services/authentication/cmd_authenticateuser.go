package authentication

import (
	"context"

	"github.com/NordGus/fnncr/internal/core/domain/session"
	"github.com/google/uuid"
)

type AuthenticateUserReq struct {
	SessionID string
}

type AuthenticateUserResp struct {
	User struct {
		ID       uuid.UUID
		Username string
	}
}

func (s *Service) AuthenticateUser(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error) {
	sessionID, err := session.ParseIDFromString(req.SessionID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	session, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	usr, err := s.userRepo.GetUserByID(ctx, session.UserID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	return AuthenticateUserResp{
		User: struct {
			ID       uuid.UUID
			Username string
		}{ID: usr.ID, Username: usr.Username.String()},
	}, nil
}

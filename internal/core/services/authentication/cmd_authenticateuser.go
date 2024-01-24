package authentication

import (
	"context"
	"errors"

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

func (s *service) AuthenticateSession(ctx context.Context, req AuthenticateUserReq) (AuthenticateUserResp, error) {
	sessionID, err := session.ParseIDFromString(req.SessionID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	sssn, err := s.sessionRepo.Get(ctx, sessionID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	usr, err := s.userRepo.GetByID(ctx, sssn.UserID)
	if err != nil {
		return AuthenticateUserResp{}, err
	}

	if sssn.Expired(s.sessionMaxAge, usr) {
		return AuthenticateUserResp{}, errors.Join(session.ErrExpired, s.sessionRepo.Delete(ctx, sssn))
	}

	return AuthenticateUserResp{
		User: struct {
			ID       uuid.UUID
			Username string
		}{ID: usr.ID, Username: usr.Username.String()},
	}, nil
}

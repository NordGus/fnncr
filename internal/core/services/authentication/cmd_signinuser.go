package authentication

import (
	"context"
	"crypto/rand"
	"io"

	"github.com/NordGus/fnncr/internal/core/domain/session"
	"github.com/NordGus/fnncr/internal/core/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type SignInUserReq struct {
	username string
	password string
}

type SignInUserResp struct {
	SessionID string
}

func (s *Service) SignInUser(ctx context.Context, req SignInUserReq) (SignInUserResp, error) {
	sessionIDBuf := make([]byte, session.IdByteSize)

	username, err := user.NewUsername(req.username)
	if err != nil {
		return SignInUserResp{}, err
	}

	usr, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return SignInUserResp{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordDigest.String()), []byte(req.password))
	if err != nil {
		return SignInUserResp{}, err
	}

	_, err = io.ReadFull(rand.Reader, sessionIDBuf)
	if err != nil {
		return SignInUserResp{}, err
	}

	sessionID, err := session.NewID([session.IdByteSize]byte(sessionIDBuf))
	if err != nil {
		return SignInUserResp{}, err
	}

	err = s.sessionRepo.CreateSession(ctx, session.New(sessionID, usr.ID))
	if err != nil {
		return SignInUserResp{}, err
	}

	return SignInUserResp{SessionID: sessionID.String()}, nil
}

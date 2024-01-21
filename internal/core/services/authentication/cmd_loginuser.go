package authentication

import (
	"context"
	"crypto/rand"
	"io"

	"github.com/NordGus/fnncr/internal/core/domain/session"
	"github.com/NordGus/fnncr/internal/core/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserReq struct {
	username string
	password string
}

type LoginUserResp struct {
	SessionID string
}

func (s *Service) LoginUser(ctx context.Context, req LoginUserReq) (LoginUserResp, error) {
	sessionIDBuf := make([]byte, session.IdByteSize)

	username, err := user.NewUsername(req.username)
	if err != nil {
		return LoginUserResp{}, err
	}

	usr, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return LoginUserResp{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordDigest.String()), []byte(req.password))
	if err != nil {
		return LoginUserResp{}, err
	}

	_, err = io.ReadFull(rand.Reader, sessionIDBuf)
	if err != nil {
		return LoginUserResp{}, err
	}

	sessionID, err := session.NewID([session.IdByteSize]byte(sessionIDBuf))
	if err != nil {
		return LoginUserResp{}, err
	}

	err = s.sessionRepo.CreateSession(ctx, session.New(sessionID, usr.ID))
	if err != nil {
		return LoginUserResp{}, err
	}

	return LoginUserResp{SessionID: sessionID.String()}, nil
}

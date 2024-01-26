package authentication

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SignOutUserReq struct {
	UserID string
}

type SignOutUserResp struct {
	Success bool
}

func (s *service) SignOut(ctx context.Context, req SignOutUserReq) (SignOutUserResp, error) {
	uid, err := uuid.Parse(req.UserID)
	if err != nil {
		return SignOutUserResp{}, err
	}

	usr, err := s.userRepo.GetByID(ctx, uid)
	if err != nil {
		return SignOutUserResp{}, err
	}

	usr.SessionVersion += 1
	usr.UpdatedAt = time.Now().UTC()

	usr, err = s.userRepo.Update(ctx, usr)
	if err != nil {
		return SignOutUserResp{}, err
	}

	return SignOutUserResp{Success: true}, nil
}

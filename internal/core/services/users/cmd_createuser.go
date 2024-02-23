package users

import (
	"context"
	"errors"
	"time"

	"financo/internal/core/domain/user"
	"github.com/google/uuid"
)

var (
	ErrPasswordAndPasswordConfirmationMismatch = errors.New("password must be equal to password_confirmation")
	ErrUsernameNotUnique                       = errors.New("username is not unique")
)

type (
	CreateReq struct {
		Username             string
		Password             string
		PasswordConfirmation string
	}

	CreateResp struct {
		User struct {
			ID        uuid.UUID
			Username  string
			CreatedAt time.Time
			UpdatedAt time.Time
		}
	}
)

func (s *service) CreateUser(ctx context.Context, req CreateReq) (CreateResp, error) {
	var err error

	if req.Password != req.PasswordConfirmation {
		err = errors.Join(err, ErrPasswordAndPasswordConfirmationMismatch)
	}

	username, errUn := user.NewUsername(req.Username)
	if errUn != nil {
		err = errors.Join(err, errUn)
	}

	pw, errPw := user.NewPasswordDigestFromPassword(req.Password)
	if errPw != nil {
		err = errors.Join(err, errPw)
	}

	_, errUnique := s.userRepo.GetByUsername(ctx, username)
	if errUnique == nil {
		err = errors.Join(err, ErrUsernameNotUnique)
	}

	if err != nil {
		return CreateResp{}, err
	}

	entity := user.New(uuid.New(), username, pw, 0, time.Now().UTC(), time.Now().UTC())

	err = s.userRepo.Create(ctx, entity)
	if err != nil {
		return CreateResp{}, err
	}

	return CreateResp{
		User: struct {
			ID        uuid.UUID
			Username  string
			CreatedAt time.Time
			UpdatedAt time.Time
		}{
			ID:        entity.ID,
			Username:  entity.Username.String(),
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}, nil
}

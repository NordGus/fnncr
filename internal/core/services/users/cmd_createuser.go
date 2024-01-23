package users

import (
	"context"
	"errors"
	"time"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (s *service) Create(ctx context.Context, req CreateReq) (CreateResp, error) {
	var err error

	if req.Password != req.PasswordConfirmation {
		err = errors.Join(err, ErrPasswordAndPasswordConfirmationMismatch)
	}

	username, errUn := user.NewUsername(req.Username)
	if errUn != nil {
		err = errors.Join(err, errUn)
	}

	pw, errPw := user.NewPassword(req.Password)
	if errPw != nil {
		err = errors.Join(err, errPw)
	}

	hash, errPwh := bcrypt.GenerateFromPassword([]byte(pw.String()), bcrypt.DefaultCost)
	if errPwh != nil {
		err = errors.Join(err, errPwh)
	}

	pwd, errPwd := user.NewPasswordDigest(string(hash))
	if errPwd != nil {
		err = errors.Join(err, errPwd)
	}

	_, errUnique := s.userRepo.GetByUsername(ctx, username)
	if errUnique == nil {
		err = errors.Join(err, ErrUsernameNotUnique)
	}

	if err != nil {
		return CreateResp{}, err
	}

	entity := user.New(uuid.New(), username, pwd, time.Now().UTC(), time.Now().UTC())

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

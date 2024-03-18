package signout

import (
	"context"
	"errors"
	"time"

	"financo/internal/core_old/authorization/domain/sessionversion"
	"financo/internal/core_old/authorization/domain/timestamp"
	"financo/internal/core_old/authorization/domain/user"
)

type (
	Command interface {
		Execute(req Request) Response
	}

	UserRepository interface {
		Save(ctx context.Context, entity user.Entity) error
	}

	command struct {
		userRepository UserRepository
	}
)

func New(userRepository UserRepository) Command {
	return &command{
		userRepository: userRepository,
	}
}

func (c *command) Execute(req Request) Response {
	var (
		res Response
	)

	ver, err := sessionversion.New(req.user.SessionVersion().Uint32() + 1)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	ut, err := timestamp.New(time.Now())
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	err = c.userRepository.Save(
		req.ctx,
		user.New(req.user.ID(), req.user.Username(), req.user.PasswordDigest(), ver, req.user.CreatedAt(), ut),
	)

	res.err = errors.Join(res.err, err)

	return res
}

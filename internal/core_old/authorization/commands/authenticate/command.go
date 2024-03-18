package authenticate

import (
	"context"
	"errors"
	"time"

	"financo/internal/core_old/authorization/domain/session"
	"financo/internal/core_old/authorization/domain/sessionID"
	"financo/internal/core_old/authorization/domain/user"
	"financo/internal/core_old/authorization/domain/userID"
)

type (
	Command interface {
		Execute(req Request) Response
	}

	UserRepository interface {
		GetByID(ctx context.Context, id userID.Value) (user.Entity, error)
	}

	SessionRepository interface {
		Get(ctx context.Context, id sessionID.Value) (session.Entity, error)
		Delete(ctx context.Context, id sessionID.Value) error
	}

	command struct {
		userRepository    UserRepository
		sessionRepository SessionRepository
		sessionMaxAge     time.Duration
		sessionStaleAge   time.Duration
		sessionIDEncoder  sessionID.Encoder
	}
)

func New(
	userRepository UserRepository,
	sessionRepository SessionRepository,
	sessionMaxAge time.Duration,
	sessionStaleAge time.Duration,
	sessionIDEncoder sessionID.Encoder,
) Command {
	return &command{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		sessionMaxAge:     sessionMaxAge,
		sessionStaleAge:   sessionStaleAge,
		sessionIDEncoder:  sessionIDEncoder,
	}
}

func (c *command) Execute(req Request) Response {
	var res Response

	id, err := sessionID.NewFromString(req.sessionID, c.sessionIDEncoder)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	sssn, err := c.sessionRepository.Get(req.ctx, id)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	usr, err := c.userRepository.GetByID(req.ctx, sssn.UserID())
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	res.err = errors.Join(res.err, sssn.Expired(usr, c.sessionMaxAge), sssn.IsTooOld(c.sessionStaleAge))

	if res.err != nil {
		res.err = errors.Join(res.err, c.sessionRepository.Delete(req.ctx, id))
	} else {
		res.user = usr
	}

	return res
}

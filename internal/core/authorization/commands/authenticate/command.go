package authenticate

import (
	"context"
	"errors"
	"time"

	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
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
	}
)

func New(
	userRepository UserRepository,
	sessionRepository SessionRepository,
	sessionMaxAge time.Duration,
	sessionStaleAge time.Duration,
) Command {
	return &command{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		sessionMaxAge:     sessionMaxAge,
		sessionStaleAge:   sessionStaleAge,
	}
}

func (c *command) Execute(req Request) Response {
	var res Response

	if res.err = errors.Join(req.err, res.err); res.err != nil {
		return res
	}

	sess, err := c.sessionRepository.Get(req.ctx, req.sessionID)
	if res.err = errors.Join(err, res.err); res.err != nil {
		return res
	}

	usr, err := c.userRepository.GetByID(req.ctx, sess.UserID())
	if res.err = errors.Join(err, res.err); res.err != nil {
		return res
	}

	res.err = errors.Join(res.err, sess.Expired(usr, c.sessionMaxAge), sess.IsTooOld(c.sessionStaleAge))

	if res.err != nil {
		res.err = errors.Join(res.err, c.sessionRepository.Delete(req.ctx, req.sessionID))
	} else {
		res.user = usr
	}

	return res
}

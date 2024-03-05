package authenticate

import (
	"context"
	"encoding/base64"
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

	Request struct {
		ctx       context.Context
		sessionID string
	}

	Response struct {
		user user.Entity
		err  error
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

func NewRequest(ctx context.Context, sessionID string) Request {
	return Request{
		ctx:       ctx,
		sessionID: sessionID,
	}
}

func (c *command) Execute(req Request) Response {
	var res Response

	id, err := sessionID.NewFromString(req.sessionID, base64.URLEncoding)
	if res.err = errors.Join(err, res.err); res.err != nil {
		return res
	}

	sess, err := c.sessionRepository.Get(req.ctx, id)
	if res.err = errors.Join(err, res.err); res.err != nil {
		return res
	}

	usr, err := c.userRepository.GetByID(req.ctx, sess.UserID())
	if res.err = errors.Join(err, res.err); res.err != nil {
		return res
	}

	// TODO: Refactor session.Entity.Expired and session.Entity.IsTooOld to return errors instead of booleans
	if sess.Expired(usr, c.sessionMaxAge) || sess.IsTooOld(c.sessionStaleAge) {
		res.err = errors.Join(errors.New("session expired"), c.sessionRepository.Delete(req.ctx, id))
	} else {
		res.user = usr
	}

	return res
}

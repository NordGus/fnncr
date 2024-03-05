package signin

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"time"

	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/username"
)

type (
	Command interface {
		Execute(req Request) Response
	}

	UserRepository interface {
		GetByUsername(ctx context.Context, value username.Value) (user.Entity, error)
	}

	SessionRepository interface {
		Create(ctx context.Context, entity session.Entity) error
	}

	Request struct {
		ctx      context.Context
		username string
		password string
	}

	Response struct {
		sessionID string
		err       error
	}

	command struct {
		userRepository    UserRepository
		sessionRepository SessionRepository
	}
)

func New(userRepository UserRepository, sessionRepository SessionRepository) Command {
	return &command{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func NewRequest(ctx context.Context, username string, password string) Request {
	return Request{
		ctx:      ctx,
		username: username,
		password: password,
	}
}

func (c *command) Execute(req Request) Response {
	var (
		res          Response
		sessionIDBuf = make([]byte, sessionID.ByteSize)
	)

	un, err := username.New(req.username)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	usr, err := c.userRepository.GetByUsername(req.ctx, un)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	err = usr.PasswordDigest().Compare(req.password)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	_, err = io.ReadFull(rand.Reader, sessionIDBuf)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	sid, err := sessionID.New([sessionID.ByteSize]byte(sessionIDBuf), base64.URLEncoding)
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	ct, err := timestamp.New(time.Now())
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	err = c.sessionRepository.Create(req.ctx, session.New(sid, usr.SessionVersion(), ct, usr.ID()))
	if res.err = errors.Join(res.err, err); res.err != nil {
		return res
	}

	res.sessionID = sid.String()

	return res
}

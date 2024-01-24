package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/NordGus/fnncr/internal/core/domain/session"
	"github.com/NordGus/fnncr/internal/ports"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCantParseSession = errors.New("session repository: failed to parse session")
)

type (
	sessionRepository struct {
		conn *redis.Client
	}

	record struct {
		ID      string `redis:"id"`
		UserID  string `redis:"userID"`
		Version int32  `redis:"version"`
	}
)

func NewSessionRepository(conn *redis.Client) ports.SessionRepository {
	return &sessionRepository{
		conn: conn,
	}
}

func (repo *sessionRepository) Create(ctx context.Context, session session.Session) error {
	var (
		key  = fmt.Sprintf("session:%s", session.ID.String())
		rcrd = record{
			ID:      session.ID.String(),
			UserID:  session.UserID.String(),
			Version: session.Version,
		}
	)

	_, err := repo.conn.HSet(ctx, key, rcrd).Result()
	if err != nil {
		return errors.Join(ports.ErrSessionNotCreated, err)
	}

	return nil
}

func (repo *sessionRepository) Get(ctx context.Context, id session.ID) (session.Session, error) {
	var (
		key = fmt.Sprintf("session:%s", id.String())

		rcrd record
	)

	err := repo.conn.HGetAll(ctx, key).Scan(&rcrd)
	if err != nil {
		return session.Session{}, errors.Join(ports.ErrSessionNotFound, err)
	}

	userID, err := uuid.Parse(rcrd.UserID)
	if err != nil {
		return session.Session{}, errors.Join(ErrCantParseSession, err)
	}

	return session.New(id, userID, rcrd.Version), nil
}

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
	ErrCantParseSession = errors.New("failed to parse session")
)

type (
	SessionRepository struct {
		conn *redis.Client
	}

	record struct {
		ID     string `redis:"id"`
		UserID string `redis:"userID"`
	}
)

func NewSessionRepository(conn *redis.Client) ports.SessionRepository {
	return &SessionRepository{
		conn: conn,
	}
}

func (repo *SessionRepository) Create(ctx context.Context, session session.Session) error {
	var (
		key   = fmt.Sprintf("session:%s", session.ID.String())
		value = record{ID: session.ID.String(), UserID: session.UserID.String()}
	)

	_, err := repo.conn.HSet(ctx, key, value).Result()
	if err != nil {
		return errors.Join(ports.ErrSessionNotCreated, err)
	}

	return nil
}

func (repo *SessionRepository) Get(ctx context.Context, id session.ID) (session.Session, error) {
	var (
		key = fmt.Sprintf("session:%s", id.String())

		result record
	)

	err := repo.conn.HGetAll(ctx, key).Scan(&result)
	if err != nil {
		return session.Session{}, errors.Join(ports.ErrSessionNotFound, err)
	}

	userID, err := uuid.Parse(result.UserID)
	if err != nil {
		return session.Session{}, errors.Join(ErrCantParseSession, err)
	}

	return session.New(id, userID), nil
}

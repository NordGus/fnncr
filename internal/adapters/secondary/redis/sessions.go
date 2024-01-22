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

type (
	SessionRepository struct {
		conn *redis.Client
	}

	record struct {
		ID     string `redis:"id"`
		UserID string `redis:"userID"`
	}
)

func NewSessionRepository(conn *redis.Client) *SessionRepository {
	return &SessionRepository{
		conn: conn,
	}
}

func (repo *SessionRepository) CreateSession(ctx context.Context, session session.Session) error {
	var (
		key   = fmt.Sprintf("session:%s", session.ID.String())
		value = record{ID: session.ID.String(), UserID: session.UserID.String()}
	)

	_, err := repo.conn.HSet(ctx, key, value).Result()
	if err != nil {
		return errors.Join(ports.ErrSessionNotFound, err)
	}

	return nil
}

func (repo *SessionRepository) GetSession(ctx context.Context, id session.ID) (session.Session, error) {
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
		return session.Session{}, errors.Join(ports.ErrSessionNotFound, err)
	}

	return session.New(id, userID), nil
}

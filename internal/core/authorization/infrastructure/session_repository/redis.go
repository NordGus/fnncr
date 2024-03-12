package session_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"financo/internal/core/authorization"
	"financo/internal/core/authorization/domain/session"
	"financo/internal/core/authorization/domain/sessionID"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/userID"
	"github.com/redis/go-redis/v9"
)

type (
	RedisService interface {
		Client() *redis.Client
	}

	redisRepository struct {
		conn          *redis.Client
		userIDEncoder userID.Encoder
	}

	redisRecord struct {
		ID        string    `redis:"id"`
		UserID    string    `redis:"userID"`
		Version   uint32    `redis:"version"`
		CreatedAt time.Time `redis:"created_at"`
	}
)

func NewRedisRepository(serv RedisService, userIDEncoder userID.Encoder) authorization.SessionRepository {
	return &redisRepository{
		conn:          serv.Client(),
		userIDEncoder: userIDEncoder,
	}
}

func redisSessionKey(id sessionID.Value) string {
	return fmt.Sprintf("session:%s", id.String())
}

func (r *redisRepository) Get(ctx context.Context, id sessionID.Value) (session.Entity, error) {
	var (
		key    = redisSessionKey(id)
		record redisRecord
	)

	err := r.conn.HGetAll(ctx, key).Scan(&record)
	if err != nil {
		return session.Entity{}, errors.Join(ErrSessionNotFound, err)
	}

	uid, err := userID.New(record.UserID, r.userIDEncoder)
	if err != nil {
		return session.Entity{}, errors.Join(ErrSessionCantBeParsed, err)
	}

	ver, err := sessionversion.New(record.Version)
	if err != nil {
		return session.Entity{}, errors.Join(ErrSessionCantBeParsed, err)
	}

	ct, err := timestamp.New(record.CreatedAt)
	if err != nil {
		return session.Entity{}, errors.Join(ErrSessionCantBeParsed, err)
	}

	return session.New(id, ver, ct, uid), nil
}

func (r *redisRepository) Create(ctx context.Context, entity session.Entity) error {
	var (
		key    = redisSessionKey(entity.ID())
		record = redisRecord{
			ID:        entity.ID().String(),
			UserID:    entity.UserID().String(),
			Version:   entity.Version().Uint32(),
			CreatedAt: entity.CreatedAt().Time(),
		}
	)

	_, err := r.conn.HSet(ctx, key, record).Result()
	if err != nil {
		return errors.Join(ErrSessionWasNotCreated, err)
	}

	return nil
}

func (r *redisRepository) Delete(ctx context.Context, id sessionID.Value) error {
	key := redisSessionKey(id)

	// This is inefficient because it's time complexity is O(M) where M is the number of keys in the session hash.
	_, err := r.conn.Del(ctx, key).Result()
	if err != nil {
		return errors.Join(ErrSessionWasNotDeleted, err)
	}

	return nil
}

package sessionrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/NordGus/fnncr/authentication"
	"github.com/redis/go-redis/v9"
)

type (
	RedisRepository struct {
		ctx    context.Context
		client *redis.Client
	}

	RedisOpts struct {
		Options redis.Options
		Ctx     context.Context
	}

	RedisConfigFunc func(opts *RedisOpts)
)

var (
	ErrInvalidRedisOpts = errors.New("sessionrepo: invalid redis opts")

	redisDefaults = RedisOpts{
		Options: redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
			PoolSize: 5,
		},
	}
)

func NewRedisRepository(configs ...RedisConfigFunc) (*RedisRepository, error) {
	var (
		opts = redisDefaults
	)

	for i := 0; i < len(configs); i++ {
		configs[i](&opts)
	}

	if opts.Ctx == nil {
		return nil, ErrInvalidRedisOpts
	}

	return &RedisRepository{
		ctx:    opts.Ctx,
		client: redis.NewClient(&opts.Options),
	}, nil
}

// Create stores a new session in the RedisRepository in a Redis Hash with key session:{sessionID}.
func (repo *RedisRepository) Create(ctx context.Context, sessionID string, userID int64) error {
	select {
	case <-repo.ctx.Done():
		return repo.ctx.Err()
	default:
		var (
			key     = fmt.Sprintf("session:%s", sessionID)
			session = Session{ID: sessionID, UserID: userID}
		)

		_, err := repo.client.HSet(ctx, key, session).Result()
		if err != nil {
			return err
		}

		return nil
	}
}

// Get retrieves session data related to the sessionID.
func (repo *RedisRepository) Get(ctx context.Context, sessionID string) (authentication.SessionRecord, error) {
	select {
	case <-repo.ctx.Done():
		return nil, repo.ctx.Err()
	default:
		var session Session

		err := repo.client.HGetAll(ctx, fmt.Sprintf("session:%s", sessionID)).Scan(&session)
		if err != nil {
			return session, err
		}

		return session, nil
	}
}

func (repo *RedisRepository) Close() error {
	return repo.client.Close()
}

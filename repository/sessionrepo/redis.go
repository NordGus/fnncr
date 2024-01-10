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

// Create implements authentication.SessionStore.
func (repo *RedisRepository) Create(sessionID string, userID int64) error {
	return nil
}

// Get implements authentication.SessionStore.
func (repo *RedisRepository) Get(sessionID string) (authentication.SessionRecord, error) {
	var (
		session = Session{ID: sessionID, UserID: 42}
	)

	values, err := repo.client.Get(repo.ctx, sessionID).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(values)

	return session, nil
}

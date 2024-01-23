package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	Service interface {
		Client() *redis.Client
		Close() error
		Health() map[string]string
	}

	service struct {
		client *redis.Client
	}
)

// New returns a new Service. It panics if it can't open a connection to the database.
func New(opts redis.Options) Service {
	var (
		client      = redis.NewClient(&opts)
		ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	)
	defer cancel()

	res, err := client.Conn().Ping(ctx).Result()
	if err != nil || res != "PONG" {
		log.Fatalln(err)
	}

	return &service{client: client}
}

func (s *service) Client() *redis.Client {
	return s.client
}

func (s *service) Close() error {
	return s.client.Close()
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := s.client.Conn().Ping(ctx).Result()
	if err != nil || res != "PONG" {
		return map[string]string{
			"message": "Database is down",
		}
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

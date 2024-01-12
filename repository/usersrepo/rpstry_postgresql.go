package usersrepo

import (
	"context"
	"errors"
	"time"

	"github.com/NordGus/fnncr/authentication"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	PostgreSQLRepository struct {
		ctx    context.Context
		client *pgxpool.Pool
	}

	PostgreSQLOpts struct {
		Options pgxpool.Config
		Ctx     context.Context
	}

	PostgreSQLConfigFunc func(opts *PostgreSQLOpts)
)

var (
	ErrInvalidPostgreSQLOpts = errors.New("usersrepo: invalid postgresql opts")

	postgresqlDefaults = PostgreSQLOpts{
		Options: pgxpool.Config{
			ConnConfig: &pgx.ConnConfig{
				Config: pgconn.Config{
					Host:     "127.0.0.1",
					Port:     5432,
					Database: "fnncr_dev",
					User:     "fnncr",
					Password: "local_dev",
				},
			},
			MaxConns:              3,
			MinConns:              1,
			MaxConnLifetime:       15 * time.Second,
			MaxConnLifetimeJitter: 2 * time.Second,
			MaxConnIdleTime:       30 * time.Second,
			HealthCheckPeriod:     20 * time.Second,
		},
	}
)

func NewPostgreSQLRepository(configs ...PostgreSQLConfigFunc) (*PostgreSQLRepository, error) {
	var (
		opts = postgresqlDefaults
	)

	for i := 0; i < len(configs); i++ {
		configs[i](&opts)
	}

	if opts.Ctx == nil {
		return nil, ErrInvalidPostgreSQLOpts
	}

	pool, err := pgxpool.NewWithConfig(opts.Ctx, &opts.Options)
	if err != nil {
		return nil, errors.Join(ErrInvalidPostgreSQLOpts, err)
	}

	return &PostgreSQLRepository{
		ctx:    opts.Ctx,
		client: pool,
	}, nil
}

// GetByID implements authentication.UserRepository.
func (repo *PostgreSQLRepository) GetByID(id int64) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

// GetByUsername implements authentication.UserRepository.
func (repo *PostgreSQLRepository) GetByUsername(username string) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

func (repo *PostgreSQLRepository) Close() error {
	repo.client.Close()

	return nil
}

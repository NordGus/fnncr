package usersrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NordGus/fnncr/authentication"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	PostgreSQLRepository struct {
		ctx    context.Context
		client *pgxpool.Pool
	}

	PostgreSQLOpts struct {
		Ctx context.Context

		Host     string
		Port     uint16
		Database string
		User     string
		Password string

		MaxConns              int32
		MinConns              int32
		MaxConnLifetime       time.Duration
		MaxConnLifetimeJitter time.Duration
		MaxConnIdleTime       time.Duration
		HealthCheckPeriod     time.Duration
	}

	PostgreSQLConfigFunc func(opts *PostgreSQLOpts)
)

var (
	ErrInvalidPostgreSQLOpts = errors.New("usersrepo: invalid postgresql opts")

	postgresqlDefaults = PostgreSQLOpts{
		Host:                  "127.0.0.1",
		Port:                  5432,
		Database:              "fnncr_dev",
		User:                  "fnncr",
		Password:              "local_dev",
		MaxConns:              3,
		MinConns:              1,
		MaxConnLifetime:       15 * time.Second,
		MaxConnLifetimeJitter: 2 * time.Second,
		MaxConnIdleTime:       30 * time.Second,
		HealthCheckPeriod:     20 * time.Second,
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

	config, err := pgxpool.ParseConfig(opts.connString())
	if err != nil {
		return nil, errors.Join(ErrInvalidPostgreSQLOpts, err)
	}

	pool, err := pgxpool.NewWithConfig(opts.Ctx, config)
	if err != nil {
		return nil, errors.Join(ErrInvalidPostgreSQLOpts, err)
	}

	return &PostgreSQLRepository{
		ctx:    opts.Ctx,
		client: pool,
	}, nil
}

// GetByID retrieves an authentication.UserRecord from the PostgreSQLRepository with the matching id.
func (repo *PostgreSQLRepository) GetByID(ctx context.Context, id int64) (authentication.UserRecord, error) {
	select {
	case <-repo.ctx.Done():
		return nil, repo.ctx.Err()
	default:
		var user User

		conn, err := repo.client.Acquire(ctx)
		if err != nil {
			return nil, err
		}
		defer conn.Release()

		err = conn.QueryRow(ctx, "SELECT id, username, password_digest FROM users WHERE id = $1", id).
			Scan(&user.ID, &user.AccessName, &user.PasswordDigest)
		if err != nil {
			return nil, err
		}

		fmt.Println(user)

		return user, nil
	}
}

// GetByUsername retrieves an authentication.UserRecord from the PostgreSQLRepository with the matching username.
func (repo *PostgreSQLRepository) GetByUsername(ctx context.Context, username string) (authentication.UserRecord, error) {
	select {
	case <-repo.ctx.Done():
		return nil, repo.ctx.Err()
	default:
		var user User

		conn, err := repo.client.Acquire(ctx)
		if err != nil {
			return nil, err
		}
		defer conn.Release()

		err = conn.QueryRow(ctx, "SELECT id, username, password_digest FROM users WHERE username = $1", username).
			Scan(&user.ID, &user.AccessName, &user.PasswordDigest)
		if err != nil {
			return nil, err
		}

		fmt.Println(user)

		return user, nil
	}
}

func (repo *PostgreSQLRepository) Close() error {
	repo.client.Close()

	return nil
}

func (opts PostgreSQLOpts) connString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s pool_max_conn_lifetime_jitter=%s",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
		opts.MaxConns,
		opts.MinConns,
		opts.MaxConnLifetime,
		opts.MaxConnIdleTime,
		opts.HealthCheckPeriod,
		opts.MaxConnLifetimeJitter,
	)
}

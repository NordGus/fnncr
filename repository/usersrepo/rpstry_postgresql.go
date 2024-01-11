package usersrepo

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/authentication"
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

	// TODO: implement sqliteDefaults SQLiteOpts values
	sqliteDefaults = PostgreSQLOpts{}
)

func NewPostgreSQLRepository(configs ...PostgreSQLConfigFunc) (*PostgreSQLRepository, error) {
	var (
		opts = sqliteDefaults
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

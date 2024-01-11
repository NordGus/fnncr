package usersrepo

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/authentication"
)

type (
	// TODO: implement PostgreSQLRepository struct
	PostgreSQLRepository struct {
	}

	// TODO: implement PostgreSQLOpts struct
	PostgreSQLOpts struct {
		Ctx context.Context
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

	return &PostgreSQLRepository{}, nil
}

// GetByID implements authentication.UserRepository.
func (repo *PostgreSQLRepository) GetByID(id int64) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

// GetByUsername implements authentication.UserRepository.
func (repo *PostgreSQLRepository) GetByUsername(username string) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

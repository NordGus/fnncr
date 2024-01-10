package usersrepo

import (
	"context"
	"errors"

	"github.com/NordGus/fnncr/authentication"
)

type (
	// TODO: implement SQLiteRepository struct
	SQLiteRepository struct {
	}

	// TODO: implement SQLiteOpts struct
	SQLiteOpts struct {
		Ctx context.Context
	}

	SQLiteConfigFunc func(opts *SQLiteOpts)
)

var (
	ErrInvalidSQLiteOpts = errors.New("usersrepo: invalid sqlite opts")

	// TODO: implement sqliteDefaults SQLiteOpts values
	sqliteDefaults = SQLiteOpts{}
)

func NewSQLiteRepository(configs ...SQLiteConfigFunc) (*SQLiteRepository, error) {
	var (
		opts = sqliteDefaults
	)

	for i := 0; i < len(configs); i++ {
		configs[i](&opts)
	}

	if opts.Ctx == nil {
		return nil, ErrInvalidSQLiteOpts
	}

	return &SQLiteRepository{}, nil
}

// GetByID implements authentication.UserRepository.
func (repo *SQLiteRepository) GetByID(id int64) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

// GetByUsername implements authentication.UserRepository.
func (repo *SQLiteRepository) GetByUsername(username string) (authentication.UserRecord, error) {
	return User{}, errors.New("unimplemented")
}

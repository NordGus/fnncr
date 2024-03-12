package user_repository

import "errors"

var (
	ErrUserNotFound                     = errors.New("user repository: user not found")
	ErrUserCantBeParsed                 = errors.New("user repository: user can't be parsed")
	ErrUserWasNotSaved                  = errors.New("user repository: user was not saved")
	ErrUserRepositoryIntegrityCorrupted = errors.New("user repository: integrity corrupted")
)

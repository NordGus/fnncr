package accounts_repository

import "errors"

var (
	ErrAccountNotFound        = errors.New("accounts_repository: not found")
	ErrAccountExists          = errors.New("accounts_repository: already exists")
	ErrFailedToSaveAccount    = errors.New("accounts_repository: failed to save account")
	ErrCorruptedAccount       = errors.New("accounts_repository: corrupted account")
	ErrInternalServiceFailure = errors.New("accounts_repository: internal service failure")
)

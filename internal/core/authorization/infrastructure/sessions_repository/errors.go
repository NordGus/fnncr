package sessions_repository

import "errors"

var (
	ErrSessionNotFound      = errors.New("session repository: session not found")
	ErrSessionCantBeParsed  = errors.New("session repository: session can't be parsed")
	ErrSessionWasNotCreated = errors.New("session repository: session was not created")
	ErrSessionWasNotDeleted = errors.New("session repository: session was not deleted")
)

package passworddigest

import "errors"

var (
	ErrPasswordEmpty       = errors.New("passworddigest: password is empty")
	ErrPasswordTooShort    = errors.New("passworddigest: password is too short")
	ErrPasswordTooLong     = errors.New("passworddigest: password is too long")
	ErrPasswordDoesntMatch = errors.New("passworddigest: password and password confirmation doesn't match")
	ErrHashInvalid         = errors.New("passworddigest: hash is invalid")
	ErrHashCostInvalid     = errors.New("passworddigest: hash cost is invalid")
)

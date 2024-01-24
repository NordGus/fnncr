package user

import "errors"

type Username string

const (
	usernameMinLen = 2
	usernameMaxLen = 20
)

var (
	ErrUsernameTooShort = errors.New("user: username is too short")
	ErrUsernameTooLong  = errors.New("user: username is too long")
)

func NewUsername(username string) (Username, error) {
	length := len(username)

	if length < usernameMinLen {
		return "", ErrUsernameTooShort
	}

	if length > usernameMaxLen {
		return "", ErrUsernameTooLong
	}

	return Username(username), nil
}

func (u Username) String() string {
	return string(u)
}

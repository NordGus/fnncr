package user

import "errors"

type Username string

const (
	usernameMinLen = 2
	usernameMaxLen = 20
)

var (
	ErrUsernameTooShort = errors.New("username is too short")
	ErrUsernameTooLong  = errors.New("username is too long")
)

func NewUsername(username string) (Username, error) {
	unLen := len(username)

	if unLen < usernameMinLen {
		return "", ErrUsernameTooShort
	}

	if unLen > usernameMaxLen {
		return "", ErrUsernameTooLong
	}

	return Username(username), nil
}

func (u Username) String() string {
	return string(u)
}

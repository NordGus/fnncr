package user

import "errors"

type Password string

const (
	passwordMinLen = 8
	passwordMaxLen = 64
)

var (
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooLong  = errors.New("password is too long")
)

func NewPassword(pw string) (Password, error) {
	pwLen := len(pw)

	if pwLen < passwordMinLen {
		return "", ErrPasswordTooShort
	}

	if pwLen > passwordMaxLen {
		return "", ErrPasswordTooLong
	}

	return Password(pw), nil
}

func (pw Password) String() string {
	return string(pw)
}

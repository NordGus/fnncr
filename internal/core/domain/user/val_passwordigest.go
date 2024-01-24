package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordDigest string

const (
	passwordMinLen = 8
	passwordMaxLen = 64
)

var (
	ErrPasswordEmpty    = errors.New("user: password is empty")
	ErrPasswordTooShort = errors.New("user: password is too short")
	ErrPasswordTooLong  = errors.New("user: password is too long")
)

func NewPasswordDigest(digest string) (PasswordDigest, error) {
	if digest == "" {
		return "", ErrPasswordEmpty
	}

	return PasswordDigest(digest), nil
}

func NewPasswordDigestFromPassword(pw string) (PasswordDigest, error) {
	length := len(pw)

	if length == 0 {
		return "", ErrPasswordEmpty
	}

	if length < passwordMinLen {
		return "", ErrPasswordTooShort
	}

	if length > passwordMaxLen {
		return "", ErrPasswordTooLong
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return PasswordDigest(hash), nil
}

func (pwd PasswordDigest) String() string {
	return string(pwd)
}

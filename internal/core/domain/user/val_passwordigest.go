package user

import "errors"

type PasswordDigest string

var (
	ErrPasswordDigestEmpty = errors.New("password digest is empty")
)

func NewPasswordDigest(digest string) (PasswordDigest, error) {
	if digest == "" {
		return "", ErrPasswordDigestEmpty
	}

	return PasswordDigest(digest), nil
}

func (pwd PasswordDigest) String() string {
	return string(pwd)
}

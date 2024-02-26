package username

import (
	"errors"
	"fmt"
)

const (
	minLen = 6
	maxLen = 20
)

type Value struct {
	username string
}

var (
	ErrBlank    = errors.New("can't be blank")
	ErrTooShort = fmt.Errorf("is too short (must be at lest %d charaters long)", minLen)
	ErrTooLong  = fmt.Errorf("is too short (must be at most %d charaters long)", maxLen)
)

func New(username string) (Value, error) {
	var (
		errs error
		uLen = len(username)
	)

	if uLen <= 0 {
		errs = errors.Join(errs, ErrBlank)
	}

	if uLen < minLen {
		errs = errors.Join(errs, ErrTooShort)
	}

	if uLen > maxLen {
		errs = errors.Join(errs, ErrTooLong)
	}

	return Value{
		username: username,
	}, errs
}

func (v *Value) String() string {
	return v.username
}

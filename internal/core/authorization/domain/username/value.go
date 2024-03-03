package username

import (
	"errors"
)

type Value struct {
	username string
}

var (
	ErrBlank = errors.New("can't be blank")
)

func New(username string) (Value, error) {
	var errs error

	if username == "" {
		errs = errors.Join(errs, ErrBlank)
	}

	return Value{
		username: username,
	}, errs
}

func (v *Value) String() string {
	return v.username
}

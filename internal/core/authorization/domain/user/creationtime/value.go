package creationtime

import (
	"errors"
	"time"
)

var (
	ErrEmpty = errors.New("must be present")
)

type Value struct {
	createdAt time.Time
}

func New(createdAt time.Time) (Value, error) {
	var errs error

	if createdAt.Equal(time.Time{}) {
		errs = errors.Join(errs, ErrEmpty)
	}

	return Value{
		createdAt: createdAt.UTC(),
	}, errs
}

func (v *Value) Time() time.Time {
	return v.createdAt
}

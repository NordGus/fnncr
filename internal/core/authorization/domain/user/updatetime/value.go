package updatetime

import (
	"errors"
	"time"
)

var (
	ErrEmpty = errors.New("must be present")
)

type Value struct {
	updatedAt time.Time
}

func New(updatedAt time.Time) (Value, error) {
	var errs error

	if updatedAt.Equal(time.Time{}) {
		errs = errors.Join(errs, ErrEmpty)
	}

	return Value{
		updatedAt: updatedAt.UTC(),
	}, errs
}

func (v *Value) Time() time.Time {
	return v.updatedAt
}

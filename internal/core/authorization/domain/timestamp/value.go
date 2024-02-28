package timestamp

import (
	"errors"
	"time"
)

var (
	ErrEmpty = errors.New("must be present")
)

type Value struct {
	timestamp time.Time
}

func New(moment time.Time) (Value, error) {
	var errs error

	if moment.Equal(time.Time{}) {
		errs = errors.Join(errs, ErrEmpty)
	}

	return Value{
		timestamp: moment.UTC(),
	}, errs
}

func (v *Value) Time() time.Time {
	return v.timestamp
}

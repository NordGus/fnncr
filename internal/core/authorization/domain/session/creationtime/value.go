package creationtime

import (
	"time"
)

const (
	maxAge = 30 * 24 * time.Hour
)

type Value struct {
	value time.Time
}

func New(when time.Time) (Value, error) {
	if isTooOld(when) {
		return Value{}, ErrCreationTimeExceedMaxAge
	}

	return Value{
		value: when.UTC(),
	}, nil
}

func (v Value) IsTooOld(maxAge time.Duration) bool {
	return time.Since(v.value) > maxAge
}

func (v Value) Time() time.Time {
	return v.value
}

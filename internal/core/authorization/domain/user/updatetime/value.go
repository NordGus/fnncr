package updatetime

import "time"

type Value struct {
	updatedAt time.Time
}

func New(updatedAt time.Time) (Value, error) {
	return Value{
		updatedAt: updatedAt,
	}, nil
}

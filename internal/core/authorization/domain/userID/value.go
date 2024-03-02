package userID

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrEmpty          = errors.New("must be present")
	ErrInvalid        = errors.New("is invalid")
	ErrFailedToParsed = errors.New("failed to parse")
)

type Value struct {
	value uuid.UUID
}

func New(id string) (Value, error) {
	var errs error

	if id == "" {
		errs = errors.Join(errs, ErrEmpty)
	}

	err := uuid.Validate(id)
	if err != nil {
		errs = errors.Join(errs, ErrInvalid)
	}

	val, err := uuid.Parse(id)
	if err != nil {
		errs = errors.Join(errs, ErrFailedToParsed)
	}

	return Value{
		value: val,
	}, errs
}

func (v Value) String() string {
	return v.value.String()
}

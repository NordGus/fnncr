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

type Encoder interface {
	Validate(s string) error
	Parse(s string) (uuid.UUID, error)
}

type Value struct {
	value   uuid.UUID
	encoder Encoder
}

func New(id string, encoder Encoder) (Value, error) {
	var errs error

	if id == "" {
		errs = errors.Join(errs, ErrEmpty)
	}

	err := encoder.Validate(id)
	if err != nil {
		errs = errors.Join(errs, ErrInvalid)
	}

	val, err := encoder.Parse(id)
	if err != nil {
		errs = errors.Join(errs, ErrFailedToParsed)
	}

	return Value{
		value:   val,
		encoder: encoder,
	}, errs
}

func (v Value) String() string {
	return v.value.String()
}

package uuid_encoder

import "github.com/google/uuid"

type Encoder struct{}

func New() Encoder {
	return Encoder{}
}

func (Encoder) Validate(s string) error {
	return uuid.Validate(s)
}

func (Encoder) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

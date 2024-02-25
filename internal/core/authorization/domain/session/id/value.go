package id

import (
	"errors"
)

const ByteSize = 64

type Encoder interface {
	EncodeToString(src []byte) string
	DecodeString(s string) ([]byte, error)
}

type Value struct {
	value   [ByteSize]byte
	encoder Encoder
}

func New(id [ByteSize]byte, encoder Encoder) (Value, error) {
	return Value{value: id, encoder: encoder}, nil
}

func NewFromString(id string, encoder Encoder) (Value, error) {
	bytes, err := encoder.DecodeString(id)
	if err != nil {
		return Value{}, errors.Join(ErrCantBeDecodedFromString, err)
	}

	return Value{value: [ByteSize]byte(bytes), encoder: encoder}, nil
}

func (v Value) String() string {
	return v.encoder.EncodeToString(v.value[:])
}

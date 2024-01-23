package session

import (
	"encoding/base64"
	"errors"
)

const (
	IdByteSize = 64
)

var (
	ErrIDStringCanNotBeParsed = errors.New("")
)

type ID [IdByteSize]byte

func NewID(id [IdByteSize]byte) (ID, error) {
	return id, nil
}

func ParseIDFromString(id string) (ID, error) {
	bytes, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return ID{}, errors.Join(ErrIDStringCanNotBeParsed, err)
	}

	return [IdByteSize]byte(bytes), nil
}

func (id ID) String() string {
	return base64.URLEncoding.EncodeToString(id[:])
}

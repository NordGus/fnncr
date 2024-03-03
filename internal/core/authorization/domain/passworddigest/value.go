package passworddigest

import "errors"

type Crypt interface {
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
	Cost(hashedPassword []byte) (int, error)
}

const (
	hashCost = 10
)

var (
	ErrHashInvalid     = errors.New("hash is invalid")
	ErrHashCostInvalid = errors.New("hash cost is invalid")
	ErrInvalidPassword = errors.New("invalid password")
)

type Value struct {
	hash  []byte
	crypt Crypt
}

func New(hash string, crypt Crypt) (Value, error) {
	var errs error

	cost, err := crypt.Cost([]byte(hash))
	if err != nil {
		errs = errors.Join(errs, ErrHashInvalid, err)
	}

	if cost != hashCost {
		errs = errors.Join(errs, ErrHashCostInvalid)
	}

	return Value{
		hash:  []byte(hash),
		crypt: crypt,
	}, errs
}

func (v Value) String() string {
	return string(v.hash)
}

func (v Value) Compare(password string) error {
	err := v.crypt.CompareHashAndPassword(v.hash, []byte(password))
	if err != nil {
		return errors.Join(err, ErrInvalidPassword)
	}

	return nil
}

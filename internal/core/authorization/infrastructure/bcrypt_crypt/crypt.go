package bcrypt_crypt

import "golang.org/x/crypto/bcrypt"

type Crypt struct{}

func New() Crypt {
	return Crypt{}
}

func (Crypt) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (Crypt) Cost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}

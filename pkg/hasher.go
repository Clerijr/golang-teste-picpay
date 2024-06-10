package pkg

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
	salt int
}

func NewHasher() *Hasher {
	return &Hasher{
		salt: 15,
	}
}

var (
	ErrEmptyString = errors.New("empty string given")
)

func (h *Hasher) Hash(password string) (string, error) {
	var err error
	var bytes []byte

	if len(password) == 0 {
		return "", ErrEmptyString
	}

	bytes, err = bcrypt.GenerateFromPassword([]byte(password), h.salt)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (h *Hasher) Compare(password, toCompare string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(toCompare), []byte(password))
	return err == nil

}

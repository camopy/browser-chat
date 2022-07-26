package bcrypthash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHash struct {
}

func (h *BcryptHash) GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error generating hash password: %v", err)
	}
	return string(hashedPassword), nil
}

func (h *BcryptHash) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

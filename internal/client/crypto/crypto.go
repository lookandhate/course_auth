package crypto

import "golang.org/x/crypto/bcrypt"

type BCryptPasswordManager struct {
}

func NewBCryptPasswordManager() *BCryptPasswordManager {
	return &BCryptPasswordManager{}
}

func (m BCryptPasswordManager) HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (m BCryptPasswordManager) ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

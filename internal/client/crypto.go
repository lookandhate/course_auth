package client

type PasswordManager interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash, password string) error
}

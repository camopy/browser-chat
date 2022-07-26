package passwordhash

type PasswordHash interface {
	GeneratePassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}

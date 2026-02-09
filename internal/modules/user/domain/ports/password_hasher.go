package ports

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

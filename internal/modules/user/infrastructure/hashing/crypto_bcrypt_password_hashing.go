package hashing

import (
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct{}

var _ ports.PasswordHasher = (*PasswordHasher)(nil)

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

// HashPassword generates a bcrypt hash for the given password.
func (p *PasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func (p *PasswordHasher) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

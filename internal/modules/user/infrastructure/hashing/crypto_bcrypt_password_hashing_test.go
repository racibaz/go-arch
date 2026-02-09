package hashing

import (
	"testing"

	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewPasswordHasher(t *testing.T) {
	// When
	hasher := NewPasswordHasher()

	// Then
	assert.NotNil(t, hasher)
	assert.IsType(t, &PasswordHasher{}, hasher)
}

func TestPasswordHasher_ImplementsPasswordHasherInterface(t *testing.T) {
	// Given
	hasher := NewPasswordHasher()

	// When & Then
	var _ ports.PasswordHasher = hasher
}

func TestPasswordHasher_HashPassword(t *testing.T) {
	testCases := []struct {
		name          string
		password      string
		expectError   bool
		expectedError string // partial match for error message
	}{
		{
			name:        "valid password",
			password:    "myStrongPassword123",
			expectError: false,
		},
		{
			name:        "empty password",
			password:    "",
			expectError: false, // bcrypt handles empty passwords, generates a hash
		},
		{
			name:        "short password",
			password:    "short",
			expectError: false, // bcrypt generates hash regardless of length
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			hasher := NewPasswordHasher()

			// When
			hash, err := hasher.HashPassword(tc.password)

			// Then
			if tc.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hash)
				assert.NotEqual(t, tc.password, hash, "Hash should not be the same as the original password")
				// Verify the generated hash can be checked
				assert.True(t, hasher.VerifyPassword(tc.password, hash), "Hashed password should be verifiable")
			}
		})
	}
}

func TestPasswordHasher_VerifyPassword(t *testing.T) {
	// Given
	hasher := NewPasswordHasher()
	password := "secret123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(t, err)
	validHash := string(hashedPassword)

	testCases := []struct {
		name            string
		passwordToCheck string
		hashToCompare   string
		expectedMatch   bool
	}{
		{
			name:            "correct password and valid hash",
			passwordToCheck: password,
			hashToCompare:   validHash,
			expectedMatch:   true,
		},
		{
			name:            "incorrect password and valid hash",
			passwordToCheck: "wrongpassword",
			hashToCompare:   validHash,
			expectedMatch:   false,
		},
		{
			name:            "empty password and valid hash",
			passwordToCheck: "",
			hashToCompare:   validHash,
			expectedMatch:   false,
		},
		{
			name:            "correct password and empty hash",
			passwordToCheck: password,
			hashToCompare:   "",
			expectedMatch:   false, // bcrypt.CompareHashAndPassword returns error for empty hash
		},
		{
			name:            "correct password and invalid hash format",
			passwordToCheck: password,
			hashToCompare:   "$2a$10$invalidhashformat", // Example of a malformed bcrypt hash
			expectedMatch:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// When
			isMatch := hasher.VerifyPassword(tc.passwordToCheck, tc.hashToCompare)

			// Then
			assert.Equal(t, tc.expectedMatch, isMatch)
		})
	}
}

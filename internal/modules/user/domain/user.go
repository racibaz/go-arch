package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/racibaz/go-arch/pkg/es"
)

const (
	UserAggregate = "users.User"
)

var (
	NameMinLength     = 3
	NameMaxLength     = 20
	PasswordMinLength = 6
)

// todo duplicate with post status
var (
	ErrNotFound      = errors.New("the user was not found")
	ErrAlreadyExists = errors.New("the user already exists")
	ErrEmptyId       = errors.New("id cannot be empty")
	ErrMinNameLength = errors.New(
		fmt.Sprintf("name must be at least %d characters long", NameMinLength),
	)
	ErrMaxNameLength = errors.New(
		fmt.Sprintf("name must not be more than %d characters", NameMaxLength),
	)
	ErrMinPasswordLength = errors.New(
		fmt.Sprintf("password must be at least %d characters long", PasswordMinLength),
	)
	ErrInvalidStatus       = errors.New("status is not valid")
	ErrInvalidCredentials  = errors.New("credentials are not valid")
	ErrStatusNotAcceptable = errors.New("status is not acceptable")
)

type User struct {
	es.Aggregate
	Name                 string
	Email                string
	Password             string
	RefreshTokenWeb      *string
	RefreshTokenWebAt    *time.Time
	RefreshTokenMobile   *string
	RefreshTokenMobileAt *time.Time
	Status               UserStatus
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// Create This factory method creates a new Post with default values if you want.
func Create(
	id, name, email, password string,
	status UserStatus,
	createdAt, updatedAt time.Time,
) (*User, error) {
	user := &User{
		Aggregate: es.NewAggregate(id, UserAggregate),
		Name:      name,
		Email:     email,
		Password:  password,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// validate the post before returning it
	err := user.validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) sanitize() {
	// Trim whitespace from the input parameters
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)
}

// Validate checks if the Post fields are valid.
func (u *User) validate() error {
	// Sanitize the input parameters
	u.sanitize()

	// Validate the input parameters
	if u.Aggregate.ID() == "" {
		return ErrEmptyId
	}

	if len(u.Name) < NameMinLength {
		return ErrMinNameLength
	}

	if len(u.Password) < PasswordMinLength {
		return ErrMinPasswordLength
	}

	if !IsValidStatus(u.Status) {
		return ErrInvalidStatus
	}

	// and more validations can be added here

	return nil
}

package domain

import (
	"errors"
	"fmt"
	"github.com/racibaz/go-arch/pkg/es"
	"strings"
	"time"
)

const (
	UserAggregate = "users.User"
)

var (
	UserNameMinLength = 3
	UserNameMaxLength = 20
	PasswordMinLength = 6
)

// todo duplicate with post status
var (
	ErrNotFound          = errors.New("the user was not found")
	ErrAlreadyExists     = errors.New("the user already exists")
	ErrEmptyId           = errors.New("id cannot be empty")
	ErrMinUserNameLength = errors.New(
		fmt.Sprintf("username must be at least %d characters long", UserNameMinLength),
	)
	ErrMaxUserNameLength = errors.New(
		fmt.Sprintf("username must not be more than %d characters", UserNameMaxLength),
	)
	ErrMinPasswordLength = errors.New(
		fmt.Sprintf("password must be at least %d characters long", PasswordMinLength),
	)
	ErrInvalidStatus = errors.New("status is not valid")
)

type User struct {
	es.Aggregate
	UserName  string
	Email     string
	Password  string
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create This factory method creates a new Post with default values if you want.
func Create(
	id, userName, email, password string,
	status UserStatus,
	createdAt, updatedAt time.Time,
) (*User, error) {
	user := &User{
		Aggregate: es.NewAggregate(id, UserAggregate),
		UserName:  userName,
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
	u.UserName = strings.TrimSpace(u.UserName)
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

	if len(u.UserName) < UserNameMinLength {
		return ErrMinUserNameLength
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

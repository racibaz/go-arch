package uuid

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrUuidCannotBeNil = errors.New("the uuid cannot be nil")

// Uuid is a wrapper around the uuid.UUID type.
type Uuid struct {
	Uuid *uuid.UUID
}

// NewUuid generates a new UUID and returns it wrapped in a Uuid struct.
func NewUuid() *Uuid {
	newUuid := uuid.New()
	return &Uuid{Uuid: &newUuid}
}

// NewID generates a new UUID and returns its string representation.
func NewID() string {
	return NewUuid().ToString()
}

// ToString converts the Uuid struct to its string representation.
func (uuid *Uuid) ToString() string {
	if uuid == nil || uuid.Uuid == nil {
		return ""
	}
	return uuid.Uuid.String()
}

func Parse(input string) (Uuid, error) {
	parsedUuid, err := uuid.Parse(input)
	if err != nil {
		return Uuid{}, fmt.Errorf("the uuid can not be parse: %w", err)
	}

	if parsedUuid == uuid.Nil {
		return Uuid{}, ErrUuidCannotBeNil
	}

	return Uuid{
		Uuid: &parsedUuid,
	}, nil
}

// ParseToString parses a string into a Uuid struct and returns its string representation.
func ParseToString(input string) (string, error) {
	parsedUuid, err := Parse(input)
	if err != nil {
		return "", fmt.Errorf("the uuid can not be parse: %w", err)
	}
	return parsedUuid.ToString(), nil
}

package uuid

import "github.com/google/uuid"

type Uuid struct {
	Uuid *uuid.UUID
}

func NewUuid() *Uuid {
	newUuid := uuid.New()
	return &Uuid{Uuid: &newUuid}
}

func (uuid *Uuid) ToString() string {
	if uuid == nil || uuid.Uuid == nil {
		return ""
	}
	return uuid.Uuid.String()
}

func Parse(input string) (Uuid, error) {
	parsedUuid, err := uuid.Parse(input)
	if err != nil {
		return Uuid{}, err
	}

	if parsedUuid == uuid.Nil {
		return Uuid{}, nil // or return an error if you prefer
	}

	return Uuid{
		Uuid: &parsedUuid,
	}, nil
}

func NewID() string {
	return NewUuid().ToString()
}

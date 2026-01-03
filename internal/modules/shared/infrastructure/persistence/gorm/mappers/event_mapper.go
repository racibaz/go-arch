package mappers

import (
	"errors"

	domain "github.com/racibaz/go-arch/internal/modules/shared/domain"
	entity "github.com/racibaz/go-arch/internal/modules/shared/infrastructure/persistence/gorm/entities"
)

func ToDomain(eventEntity *entity.Event) (*domain.Event, error) {
	if eventEntity == nil {
		return nil, errors.New("event entity is nil")
	}

	return &domain.Event{
		StreamID:      eventEntity.StreamID,
		StreamName:    eventEntity.StreamName,
		StreamVersion: eventEntity.StreamVersion,
		EventID:       eventEntity.EventID,
		EventName:     eventEntity.EventName,
		EventData:     eventEntity.EventData,
		OccurredAt:    eventEntity.OccurredAt,
	}, nil
}

func ToPersistence(event *domain.Event) (*entity.Event, error) {
	if event == nil {
		return nil, errors.New("event domain is nil")
	}

	return &entity.Event{
		StreamID:      event.StreamID,
		StreamName:    event.StreamName,
		StreamVersion: event.StreamVersion,
		EventID:       event.EventID,
		EventName:     event.EventName,
		EventData:     event.EventData,
		OccurredAt:    event.OccurredAt,
	}, nil
}

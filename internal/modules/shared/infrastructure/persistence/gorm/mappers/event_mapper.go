package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/shared/domain"
	entity "github.com/racibaz/go-arch/internal/modules/shared/infrastructure/persistence/gorm/entities"
)

func ToDomain(eventEntity entity.Event) domain.Event {

	return domain.Event{
		StreamID:      eventEntity.StreamID,
		StreamName:    eventEntity.StreamName,
		StreamVersion: eventEntity.StreamVersion,
		EventID:       eventEntity.EventID,
		EventName:     eventEntity.EventName,
		EventData:     eventEntity.EventData,
		OccurredAt:    eventEntity.OccurredAt,
	}

}

func ToPersistence(event domain.Event) entity.Event {

	return entity.Event{
		StreamID:      event.StreamID,
		StreamName:    event.StreamName,
		StreamVersion: event.StreamVersion,
		EventID:       event.EventID,
		EventName:     event.EventName,
		EventData:     event.EventData,
		OccurredAt:    event.OccurredAt,
	}
}

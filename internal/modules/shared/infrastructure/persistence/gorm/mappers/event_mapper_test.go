package mappers

import (
	"testing"
	"time"

	domain "github.com/racibaz/go-arch/internal/modules/shared/domain"
	entity "github.com/racibaz/go-arch/internal/modules/shared/infrastructure/persistence/gorm/entities"
	"github.com/stretchr/testify/assert"
)

func TestToDomain(t *testing.T) {
	testCases := []struct {
		name        string
		eventEntity *entity.Event
		expected    *domain.Event
		expectError bool
		errorMsg    string
	}{
		{
			name: "successful conversion",
			eventEntity: &entity.Event{
				StreamID:      "stream-123",
				StreamName:    "TestStream",
				StreamVersion: "1",
				EventID:       "event-456",
				EventName:     "TestEvent",
				EventData:     `{"key": "value"}`,
				OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			expected: &domain.Event{
				StreamID:      "stream-123",
				StreamName:    "TestStream",
				StreamVersion: "1",
				EventID:       "event-456",
				EventName:     "TestEvent",
				EventData:     `{"key": "value"}`,
				OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			expectError: false,
		},
		{
			name:        "nil event entity",
			eventEntity: nil,
			expected:    nil,
			expectError: true,
			errorMsg:    "event entity is nil",
		},
		{
			name:        "empty event entity",
			eventEntity: &entity.Event{},
			expected: &domain.Event{
				StreamID:      "",
				StreamName:    "",
				StreamVersion: "",
				EventID:       "",
				EventName:     "",
				EventData:     "",
				OccurredAt:    time.Time{},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// When
			result, err := ToDomain(tc.eventEntity)

			// Then
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expected.StreamID, result.StreamID)
				assert.Equal(t, tc.expected.StreamName, result.StreamName)
				assert.Equal(t, tc.expected.StreamVersion, result.StreamVersion)
				assert.Equal(t, tc.expected.EventID, result.EventID)
				assert.Equal(t, tc.expected.EventName, result.EventName)
				assert.Equal(t, tc.expected.EventData, result.EventData)
				assert.Equal(t, tc.expected.OccurredAt, result.OccurredAt)
			}
		})
	}
}

func TestToPersistence(t *testing.T) {
	testCases := []struct {
		name        string
		event       *domain.Event
		expected    *entity.Event
		expectError bool
		errorMsg    string
	}{
		{
			name: "successful conversion",
			event: &domain.Event{
				StreamID:      "stream-123",
				StreamName:    "TestStream",
				StreamVersion: "1",
				EventID:       "event-456",
				EventName:     "TestEvent",
				EventData:     `{"key": "value"}`,
				OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			expected: &entity.Event{
				StreamID:      "stream-123",
				StreamName:    "TestStream",
				StreamVersion: "1",
				EventID:       "event-456",
				EventName:     "TestEvent",
				EventData:     `{"key": "value"}`,
				OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			expectError: false,
		},
		{
			name:        "nil event domain",
			event:       nil,
			expected:    nil,
			expectError: true,
			errorMsg:    "event domain is nil",
		},
		{
			name:  "empty event domain",
			event: &domain.Event{},
			expected: &entity.Event{
				StreamID:      "",
				StreamName:    "",
				StreamVersion: "",
				EventID:       "",
				EventName:     "",
				EventData:     "",
				OccurredAt:    time.Time{},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// When
			result, err := ToPersistence(tc.event)

			// Then
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expected.StreamID, result.StreamID)
				assert.Equal(t, tc.expected.StreamName, result.StreamName)
				assert.Equal(t, tc.expected.StreamVersion, result.StreamVersion)
				assert.Equal(t, tc.expected.EventID, result.EventID)
				assert.Equal(t, tc.expected.EventName, result.EventName)
				assert.Equal(t, tc.expected.EventData, result.EventData)
				assert.Equal(t, tc.expected.OccurredAt, result.OccurredAt)
			}
		})
	}
}

func TestRoundTripConversion(t *testing.T) {
	// Test that converting entity -> domain -> entity preserves data
	originalEntity := &entity.Event{
		StreamID:      "stream-123",
		StreamName:    "TestStream",
		StreamVersion: "1",
		EventID:       "event-456",
		EventName:     "TestEvent",
		EventData:     `{"key": "value"}`,
		OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	// When: entity -> domain
	domainEvent, err := ToDomain(originalEntity)
	assert.NoError(t, err)
	assert.NotNil(t, domainEvent)

	// When: domain -> entity
	resultEntity, err := ToPersistence(domainEvent)
	assert.NoError(t, err)
	assert.NotNil(t, resultEntity)

	// Then: check that all fields are preserved
	assert.Equal(t, originalEntity.StreamID, resultEntity.StreamID)
	assert.Equal(t, originalEntity.StreamName, resultEntity.StreamName)
	assert.Equal(t, originalEntity.StreamVersion, resultEntity.StreamVersion)
	assert.Equal(t, originalEntity.EventID, resultEntity.EventID)
	assert.Equal(t, originalEntity.EventName, resultEntity.EventName)
	assert.Equal(t, originalEntity.EventData, resultEntity.EventData)
	assert.Equal(t, originalEntity.OccurredAt, resultEntity.OccurredAt)
}

func TestMapperFunctions_DoNotModifyInput(t *testing.T) {
	// Test that mapper functions don't modify the input structs
	originalEntity := &entity.Event{
		StreamID:      "stream-123",
		StreamName:    "TestStream",
		StreamVersion: "1",
		EventID:       "event-456",
		EventName:     "TestEvent",
		EventData:     `{"key": "value"}`,
		OccurredAt:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	originalDomain := &domain.Event{
		StreamID:      "domain-stream-123",
		StreamName:    "DomainTestStream",
		StreamVersion: "2",
		EventID:       "domain-event-456",
		EventName:     "DomainTestEvent",
		EventData:     `{"domain": "value"}`,
		OccurredAt:    time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
	}

	// Store original values
	entityBefore := *originalEntity
	domainBefore := *originalDomain

	// Perform conversions
	_, _ = ToDomain(originalEntity)
	_, _ = ToPersistence(originalDomain)

	// Check that originals weren't modified
	assert.Equal(t, entityBefore, *originalEntity)
	assert.Equal(t, domainBefore, *originalDomain)
}

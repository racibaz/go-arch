package logging

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/racibaz/go-arch/pkg/ddd"
)

// Mock event for testing
type mockEvent struct {
	ddd.Entity
	payload    ddd.EventPayload
	metadata   ddd.Metadata
	occurredAt time.Time
}

func (e mockEvent) EventName() string         { return e.EntityName() }
func (e mockEvent) Payload() ddd.EventPayload { return e.payload }
func (e mockEvent) Metadata() ddd.Metadata    { return e.metadata }
func (e mockEvent) OccurredAt() time.Time     { return e.occurredAt }

func newMockEvent() mockEvent {
	return mockEvent{
		Entity:     ddd.NewEntity(uuid.New().String(), "MockEvent"),
		payload:    nil,
		metadata:   make(ddd.Metadata),
		occurredAt: time.Now(),
	}
}

// Mock event handler for testing
type mockEventHandler struct {
	shouldError bool
}

func (h *mockEventHandler) HandleEvent(ctx context.Context, event ddd.Event) error {
	if h.shouldError {
		return errors.New("mock handler error")
	}
	return nil
}

// Mock logger for testing
type mockLogger struct {
	loggedMessages []string
}

func (m *mockLogger) Debug(msg string, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, msg)
}

func (m *mockLogger) Info(msg string, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, msg)
}

func (m *mockLogger) Warn(msg string, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, msg)
}

func (m *mockLogger) Error(msg string, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, msg)
}

func (m *mockLogger) Fatal(msg string, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, msg)
}

func TestLogEventHandlerAccess(t *testing.T) {
	mockLogger := &mockLogger{}
	handler := &mockEventHandler{shouldError: false}

	loggedHandler := LogEventHandlerAccess[ddd.Event](handler, "TestHandler", mockLogger)

	if loggedHandler.label != "TestHandler" {
		t.Errorf("Expected label 'TestHandler', got '%s'", loggedHandler.label)
	}
}

func TestEventHandlers_HandleEvent_Success(t *testing.T) {
	mockLogger := &mockLogger{}
	handler := &mockEventHandler{shouldError: false}

	loggedHandler := LogEventHandlerAccess[ddd.Event](handler, "TestHandler", mockLogger)

	event := newMockEvent()
	err := loggedHandler.HandleEvent(context.Background(), event)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that info log was called (format string)
	foundInfo := false
	for _, msg := range mockLogger.loggedMessages {
		if msg == "--> Post.%s.On(%s)" {
			foundInfo = true
			break
		}
	}
	if !foundInfo {
		t.Errorf("Expected info log message not found. Got: %v", mockLogger.loggedMessages)
	}

	// Check that error log was called (deferred, format string)
	foundError := false
	for _, msg := range mockLogger.loggedMessages {
		if msg == "<-- Post.%s.On(%s)" {
			foundError = true
			break
		}
	}
	if !foundError {
		t.Errorf("Expected error log message not found. Got: %v", mockLogger.loggedMessages)
	}
}

func TestEventHandlers_HandleEvent_Error(t *testing.T) {
	mockLogger := &mockLogger{}
	handler := &mockEventHandler{shouldError: true}

	loggedHandler := LogEventHandlerAccess[ddd.Event](handler, "TestHandler", mockLogger)

	event := newMockEvent()
	err := loggedHandler.HandleEvent(context.Background(), event)

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err.Error() != "mock handler error" {
		t.Errorf("Expected 'mock handler error', got '%s'", err.Error())
	}

	// Check that logs were still called
	foundInfo := false
	foundError := false
	for _, msg := range mockLogger.loggedMessages {
		if msg == "--> Post.%s.On(%s)" {
			foundInfo = true
		}
		if msg == "<-- Post.%s.On(%s)" {
			foundError = true
		}
	}
	if !foundInfo {
		t.Errorf("Expected info log message not found. Got: %v", mockLogger.loggedMessages)
	}
	if !foundError {
		t.Errorf("Expected error log message not found. Got: %v", mockLogger.loggedMessages)
	}
}

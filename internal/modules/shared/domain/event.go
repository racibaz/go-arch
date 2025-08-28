package domain

import "time"

type Event struct {
	StreamID      string
	StreamName    string
	StreamVersion string
	EventID       string
	EventName     string
	EventData     string
	OccurredAt    time.Time
}

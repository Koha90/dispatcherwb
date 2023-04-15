// Package events keep types of events and ...
package events

// Fetcher interface for fetcher
type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

// Processor interface for processor
type Processor interface {
	Process(e Event) error
}

// Type users type for event. Is int
type Type int

// Unknown, Message constants for Event
const (
	Unknown Type = iota
	Message
)

// Event struct keep field Type, Text and Meta
type Event struct {
	Type Type
	Text string
	Meta interface{}
}

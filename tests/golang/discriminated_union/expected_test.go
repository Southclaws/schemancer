package discriminated_union_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type BaseEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

type EventUnion interface {
	EventType() string
	isEvent()
}

type Event struct {
	EventUnion
}

func (w Event) MarshalJSON() ([]byte, error) {
	if w.EventUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.EventUnion)
}

func (w *Event) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.EventUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("Event: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("Event: missing discriminator field %q", "type")
	}

	var v EventUnion
	switch peek.Type {
	case "created":
		v = &CreatedEvent{}
	case "updated":
		v = &UpdatedEvent{}
	case "deleted":
		v = &DeletedEvent{}
	default:
		return fmt.Errorf("Event: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("Event: invalid %q payload: %w", peek.Type, err)
	}

	w.EventUnion = v
	return nil
}

type CreatedEvent struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

func (CreatedEvent) isEvent() {}

func (CreatedEvent) EventType() string { return "created" }

type UpdatedEvent struct {
	Changes   map[string]interface{} `json:"changes"`
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	Type      string                 `json:"type"`
}

func (UpdatedEvent) isEvent() {}

func (UpdatedEvent) EventType() string { return "updated" }

type DeletedEvent struct {
	ID        string    `json:"id"`
	Reason    *string   `json:"reason,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

func (DeletedEvent) isEvent() {}

func (DeletedEvent) EventType() string { return "deleted" }

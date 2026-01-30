package external_refs

import (
	"github.com/google/uuid"
	"time"
)

type EventMetadata struct {
	Source    *string   `json:"source,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type EventPayload struct {
	Data      map[string]interface{} `json:"data,omitempty"`
	EventID   uuid.UUID              `json:"eventId"`
	EventType string                 `json:"eventType"`
}

type EventHandler struct {
	HandlerID       string         `json:"handlerId"`
	Metadata        *EventMetadata `json:"metadata,omitempty"`
	Payload         EventPayload   `json:"payload"`
	RelatedPayloads []EventPayload `json:"relatedPayloads,omitempty"`
}

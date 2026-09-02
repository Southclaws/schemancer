package union_with_refs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Command struct {
	Name string `json:"name"`
}

type Status struct {
	Error   *string `json:"error,omitempty"`
	Success bool    `json:"success"`
}

type MessagesUnion interface {
	MessagesType() string
	isMessages()
}

type Messages struct {
	MessagesUnion
}

func (w Messages) MarshalJSON() ([]byte, error) {
	if w.MessagesUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.MessagesUnion)
}

func (w *Messages) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.MessagesUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("Messages: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("Messages: missing discriminator field %q", "type")
	}

	var v MessagesUnion
	switch peek.Type {
	case "control":
		v = &ControlMessage{}
	case "status":
		v = &StatusMessage{}
	default:
		return fmt.Errorf("Messages: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("Messages: invalid %q payload: %w", peek.Type, err)
	}

	w.MessagesUnion = v
	return nil
}

type ControlMessage struct {
	Data Command `json:"data"`
	Type string  `json:"type"`
}

func (ControlMessage) isMessages() {}

func (ControlMessage) MessagesType() string { return "control" }

type StatusMessage struct {
	Data *Status `json:"data,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (StatusMessage) isMessages() {}

func (StatusMessage) MessagesType() string { return "status" }

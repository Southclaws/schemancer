package external_refs_only_defs_test

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type EventType string

const (
	EventTypeThreadPublished   EventType = "thread_published"
	EventTypeThreadUnpublished EventType = "thread_unpublished"
	EventTypeThreadUpdated     EventType = "thread_updated"
)

type EventPayload struct {
	Data      map[string]interface{} `json:"data"`
	EventType EventType              `json:"event_type"`
}

type RPCRequestBase struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type RPCRequestOtherParams struct {
	Info *string `json:"info,omitempty"`
}

type RPCRequestToPluginUnion interface {
	RPCRequestToPluginType() string
	isRPCRequestToPlugin()
}

type RPCRequestToPlugin struct {
	RPCRequestToPluginUnion
}

func (w RPCRequestToPlugin) MarshalJSON() ([]byte, error) {
	if w.RPCRequestToPluginUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.RPCRequestToPluginUnion)
}

func (w *RPCRequestToPlugin) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.RPCRequestToPluginUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"method"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("RPCRequestToPlugin: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("RPCRequestToPlugin: missing discriminator field %q", "method")
	}

	var v RPCRequestToPluginUnion
	switch peek.Type {
	case "event":
		v = &RPCRequestEvent{}
	case "other":
		v = &RPCRequestOther{}
	default:
		return fmt.Errorf("RPCRequestToPlugin: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("RPCRequestToPlugin: invalid %q payload: %w", peek.Type, err)
	}

	w.RPCRequestToPluginUnion = v
	return nil
}

type RPCRequestEvent struct {
	ID      int          `json:"id"`
	Jsonrpc string       `json:"jsonrpc"`
	Method  string       `json:"method"`
	Params  EventPayload `json:"params"`
}

func (RPCRequestEvent) isRPCRequestToPlugin() {}

func (RPCRequestEvent) RPCRequestToPluginType() string { return "event" }

type RPCRequestOther struct {
	ID      int                   `json:"id"`
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  RPCRequestOtherParams `json:"params"`
}

func (RPCRequestOther) isRPCRequestToPlugin() {}

func (RPCRequestOther) RPCRequestToPluginType() string { return "other" }

package jsonrpc_test

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RPCRequestBase struct {
	ID      int                    `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCRequestGetConfigParams struct {
	// Specific config keys to retrieve. If empty, returns all config.
	Keys []string `json:"keys,omitempty"`
}

type RPCRequestPingParams struct {
}

type RPCRequestPublishEventParams struct {
	// Event type to publish (must be in snake_case)
	EventType string `json:"event_type"`
	// Event data
	Payload map[string]interface{} `json:"payload"`
}

type RPCRequestToHostUnion interface {
	RPCRequestToHostType() string
	isRPCRequestToHost()
}

type RPCRequestToHost struct {
	RPCRequestToHostUnion
}

func (w RPCRequestToHost) MarshalJSON() ([]byte, error) {
	if w.RPCRequestToHostUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.RPCRequestToHostUnion)
}

func (w *RPCRequestToHost) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.RPCRequestToHostUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"method"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("RPCRequestToHost: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("RPCRequestToHost: missing discriminator field %q", "method")
	}

	var v RPCRequestToHostUnion
	switch peek.Type {
	case "get_config":
		v = &RPCRequestGetConfig{}
	case "publish_event":
		v = &RPCRequestPublishEvent{}
	default:
		return fmt.Errorf("RPCRequestToHost: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("RPCRequestToHost: invalid %q payload: %w", peek.Type, err)
	}

	w.RPCRequestToHostUnion = v
	return nil
}

type RPCRequestGetConfig struct {
	ID      int                        `json:"id"`
	Jsonrpc string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  *RPCRequestGetConfigParams `json:"params,omitempty"`
}

func (RPCRequestGetConfig) isRPCRequestToHost() {}

func (RPCRequestGetConfig) RPCRequestToHostType() string { return "get_config" }

type RPCRequestPublishEvent struct {
	ID      int                          `json:"id"`
	Jsonrpc string                       `json:"jsonrpc"`
	Method  string                       `json:"method"`
	Params  RPCRequestPublishEventParams `json:"params"`
}

func (RPCRequestPublishEvent) isRPCRequestToHost() {}

func (RPCRequestPublishEvent) RPCRequestToHostType() string { return "publish_event" }

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
	case "ping":
		v = &RPCRequestPing{}
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
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	// Event payload - structure depends on event type
	Params map[string]interface{} `json:"params"`
}

func (RPCRequestEvent) isRPCRequestToPlugin() {}

func (RPCRequestEvent) RPCRequestToPluginType() string { return "event" }

type RPCRequestPing struct {
	ID      int                   `json:"id"`
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  *RPCRequestPingParams `json:"params,omitempty"`
}

func (RPCRequestPing) isRPCRequestToPlugin() {}

func (RPCRequestPing) RPCRequestToPluginType() string { return "ping" }

type RPCResponseBaseError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponseBase struct {
	Error   *RPCResponseBaseError  `json:"error,omitempty"`
	ID      int                    `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Result  map[string]interface{} `json:"result,omitempty"`
}

type RPCResponseEventError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponseEventResult struct {
	Ok *bool `json:"ok,omitempty"`
}

type RPCResponseEvent struct {
	Error   *RPCResponseEventError  `json:"error,omitempty"`
	ID      int                     `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Result  *RPCResponseEventResult `json:"result,omitempty"`
}

type RPCResponseGetConfigError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponseGetConfigResult struct {
	// Configuration key-value pairs
	Config map[string]interface{} `json:"config,omitempty"`
}

type RPCResponseGetConfig struct {
	Error   *RPCResponseGetConfigError  `json:"error,omitempty"`
	ID      int                         `json:"id"`
	Jsonrpc string                      `json:"jsonrpc"`
	Result  *RPCResponseGetConfigResult `json:"result,omitempty"`
}

type RPCResponsePublishEventError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponsePublishEventResult struct {
	Ok *bool `json:"ok,omitempty"`
}

type RPCResponsePublishEvent struct {
	Error   *RPCResponsePublishEventError  `json:"error,omitempty"`
	ID      int                            `json:"id"`
	Jsonrpc string                         `json:"jsonrpc"`
	Result  *RPCResponsePublishEventResult `json:"result,omitempty"`
}

type RPCResponseFromHost = interface{}

type RPCResponsePingError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponsePingResult struct {
	Pong bool `json:"pong"`
	// Optional status message
	Status *string `json:"status,omitempty"`
	// How long the plugin has been running
	UptimeSeconds *float64 `json:"uptime_seconds,omitempty"`
}

type RPCResponsePing struct {
	Error   *RPCResponsePingError  `json:"error,omitempty"`
	ID      int                    `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Result  *RPCResponsePingResult `json:"result,omitempty"`
}

type RPCResponseFromPlugin = interface{}

type RPCResponseShutdownError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type RPCResponseShutdownResult struct {
	Ok *bool `json:"ok,omitempty"`
}

type RPCResponseShutdown struct {
	Error   *RPCResponseShutdownError  `json:"error,omitempty"`
	ID      int                        `json:"id"`
	Jsonrpc string                     `json:"jsonrpc"`
	Result  *RPCResponseShutdownResult `json:"result,omitempty"`
}

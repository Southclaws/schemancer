package composed_union

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type PluginConfigurationFieldUnion interface {
	PluginConfigurationFieldType() string
	isPluginConfigurationField()
}

type PluginConfigurationField struct {
	PluginConfigurationFieldUnion
}

func (w PluginConfigurationField) MarshalJSON() ([]byte, error) {
	if w.PluginConfigurationFieldUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.PluginConfigurationFieldUnion)
}

func (w *PluginConfigurationField) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.PluginConfigurationFieldUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("PluginConfigurationField: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("PluginConfigurationField: missing discriminator field %q", "type")
	}

	var v PluginConfigurationFieldUnion
	switch peek.Type {
	case "string":
		v = &PluginConfigurationFieldString{}
	case "number":
		v = &PluginConfigurationFieldNumber{}
	case "boolean":
		v = &PluginConfigurationFieldBoolean{}
	default:
		return fmt.Errorf("PluginConfigurationField: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("PluginConfigurationField: invalid %q payload: %w", peek.Type, err)
	}

	w.PluginConfigurationFieldUnion = v
	return nil
}

type PluginConfigurationFieldString struct {
	Default *string `json:"default,omitempty"`
	// A description of the configuration field.
	Description *string `json:"description,omitempty"`
	// A unique identifier for this configuration field, used for
	// referencing the field in the plugin configuration object.
	//
	ID *string `json:"id,omitempty"`
	// A human-readable label for the configuration field.
	Label *string `json:"label,omitempty"`
	Type  string  `json:"type"`
}

func (PluginConfigurationFieldString) isPluginConfigurationField() {}

func (PluginConfigurationFieldString) PluginConfigurationFieldType() string { return "string" }

type PluginConfigurationFieldNumber struct {
	Default *float64 `json:"default,omitempty"`
	// A description of the configuration field.
	Description *string `json:"description,omitempty"`
	// A unique identifier for this configuration field, used for
	// referencing the field in the plugin configuration object.
	//
	ID *string `json:"id,omitempty"`
	// A human-readable label for the configuration field.
	Label *string `json:"label,omitempty"`
	Type  string  `json:"type"`
}

func (PluginConfigurationFieldNumber) isPluginConfigurationField() {}

func (PluginConfigurationFieldNumber) PluginConfigurationFieldType() string { return "number" }

type PluginConfigurationFieldBoolean struct {
	Default *bool `json:"default,omitempty"`
	// A description of the configuration field.
	Description *string `json:"description,omitempty"`
	// A unique identifier for this configuration field, used for
	// referencing the field in the plugin configuration object.
	//
	ID *string `json:"id,omitempty"`
	// A human-readable label for the configuration field.
	Label *string `json:"label,omitempty"`
	Type  string  `json:"type"`
}

func (PluginConfigurationFieldBoolean) isPluginConfigurationField() {}

func (PluginConfigurationFieldBoolean) PluginConfigurationFieldType() string { return "boolean" }

type PluginConfigurationFieldSchema = PluginConfigurationField

type PluginConfigurationSchema struct {
	Fields []PluginConfigurationFieldSchema `json:"fields,omitempty"`
}

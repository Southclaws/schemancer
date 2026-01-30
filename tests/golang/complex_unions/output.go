package complex_unions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AllOfComposition struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Timestamp *string `json:"timestamp,omitempty"`
}

type MultiType = interface{}

type NestedUnion struct {
	Data interface{} `json:"data,omitempty"`
}

// Test complex anyOf/oneOf scenarios
type ObjectUnionUnion interface {
	ObjectUnionType() string
	isObjectUnion()
}

type ObjectUnion struct {
	ObjectUnionUnion
}

func (w ObjectUnion) MarshalJSON() ([]byte, error) {
	if w.ObjectUnionUnion == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.ObjectUnionUnion)
}

func (w *ObjectUnion) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.ObjectUnionUnion = nil
		return nil
	}

	var peek struct {
		Type string `json:"kind"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("ObjectUnion: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("ObjectUnion: missing discriminator field %q", "kind")
	}

	var v ObjectUnionUnion
	switch peek.Type {
	case "a":
		v = &ObjectUnionA{}
	case "b":
		v = &ObjectUnionB{}
	case "c":
		v = &ObjectUnionC{}
	default:
		return fmt.Errorf("ObjectUnion: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("ObjectUnion: invalid %q payload: %w", peek.Type, err)
	}

	w.ObjectUnionUnion = v
	return nil
}

type ObjectUnionA struct {
	AField string `json:"aField"`
	Kind   string `json:"kind"`
}

func (ObjectUnionA) isObjectUnion() {}

func (ObjectUnionA) ObjectUnionType() string { return "a" }

type ObjectUnionB struct {
	BField int    `json:"bField"`
	Kind   string `json:"kind"`
}

func (ObjectUnionB) isObjectUnion() {}

func (ObjectUnionB) ObjectUnionType() string { return "b" }

type ObjectUnionC struct {
	CField bool   `json:"cField"`
	Kind   string `json:"kind"`
}

func (ObjectUnionC) isObjectUnion() {}

func (ObjectUnionC) ObjectUnionType() string { return "c" }

type StringOrNull = interface{}

type StringOrNumber = interface{}

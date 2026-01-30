package loader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Southclaws/schemancer/schemancer/deref"
	"github.com/goccy/go-yaml"
	"github.com/google/jsonschema-go/jsonschema"
)

func FromFile(filename string) (*jsonschema.Schema, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	schema, err := FromReader(f)
	if err != nil {
		return nil, err
	}

	// Resolve external $ref references
	baseDir := filepath.Dir(filename)
	if err := deref.Schema(schema, baseDir); err != nil {
		return nil, fmt.Errorf("failed to dereference schema: %w", err)
	}

	return schema, nil
}

func FromReader(r io.Reader) (*jsonschema.Schema, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	var yamlData interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	jsonData, err := json.Marshal(yamlData)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to JSON: %w", err)
	}

	var schema jsonschema.Schema
	if err := json.Unmarshal(jsonData, &schema); err != nil {
		return nil, fmt.Errorf("failed to unmarshal schema: %w", err)
	}

	return &schema, nil
}

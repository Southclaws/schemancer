// Package deref provides utilities for dereferencing JSON Schema files.
// It resolves external $ref references and hoists nested definitions to the root.
package deref

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/google/jsonschema-go/jsonschema"
)

// Schema dereferences all external $ref references in a schema in-place.
// baseDir is used to resolve relative file paths in $ref values.
func Schema(schema *jsonschema.Schema, baseDir string) error {
	rootDefs := make(map[string]*jsonschema.Schema)

	// Copy existing definitions
	for k, v := range schema.Defs {
		rootDefs[k] = v
	}
	for k, v := range schema.Definitions {
		rootDefs[k] = v
	}

	if err := dereferenceRefs(schema, baseDir, rootDefs); err != nil {
		return fmt.Errorf("dereference schema: %w", err)
	}

	// Hoist all definitions to $defs
	if len(rootDefs) > 0 {
		if schema.Defs == nil {
			schema.Defs = make(map[string]*jsonschema.Schema)
		}
		for k, v := range rootDefs {
			schema.Defs[k] = v
		}
	}

	return nil
}

func dereferenceRefs(schema *jsonschema.Schema, baseDir string, rootDefs map[string]*jsonschema.Schema) error {
	if schema == nil {
		return nil
	}

	// Handle external $ref (any ref that doesn't start with # is potentially external)
	if schema.Ref != "" && !strings.HasPrefix(schema.Ref, "#") {
		refPath := schema.Ref
		var jsonPointer string

		// Split file path from JSON Pointer fragment (e.g., "./events.yaml#/$defs/EventPayload")
		if idx := strings.Index(refPath, "#"); idx != -1 {
			jsonPointer = refPath[idx+1:] // Get everything after #
			refPath = refPath[:idx]       // Get file path before #
		}

		fullPath := filepath.Join(baseDir, refPath)

		refData, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("read ref %s: %w", schema.Ref, err)
		}

		// Use YAML→JSON conversion to trigger jsonschema's custom unmarshaler
		var yamlData any
		if err := yaml.Unmarshal(refData, &yamlData); err != nil {
			return fmt.Errorf("parse ref %s: %w", schema.Ref, err)
		}
		jsonData, err := json.Marshal(yamlData)
		if err != nil {
			return fmt.Errorf("convert ref %s: %w", schema.Ref, err)
		}
		var refSchema jsonschema.Schema
		if err := json.Unmarshal(jsonData, &refSchema); err != nil {
			return fmt.Errorf("unmarshal ref %s: %w", schema.Ref, err)
		}

		refBaseDir := filepath.Dir(fullPath)
		if err := dereferenceRefs(&refSchema, refBaseDir, rootDefs); err != nil {
			return err
		}

		// If there's a JSON Pointer, resolve it and add to root defs
		if jsonPointer != "" {
			targetSchema, err := resolveJSONPointer(&refSchema, jsonPointer)
			if err != nil {
				return fmt.Errorf("resolve pointer %s in %s: %w", jsonPointer, schema.Ref, err)
			}

			// Extract the definition name from the pointer
			parts := strings.Split(strings.TrimPrefix(jsonPointer, "/"), "/")
			var defName string

			if len(parts) >= 2 && (parts[0] == "$defs" || parts[0] == "definitions") {
				// Standard $defs/EventPayload format
				defName = parts[1]
			} else if len(parts) == 1 && parts[0] != "" {
				// Root-level property like EventPayload (stored in Extra)
				defName = parts[0]
			}

			if defName != "" {
				// Add to root defs if not already present
				if _, exists := rootDefs[defName]; !exists {
					rootDefs[defName] = targetSchema
				}

				// Convert external ref to internal ref
				schema.Ref = "#/$defs/" + defName
			} else {
				// Fallback: inline the schema if we can't determine a name
				inlineSchema(schema, targetSchema)
			}
		} else {
			// No JSON Pointer: inline the entire file
			inlineSchema(schema, &refSchema)
		}
	}

	// Recursively process all nested schemas
	if err := dereferenceRefs(schema.Items, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.AdditionalItems, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.AdditionalProperties, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.Contains, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.PropertyNames, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.UnevaluatedItems, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.UnevaluatedProperties, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.If, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.Then, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.Else, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.Not, baseDir, rootDefs); err != nil {
		return err
	}
	if err := dereferenceRefs(schema.ContentSchema, baseDir, rootDefs); err != nil {
		return err
	}

	// Process schema arrays
	for _, s := range schema.PrefixItems {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.ItemsArray {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.AllOf {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.AnyOf {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.OneOf {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}

	// Process schema maps
	for _, s := range schema.Properties {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.PatternProperties {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.DependentSchemas {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}
	for _, s := range schema.DependencySchemas {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
	}

	// Process and hoist nested definitions
	for name, s := range schema.Defs {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
		if _, exists := rootDefs[name]; !exists {
			rootDefs[name] = s
		}
	}
	for name, s := range schema.Definitions {
		if err := dereferenceRefs(s, baseDir, rootDefs); err != nil {
			return err
		}
		if _, exists := rootDefs[name]; !exists {
			rootDefs[name] = s
		}
	}

	return nil
}

// inlineSchema merges the fields from src into dst, clearing the $ref.
func inlineSchema(dst, src *jsonschema.Schema) {
	// Clear the ref since we're inlining
	dst.Ref = ""

	// Copy metadata if not set
	if dst.Title == "" {
		dst.Title = src.Title
	}
	if dst.Description == "" {
		dst.Description = src.Description
	}
	if dst.Type == "" {
		dst.Type = src.Type
	}
	if dst.Types == nil {
		dst.Types = src.Types
	}
	if dst.Format == "" {
		dst.Format = src.Format
	}

	// Copy validation
	if dst.Enum == nil {
		dst.Enum = src.Enum
	}
	if dst.Const == nil {
		dst.Const = src.Const
	}
	if dst.Required == nil {
		dst.Required = src.Required
	}

	// Copy nested schemas
	if dst.Properties == nil {
		dst.Properties = src.Properties
	}
	if dst.Items == nil {
		dst.Items = src.Items
	}
	if dst.AdditionalProperties == nil {
		dst.AdditionalProperties = src.AdditionalProperties
	}
	if dst.AllOf == nil {
		dst.AllOf = src.AllOf
	}
	if dst.AnyOf == nil {
		dst.AnyOf = src.AnyOf
	}
	if dst.OneOf == nil {
		dst.OneOf = src.OneOf
	}

	// Copy definitions
	if dst.Defs == nil {
		dst.Defs = src.Defs
	}
	if dst.Definitions == nil {
		dst.Definitions = src.Definitions
	}
}

// resolveJSONPointer resolves a JSON Pointer within a schema.
// For example, "/$defs/EventPayload" returns the EventPayload definition.
// Also handles root-level properties like "/EventPayload" stored in schema.Extra.
func resolveJSONPointer(schema *jsonschema.Schema, pointer string) (*jsonschema.Schema, error) {
	// Remove leading slash if present
	pointer = strings.TrimPrefix(pointer, "/")

	if pointer == "" {
		return schema, nil
	}

	// Split the pointer into parts
	parts := strings.Split(pointer, "/")

	// Handle $defs and definitions
	if parts[0] == "$defs" || parts[0] == "definitions" {
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid JSON Pointer: %s", pointer)
		}

		defName := parts[1]
		var def *jsonschema.Schema

		if parts[0] == "$defs" && schema.Defs != nil {
			def = schema.Defs[defName]
		} else if parts[0] == "definitions" && schema.Definitions != nil {
			def = schema.Definitions[defName]
		}

		if def == nil {
			return nil, fmt.Errorf("definition not found: %s", defName)
		}

		// If there are more parts, continue resolving
		if len(parts) > 2 {
			remainingPointer := strings.Join(parts[2:], "/")
			return resolveJSONPointer(def, remainingPointer)
		}

		return def, nil
	}

	// Handle root-level properties (stored in schema.Extra)
	if schema.Extra != nil {
		if extraValue, ok := schema.Extra[parts[0]]; ok {
			// Parse the Extra value into a schema
			extraSchema := parseExtraSchema(extraValue)
			if extraSchema == nil {
				return nil, fmt.Errorf("failed to parse root-level property: %s", parts[0])
			}

			// If there are more parts, continue resolving
			if len(parts) > 1 {
				remainingPointer := strings.Join(parts[1:], "/")
				return resolveJSONPointer(extraSchema, remainingPointer)
			}

			return extraSchema, nil
		}
	}

	return nil, fmt.Errorf("unsupported JSON Pointer: %s", pointer)
}

// parseExtraSchema converts a value from Schema.Extra into a jsonschema.Schema.
// Extra values are stored as map[string]any, so we need to marshal/unmarshal to get a proper Schema.
func parseExtraSchema(v any) *jsonschema.Schema {
	if v == nil {
		return nil
	}

	// Marshal to JSON and unmarshal into a Schema
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	var schema jsonschema.Schema
	if err := json.Unmarshal(jsonBytes, &schema); err != nil {
		return nil
	}

	return &schema
}

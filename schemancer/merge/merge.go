// Package merge provides utilities for merging JSON Schema definitions.
// This is used for handling allOf composition and resolving references.
package merge

import (
	"encoding/json"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
)

// ResolveSchema resolves a $ref reference within a root schema.
// If the schema has no $ref, returns the schema as-is.
// Returns nil only if the reference cannot be resolved.
func ResolveSchema(root *jsonschema.Schema, s *jsonschema.Schema) *jsonschema.Schema {
	if s == nil {
		return nil
	}

	// If no ref, return the schema itself
	if s.Ref == "" {
		return s
	}

	// Only handle local references (#/$defs/... or #/Name)
	if len(s.Ref) > 0 && s.Ref[0] == '#' {
		refPath := s.Ref[1:]

		// Handle $defs references (#/$defs/...)
		if len(refPath) > 7 && refPath[:7] == "/$defs/" {
			defName := refPath[7:]
			// Handle nested $defs (e.g., #/$defs/A/$defs/B)
			if idx := strings.Index(defName, "/"); idx != -1 {
				// For now, we only support top-level $defs
				defName = defName[:idx]
			}
			if root.Defs != nil {
				if defSchema, exists := root.Defs[defName]; exists {
					return defSchema
				}
			}
		} else if len(refPath) > 1 && refPath[0] == '/' {
			// Handle root-level references (#/Name) - look in Extra
			defName := refPath[1:]
			if root.Extra != nil {
				if extraVal, exists := root.Extra[defName]; exists {
					// Parse the Extra value into a Schema
					return parseExtraSchemaHelper(extraVal)
				}
			}
		}
	}
	return nil
}

// parseExtraSchemaHelper converts an Extra value to a jsonschema.Schema
func parseExtraSchemaHelper(v any) *jsonschema.Schema {
	if v == nil {
		return nil
	}

	// Extra values are typically map[string]any, marshal/unmarshal to get a proper Schema
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

// AllOf merges all schemas in an allOf array into a single schema.
// This handles nested allOf structures and $ref resolution recursively.
// Returns a new schema with merged properties and required fields.
func AllOf(root *jsonschema.Schema, s *jsonschema.Schema) *jsonschema.Schema {
	if s == nil {
		return nil
	}

	if len(s.AllOf) == 0 {
		return s
	}

	merged := &jsonschema.Schema{
		Type:        "object",
		Properties:  make(map[string]*jsonschema.Schema),
		Required:    []string{},
		Description: s.Description,
	}

	// First, include any properties directly on the schema
	if s.Properties != nil {
		for k, v := range s.Properties {
			merged.Properties[k] = v
		}
	}
	merged.Required = append(merged.Required, s.Required...)

	// Then merge all allOf parts
	for _, part := range s.AllOf {
		resolved := part
		if part.Ref != "" {
			resolved = ResolveSchema(root, part)
			if resolved == nil {
				resolved = part
			}
		}

		// Recursively merge nested allOf
		sub := AllOf(root, resolved)
		if sub == nil {
			continue
		}

		// Merge properties (last definition wins, allowing overrides like const)
		if sub.Properties != nil {
			for k, v := range sub.Properties {
				merged.Properties[k] = v
			}
		}

		// Union required arrays
		merged.Required = append(merged.Required, sub.Required...)
	}

	// Deduplicate required
	merged.Required = uniqueStrings(merged.Required)

	return merged
}

// uniqueStrings returns a new slice with duplicate strings removed.
func uniqueStrings(s []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(s))
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

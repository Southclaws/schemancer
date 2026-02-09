package detect

import (
	"fmt"
	"strings"

	"github.com/Southclaws/schemancer/schemancer/merge"
	"github.com/google/jsonschema-go/jsonschema"
)

type UnionResult struct {
	DiscriminatorField string
	Variants           []Variant
}

type Variant struct {
	Name       string
	ConstValue string
	Schema     *jsonschema.Schema
	BaseType   string // Name of the base type from allOf $ref (if any)
}

func DiscriminatedUnion(schema *jsonschema.Schema) (*UnionResult, error) {
	if len(schema.OneOf) == 0 {
		return nil, nil
	}

	// First pass: collect all merged variants and their const fields
	type variantInfo struct {
		name        string
		schema      *jsonschema.Schema
		constFields map[string]string // field name -> const value
		baseType    string            // base type name from allOf $ref
	}

	variantInfos := []variantInfo{}
	for _, variant := range schema.OneOf {
		resolved := merge.ResolveSchema(schema, variant)
		if resolved == nil {
			continue
		}

		// Extract base type from allOf $ref before merging
		baseType := FindAllOfBaseRef(resolved)

		merged := merge.AllOf(schema, resolved)
		if merged == nil || merged.Properties == nil {
			continue
		}

		variantName := ""
		if variant.Ref != "" {
			parts := strings.Split(variant.Ref, "/")
			variantName = parts[len(parts)-1]
		}

		constFields := make(map[string]string)
		for propName, propSchema := range merged.Properties {
			if propSchema.Const != nil {
				constVal := getConstValue(propSchema.Const)
				constFields[propName] = constVal
			}
		}

		variantInfos = append(variantInfos, variantInfo{
			name:        variantName,
			schema:      merged,
			constFields: constFields,
			baseType:    baseType,
		})
	}

	if len(variantInfos) == 0 {
		return nil, nil
	}

	// Second pass: find a discriminator field where all variants have unique const values
	discriminatorField := ""
	constValues := make(map[string]bool)

	// Get all potential discriminator fields (fields with const in first variant)
	for fieldName := range variantInfos[0].constFields {
		constValues = make(map[string]bool)
		isValidDiscriminator := true

		// Check if all variants have this field with a const value
		for _, vi := range variantInfos {
			constVal, hasConst := vi.constFields[fieldName]
			if !hasConst {
				isValidDiscriminator = false
				break
			}

			// Check if this const value is unique across variants
			if constValues[constVal] {
				isValidDiscriminator = false
				break
			}
			constValues[constVal] = true
		}

		// If this field has unique const values across all variants, use it as discriminator
		if isValidDiscriminator && len(constValues) == len(variantInfos) {
			discriminatorField = fieldName
			break
		}
	}

	if discriminatorField == "" {
		return nil, nil
	}

	// Build final variants using the identified discriminator
	variants := make([]Variant, 0, len(variantInfos))
	for _, vi := range variantInfos {
		variants = append(variants, Variant{
			Name:       vi.name,
			ConstValue: vi.constFields[discriminatorField],
			Schema:     vi.schema,
			BaseType:   vi.baseType,
		})
	}

	return &UnionResult{
		DiscriminatorField: discriminatorField,
		Variants:           variants,
	}, nil
}

// FindAllOfBaseRef extracts the base type name from allOf $ref composition.
// Returns the type name (e.g., "BaseField") if exactly one $ref is found in allOf, else "".
func FindAllOfBaseRef(schema *jsonschema.Schema) string {
	if schema == nil || len(schema.AllOf) == 0 {
		return ""
	}
	for _, part := range schema.AllOf {
		if part.Ref != "" {
			parts := strings.Split(part.Ref, "/")
			return parts[len(parts)-1]
		}
	}
	return ""
}

func ResolveVariantSchema(rootSchema *jsonschema.Schema, variant *jsonschema.Schema) *jsonschema.Schema {
	resolved := merge.ResolveSchema(rootSchema, variant)
	if resolved == nil {
		return variant
	}
	return merge.AllOf(rootSchema, resolved)
}

func ValidateEvent(schema *jsonschema.Schema, event map[string]interface{}) error {
	resolved, err := schema.Resolve(nil)
	if err != nil {
		return fmt.Errorf("failed to resolve schema: %w", err)
	}

	if err := resolved.Validate(event); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

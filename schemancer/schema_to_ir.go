package schemancer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/Southclaws/schemancer/schemancer/detect"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
	"github.com/Southclaws/schemancer/schemancer/merge"

	"github.com/google/jsonschema-go/jsonschema"
)

func SchemaToIR(schema *jsonschema.Schema) (*ir.IR, error) {
	result := &ir.IR{
		Schema: schema,
		Types:  []ir.IRType{},
	}

	union, _ := detect.DiscriminatedUnion(schema)
	if union != nil {
		unionName := ""
		if schema.Title != "" {
			unionName = schema.Title
		}
		irUnion := convertDiscriminatedUnion(schema, union, unionName, &result.Types)
		result.Types = append(result.Types, irUnion)
		return result, nil
	}

	// Track which $defs names are used in discriminated unions to avoid duplicates
	usedInUnions := make(map[string]bool)

	// Process root-level named schemas from Extra field (e.g., RPCRequestToPlugin)
	if schema.Extra != nil {
		extraNames := make([]string, 0, len(schema.Extra))
		for name := range schema.Extra {
			extraNames = append(extraNames, name)
		}
		sort.Strings(extraNames)

		for _, name := range extraNames {
			extraSchema := parseExtraSchema(schema.Extra[name])
			if extraSchema == nil {
				continue
			}

			// Check if this is a discriminated union
			// We need to create a wrapper schema that has access to root $defs for ref resolution
			wrapperSchema := &jsonschema.Schema{
				OneOf: extraSchema.OneOf,
				Defs:  schema.Defs,  // Use root schema's $defs for ref resolution
				Extra: schema.Extra, // Also include Extra for resolving root-level refs
			}
			union, _ := detect.DiscriminatedUnion(wrapperSchema)
			if union != nil {
				irUnion := convertDiscriminatedUnion(schema, union, name, &result.Types)
				result.Types = append(result.Types, irUnion)

				// Mark all variant names and base types as used in unions
				for _, v := range union.Variants {
					usedInUnions[v.Name] = true
					if v.BaseType != "" {
						usedInUnions[v.BaseType] = true
					}
				}
				continue
			}

			// Otherwise, process as a regular type
			irType := convertSchemaToIRType(schema, name, extraSchema, &result.Types)
			if irType != nil {
				result.Types = append(result.Types, *irType)
			}
		}
	}

	if schema.Defs != nil {
		names := make([]string, 0, len(schema.Defs))
		for name := range schema.Defs {
			names = append(names, name)
		}
		sort.Strings(names)

		// First pass: detect all discriminated unions and mark their variants as used
		for _, name := range names {
			def := schema.Defs[name]

			// Check if this definition is a discriminated union
			wrapperSchema := &jsonschema.Schema{
				OneOf: def.OneOf,
				Defs:  schema.Defs,  // Use root schema's $defs for ref resolution
				Extra: schema.Extra, // Also include Extra for resolving root-level refs
			}
			union, _ := detect.DiscriminatedUnion(wrapperSchema)
			if union != nil {
				irUnion := convertDiscriminatedUnion(schema, union, name, &result.Types)
				result.Types = append(result.Types, irUnion)

				// Mark all variant names and base types as used in unions
				for _, v := range union.Variants {
					usedInUnions[v.Name] = true
					if v.BaseType != "" {
						usedInUnions[v.BaseType] = true
					}
				}
			}
		}

		// First-and-a-half pass: detect allOf schemas that compose a base struct with a
		// discriminated union. E.g.:
		//   PluginConfigurationFieldSchema:
		//     allOf:
		//       - $ref: "#/$defs/PluginConfigurationFieldBase"   # plain struct
		//       - $ref: "#/$defs/PluginConfigurationField"        # discriminated union
		//
		// In this case we merge the base fields into each union variant and make the
		// allOf schema an alias for the union wrapper.
		for _, name := range names {
			if usedInUnions[name] {
				continue
			}
			def := schema.Defs[name]
			if len(def.AllOf) == 0 {
				continue
			}

			// Resolve all allOf members and separate union refs from base refs
			var unionRefName string
			var baseSchemas []*jsonschema.Schema

			for _, part := range def.AllOf {
				if part.Ref == "" {
					continue
				}
				refName := refToTypeName(part.Ref)
				refSchema := merge.ResolveSchema(schema, part)
				if refSchema == nil {
					continue
				}
				// Check whether the resolved ref is a discriminated union
				wrapperSchema := &jsonschema.Schema{
					OneOf: refSchema.OneOf,
					Defs:  schema.Defs,
					Extra: schema.Extra,
				}
				u, _ := detect.DiscriminatedUnion(wrapperSchema)
				if u != nil {
					unionRefName = refName
				} else {
					baseSchemas = append(baseSchemas, refSchema)
				}
			}

			if unionRefName == "" {
				continue
			}

			// Collect fields from all base schemas
			baseFields := collectBaseFields(schema, baseSchemas, &result.Types)

			// Find the already-generated union IR type and merge base fields into its variants
			mergeBaseFieldsIntoUnion(unionRefName, baseFields, result.Types)

			// Mark base schemas as used so they are not emitted as standalone types
			for _, part := range def.AllOf {
				if part.Ref == "" {
					continue
				}
				refName := refToTypeName(part.Ref)
				if refName != unionRefName {
					usedInUnions[refName] = true
				}
			}

			// Make this schema an alias for the union wrapper
			usedInUnions[name] = true
			result.Types = append(result.Types, ir.IRType{
				Name: symbolName(name),
				Kind: ir.IRKindAlias,
				Element: &ir.IRTypeRef{
					Name: unionRefName,
				},
			})
		}

		// Second pass: process remaining types that aren't part of a union
		for _, name := range names {
			// Skip types that are already part of a discriminated union
			if usedInUnions[name] {
				continue
			}

			def := schema.Defs[name]

			// Skip if this was already processed as a union
			if def.OneOf != nil {
				wrapperSchema := &jsonschema.Schema{
					OneOf: def.OneOf,
					Defs:  schema.Defs,
					Extra: schema.Extra,
				}
				union, _ := detect.DiscriminatedUnion(wrapperSchema)
				if union != nil {
					// Already processed in first pass
					continue
				}
			}

			// Process as a regular type
			irType := convertSchemaToIRType(schema, name, def, &result.Types)
			if irType != nil {
				result.Types = append(result.Types, *irType)
			}
		}
	}

	if schema.Title != "" && schema.Type == "object" {
		irType := convertSchemaToIRType(schema, schema.Title, schema, &result.Types)
		if irType != nil {
			result.Types = append(result.Types, *irType)
		}
	}

	// Topologically sort types so dependencies are declared before dependents
	result.Types = topologicalSort(result.Types)

	return result, nil
}

func convertDiscriminatedUnion(root *jsonschema.Schema, union *detect.UnionResult, unionName string, inlineTypes *[]ir.IRType) ir.IRType {
	rootName := symbolName(unionName)
	if rootName == "" {
		rootName = "Union"
		if root.Title != "" {
			rootName = symbolName(root.Title)
		}
	}

	// Track base types that need to be generated
	baseTypesNeeded := make(map[string]bool)
	for _, v := range union.Variants {
		if v.BaseType != "" {
			baseTypesNeeded[v.BaseType] = true
		}
	}

	// Generate base types from $defs if they exist
	for baseTypeName := range baseTypesNeeded {
		if root.Defs != nil {
			if baseSchema, ok := root.Defs[baseTypeName]; ok {
				baseIRType := convertSchemaToIRType(root, baseTypeName, baseSchema, inlineTypes)
				if baseIRType != nil {
					*inlineTypes = append(*inlineTypes, *baseIRType)
				}
			}
		}
	}

	variants := make([]ir.IRVariant, 0, len(union.Variants))
	for _, v := range union.Variants {
		// Generate a name for inline variants using union name + const value
		variantName := v.Name
		if variantName == "" {
			// For inline variants, generate name from const value
			variantName = rootName + symbolName(v.ConstValue)
		}

		variantType := convertStructToIRType(root, variantName, v.Schema, inlineTypes)
		if v.BaseType != "" {
			variantType.BaseType = symbolName(v.BaseType)
		}
		variants = append(variants, ir.IRVariant{
			Name:       symbolName(variantName),
			ConstValue: v.ConstValue,
			Type:       *variantType,
		})
	}

	return ir.IRType{
		Name:        rootName,
		Description: root.Description,
		Kind:        ir.IRKindDiscriminatedUnion,
		Union: &ir.IRDiscriminatedUnion{
			InterfaceName:      rootName + "Union",
			WrapperName:        rootName,
			DiscriminatorField: symbolName(union.DiscriminatorField),
			DiscriminatorJSON:  union.DiscriminatorField,
			Variants:           variants,
		},
	}
}

func convertSchemaToIRType(root *jsonschema.Schema, name string, schema *jsonschema.Schema, inlineTypes *[]ir.IRType) *ir.IRType {
	goName := symbolName(name)

	// Handle allOf composition - merge all schemas into one struct
	if len(schema.AllOf) > 0 {
		baseTypeName := detect.FindAllOfBaseRef(schema)
		merged := merge.AllOf(root, schema)
		if merged != nil && merged.Properties != nil {
			irType := convertStructToIRType(root, name, merged, inlineTypes)
			if baseTypeName != "" {
				irType.BaseType = symbolName(baseTypeName)
			}
			return irType
		}
		// If merge fails, fall through to other handling
	}

	if len(schema.AnyOf) > 0 || len(schema.OneOf) > 0 {
		// Build a non-discriminated union from the variants
		variants := collectUnionVariants(root, schema, goName, inlineTypes)
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindUnion,
			SimpleUnion: &ir.IRUnion{
				Variants: variants,
			},
		}
	}

	if len(schema.Enum) > 0 {
		enumType, enumValues, hasNull := classifyEnum(schema.Enum)

		// Mixed enums (string + int) fall back to any type
		if enumType == ir.IRBuiltinAny {
			return &ir.IRType{
				Name:        goName,
				Description: schema.Description,
				Kind:        ir.IRKindAlias,
				Element: &ir.IRTypeRef{
					Builtin:  ir.IRBuiltinAny,
					Nullable: hasNull,
				},
			}
		}

		// Build backwards-compatible string enum values
		stringValues := make([]string, 0, len(enumValues))
		for _, v := range enumValues {
			if !v.IsNull {
				stringValues = append(stringValues, v.StringValue)
			}
		}

		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindEnum,
			Enum:        stringValues,
			EnumValues:  enumValues,
			EnumType:    enumType,
		}
	}

	if schema.Type == "string" {
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinString,
			},
		}
	}

	if schema.Type == "integer" {
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinInt,
			},
		}
	}

	if schema.Type == "number" {
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinFloat,
			},
		}
	}

	if schema.Type == "boolean" {
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinBool,
			},
		}
	}

	if schema.Type == "array" {
		elem := schemaToIRTypeRefWithContext(root, schema.Items, goName+"Item", inlineTypes)
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Array: &elem,
			},
		}
	}

	if schema.Type == "" && schema.Properties == nil {
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinAny,
			},
		}
	}

	return convertStructToIRType(root, name, schema, inlineTypes)
}

func convertStructToIRType(root *jsonschema.Schema, name string, schema *jsonschema.Schema, inlineTypes *[]ir.IRType) *ir.IRType {
	goName := symbolName(name)

	if schema.Type != "object" && schema.Properties == nil {
		return nil
	}

	if schema.Properties == nil {
		// Check for typed additionalProperties (e.g. map[string]SomeType)
		if schema.AdditionalProperties != nil && (schema.AdditionalProperties.Ref != "" || schema.AdditionalProperties.Type != "") {
			valueRef := schemaToIRTypeRefWithContext(root, schema.AdditionalProperties, goName+"Value", inlineTypes)
			return &ir.IRType{
				Name:        goName,
				Description: schema.Description,
				Kind:        ir.IRKindAlias,
				Element: &ir.IRTypeRef{
					Map: &valueRef,
				},
			}
		}
		return &ir.IRType{
			Name:        goName,
			Description: schema.Description,
			Kind:        ir.IRKindAlias,
			Element: &ir.IRTypeRef{
				Builtin: ir.IRBuiltinAny,
				Map:     &ir.IRTypeRef{Builtin: ir.IRBuiltinAny},
			},
		}
	}

	requiredSet := make(map[string]bool)
	for _, r := range schema.Required {
		requiredSet[r] = true
	}

	propNames := make([]string, 0, len(schema.Properties))
	for propName := range schema.Properties {
		propNames = append(propNames, propName)
	}
	sort.Strings(propNames)

	fields := make([]ir.IRField, 0, len(propNames))
	for _, propName := range propNames {
		propSchema := schema.Properties[propName]
		fieldName := symbolName(propName)
		var fieldDesc string
		if propSchema != nil {
			fieldDesc = propSchema.Description
		}
		field := ir.IRField{
			Name:        fieldName,
			Description: fieldDesc,
			JSONName:    propName,
			Type:        schemaToIRTypeRefWithContext(root, propSchema, goName+fieldName, inlineTypes),
			Required:    requiredSet[propName],
		}

		// Parse default value
		if propSchema != nil && propSchema.Default != nil {
			field.Default = parseDefault(propSchema.Default, &field.Type)
		}

		// Parse language-specific extensions (x-java-name, x-go-name, etc.)
		if propSchema != nil && propSchema.Extra != nil {
			extensions := make(map[string]string)
			for key, val := range propSchema.Extra {
				if strings.HasPrefix(key, "x-") {
					if strVal, ok := val.(string); ok {
						extensions[key] = strVal
					}
				}
			}
			if len(extensions) > 0 {
				field.Extensions = extensions
			}
		}

		fields = append(fields, field)
	}

	return &ir.IRType{
		Name:        goName,
		Description: schema.Description,
		Kind:        ir.IRKindStruct,
		Fields:      fields,
	}
}

// collectUnionVariants builds IRTypeRef variants from oneOf/anyOf schemas.
// For primitive types, it returns builtin refs. For complex types, it creates
// inline types and returns named refs.
func collectUnionVariants(root *jsonschema.Schema, schema *jsonschema.Schema, contextName string, inlineTypes *[]ir.IRType) []ir.IRTypeRef {
	var schemas []*jsonschema.Schema
	if len(schema.OneOf) > 0 {
		schemas = schema.OneOf
	} else {
		schemas = schema.AnyOf
	}

	variants := make([]ir.IRTypeRef, 0, len(schemas))
	for _, s := range schemas {
		if s == nil {
			continue
		}

		// Handle null type
		if s.Type == "null" {
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinAny, Nullable: true})
			continue
		}

		// Handle $ref
		if s.Ref != "" {
			typeName := refToTypeName(s.Ref)
			variants = append(variants, ir.IRTypeRef{Name: typeName})
			continue
		}

		// Handle primitive types
		switch s.Type {
		case "string":
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinString, Format: schemaFormatToIRFormat(s.Format)})
		case "integer":
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinInt})
		case "number":
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinFloat})
		case "boolean":
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinBool})
		case "array":
			elem := schemaToIRTypeRefWithContext(root, s.Items, contextName+"Item", inlineTypes)
			variants = append(variants, ir.IRTypeRef{Array: &elem})
		case "object":
			// For non-discriminated unions, inline objects are represented as any/map
			// since most languages can't express anonymous object types in unions.
			// If the object has no properties, it's a generic map.
			if s.Properties != nil {
				variants = append(variants, ir.IRTypeRef{Map: &ir.IRTypeRef{Builtin: ir.IRBuiltinAny}})
			} else {
				variants = append(variants, ir.IRTypeRef{Map: &ir.IRTypeRef{Builtin: ir.IRBuiltinAny}})
			}
		default:
			// Unknown or empty type, use any
			variants = append(variants, ir.IRTypeRef{Builtin: ir.IRBuiltinAny})
		}
	}

	return variants
}

func schemaToIRTypeRefWithContext(root *jsonschema.Schema, schema *jsonschema.Schema, contextName string, inlineTypes *[]ir.IRType) ir.IRTypeRef {
	if schema == nil {
		return ir.IRTypeRef{Builtin: ir.IRBuiltinAny}
	}

	constraints := extractConstraints(schema)

	if schema.Ref != "" {
		typeName := refToTypeName(schema.Ref)
		return ir.IRTypeRef{Name: typeName, Constraints: constraints}
	}

	if schema.Const != nil {
		return ir.IRTypeRef{Builtin: ir.IRBuiltinString, Constraints: constraints}
	}

	if len(schema.AnyOf) > 0 || len(schema.OneOf) > 0 {
		return ir.IRTypeRef{Builtin: ir.IRBuiltinAny, Constraints: constraints}
	}

	switch schema.Type {
	case "string":
		return ir.IRTypeRef{Builtin: ir.IRBuiltinString, Format: schemaFormatToIRFormat(schema.Format), Constraints: constraints}
	case "integer":
		return ir.IRTypeRef{Builtin: ir.IRBuiltinInt, Constraints: constraints}
	case "number":
		return ir.IRTypeRef{Builtin: ir.IRBuiltinFloat, Constraints: constraints}
	case "boolean":
		return ir.IRTypeRef{Builtin: ir.IRBuiltinBool, Constraints: constraints}
	case "array":
		elem := schemaToIRTypeRefWithContext(root, schema.Items, contextName+"Item", inlineTypes)
		return ir.IRTypeRef{Array: &elem, Constraints: constraints}
	case "object":
		// Check if this is an inline object with properties
		if schema.Properties != nil {
			// Generate an inline type
			inlineType := convertStructToIRType(root, contextName, schema, inlineTypes)
			if inlineType != nil {
				*inlineTypes = append(*inlineTypes, *inlineType)
				return ir.IRTypeRef{Name: inlineType.Name, Constraints: constraints}
			}
		}
		// Check for typed additionalProperties (e.g. map[string]SomeType)
		if schema.AdditionalProperties != nil && (schema.AdditionalProperties.Ref != "" || schema.AdditionalProperties.Type != "") {
			valueRef := schemaToIRTypeRefWithContext(root, schema.AdditionalProperties, contextName+"Value", inlineTypes)
			return ir.IRTypeRef{Map: &valueRef, Constraints: constraints}
		}
		return ir.IRTypeRef{Map: &ir.IRTypeRef{Builtin: ir.IRBuiltinAny}, Constraints: constraints}
	}

	return ir.IRTypeRef{Builtin: ir.IRBuiltinAny, Constraints: constraints}
}

// parseDefault converts a JSON Schema default value (json.RawMessage) into an IRDefault.
// Returns nil if the raw value is empty or null.
func parseDefault(raw json.RawMessage, typeRef *ir.IRTypeRef) *ir.IRDefault {
	if raw == nil {
		return nil
	}

	rawStr := strings.TrimSpace(string(raw))
	if rawStr == "" || rawStr == "null" {
		return nil
	}

	builtin := typeRef.Builtin
	if builtin == ir.IRBuiltinNone && typeRef.Name != "" {
		// Named type reference (e.g., enum default) — treat as string
		builtin = ir.IRBuiltinString
	}

	return &ir.IRDefault{
		RawValue: rawStr,
		Builtin:  builtin,
	}
}

func getConstValue(v interface{}) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return ""
		}
		rv = rv.Elem()
	}

	return fmt.Sprintf("%v", rv.Interface())
}

func symbolName(s string) string {
	if s == "" {
		return ""
	}

	// Strip leading special chars (common in JSON Schema: $schema, @type, #id, etc.)
	s = strings.TrimLeft(s, "$@#")

	// If the name is already in valid PascalCase format (starts with uppercase, no delimiters), keep it as-is
	// This preserves acronyms like "RPCRequestBase" instead of converting to "RpcrequestBase"
	if isValidPascalCase(s) {
		return s
	}

	// Use casing package which handles acronyms properly
	name := casing.ToPascalCase(s)

	// If name starts with a number, prefix with "N"
	if len(name) > 0 && name[0] >= '0' && name[0] <= '9' {
		name = "N" + name
	}

	return name
}

// isValidPascalCase checks if a string is already in PascalCase format
// (starts with uppercase, contains no delimiters like _ or -)
func isValidPascalCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	// Must start with uppercase letter
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}
	// Must not contain delimiters
	for _, r := range s {
		if r == '_' || r == '-' || r == ' ' {
			return false
		}
	}
	return true
}

func refToTypeName(ref string) string {
	if idx := strings.LastIndex(ref, "/"); idx >= 0 {
		ref = ref[idx+1:]
	}
	ref = strings.TrimSuffix(ref, ".json")
	ref = strings.TrimSuffix(ref, ".yaml")
	return symbolName(ref)
}

func schemaFormatToIRFormat(format string) ir.IRFormat {
	if format == "" {
		return ir.IRFormatNone
	}

	switch format {
	case "byte":
		return ir.IRFormatByte
	case "date-time":
		return ir.IRFormatDateTime
	case "date":
		return ir.IRFormatDate
	case "uuid":
		return ir.IRFormatUUID
	case "email":
		return ir.IRFormatEmail
	case "uri":
		return ir.IRFormatURI
	default:
		// Pass through custom formats as-is
		// This allows users to define custom format mappings in their config
		return ir.IRFormat(format)
	}
}

// classifyEnum analyzes enum values and returns the enum type, typed values, and whether null is present.
// Returns IRBuiltinAny if the enum has mixed types (string + int).
func classifyEnum(values []any) (ir.IRBuiltin, []ir.IREnumValue, bool) {
	var result []ir.IREnumValue
	hasString := false
	hasInt := false
	hasNull := false

	for _, v := range values {
		switch val := v.(type) {
		case string:
			hasString = true
			result = append(result, ir.IREnumValue{StringValue: val})
		case float64:
			// JSON numbers come as float64, check if it's actually an integer
			if val == float64(int(val)) {
				hasInt = true
				intVal := int(val)
				// Use the string representation as the name
				result = append(result, ir.IREnumValue{
					StringValue: fmt.Sprintf("%d", intVal),
					IntValue:    &intVal,
				})
			} else {
				// Non-integer number, treat as mixed
				return ir.IRBuiltinAny, nil, false
			}
		case nil:
			hasNull = true
			result = append(result, ir.IREnumValue{IsNull: true})
		default:
			// Unknown type, fall back to any
			return ir.IRBuiltinAny, nil, false
		}
	}

	// Determine enum type
	if hasString && !hasInt {
		return ir.IRBuiltinString, result, hasNull
	}
	if hasInt && !hasString {
		return ir.IRBuiltinInt, result, hasNull
	}
	// Mixed string + int
	return ir.IRBuiltinAny, nil, hasNull
}

// extractConstraints extracts JSON Schema validation constraints from a schema.
// Returns nil if no constraints are present.
func extractConstraints(schema *jsonschema.Schema) *ir.IRConstraints {
	if schema == nil {
		return nil
	}

	c := &ir.IRConstraints{}
	hasConstraints := false

	// String constraints
	if schema.MinLength != nil {
		v := int(*schema.MinLength)
		c.MinLength = &v
		hasConstraints = true
	}
	if schema.MaxLength != nil {
		v := int(*schema.MaxLength)
		c.MaxLength = &v
		hasConstraints = true
	}
	if schema.Pattern != "" {
		c.Pattern = schema.Pattern
		hasConstraints = true
	}

	// Numeric constraints
	if schema.Minimum != nil {
		c.Minimum = schema.Minimum
		hasConstraints = true
	}
	if schema.Maximum != nil {
		c.Maximum = schema.Maximum
		hasConstraints = true
	}
	if schema.ExclusiveMinimum != nil {
		c.ExclusiveMinimum = schema.ExclusiveMinimum
		hasConstraints = true
	}
	if schema.ExclusiveMaximum != nil {
		c.ExclusiveMaximum = schema.ExclusiveMaximum
		hasConstraints = true
	}
	if schema.MultipleOf != nil {
		c.MultipleOf = schema.MultipleOf
		hasConstraints = true
	}

	// Array constraints
	if schema.MinItems != nil {
		v := int(*schema.MinItems)
		c.MinItems = &v
		hasConstraints = true
	}
	if schema.MaxItems != nil {
		v := int(*schema.MaxItems)
		c.MaxItems = &v
		hasConstraints = true
	}
	if schema.UniqueItems {
		c.UniqueItems = true
		hasConstraints = true
	}

	if !hasConstraints {
		return nil
	}

	return c
}

// topologicalSort sorts types so that dependencies are declared before dependents.
// Uses depth-first search for topological sorting with deterministic ordering.
func topologicalSort(types []ir.IRType) []ir.IRType {
	if len(types) == 0 {
		return types
	}

	// Build name -> type map and collect names for deterministic ordering
	typeMap := make(map[string]ir.IRType)
	var typeNames []string
	for _, t := range types {
		typeMap[t.Name] = t
		typeNames = append(typeNames, t.Name)
	}
	// Sort names alphabetically for deterministic processing
	sort.Strings(typeNames)

	// Build dependency graph (type -> types it depends on)
	deps := make(map[string][]string)
	for _, t := range types {
		depSet := extractTypeDependencies(t, typeMap)
		var depList []string
		for dep := range depSet {
			depList = append(depList, dep)
		}
		// Sort dependencies for deterministic ordering
		sort.Strings(depList)
		deps[t.Name] = depList
	}

	// DFS-based topological sort
	var result []ir.IRType
	visited := make(map[string]bool)

	var visit func(name string)
	visit = func(name string) {
		if visited[name] {
			return
		}
		visited[name] = true

		// Visit all dependencies first (in sorted order)
		for _, dep := range deps[name] {
			if _, exists := typeMap[dep]; exists {
				visit(dep)
			}
		}

		// Add this type after its dependencies
		if t, exists := typeMap[name]; exists {
			result = append(result, t)
		}
	}

	// Visit all types in sorted order
	for _, name := range typeNames {
		visit(name)
	}

	return result
}

// extractTypeDependencies returns the set of type names that a type depends on.
func extractTypeDependencies(t ir.IRType, typeMap map[string]ir.IRType) map[string]bool {
	deps := make(map[string]bool)

	// Helper to extract deps from a type ref
	var extractFromRef func(ref *ir.IRTypeRef)
	extractFromRef = func(ref *ir.IRTypeRef) {
		if ref == nil {
			return
		}
		if ref.Name != "" {
			deps[ref.Name] = true
		}
		if ref.Array != nil {
			extractFromRef(ref.Array)
		}
		if ref.Map != nil {
			extractFromRef(ref.Map)
		}
	}

	// Extract from fields
	for _, f := range t.Fields {
		extractFromRef(&f.Type)
	}

	// Extract from element (for aliases)
	if t.Element != nil {
		extractFromRef(t.Element)
	}

	// Extract from base type (allOf $ref inheritance)
	if t.BaseType != "" {
		deps[t.BaseType] = true
	}

	// Extract from discriminated union variants
	if t.Union != nil {
		for _, v := range t.Union.Variants {
			for _, f := range v.Type.Fields {
				extractFromRef(&f.Type)
			}
			// Also depend on the variant's base type
			if v.Type.BaseType != "" {
				deps[v.Type.BaseType] = true
			}
		}
	}

	// Extract from simple union variants
	if t.SimpleUnion != nil {
		for i := range t.SimpleUnion.Variants {
			extractFromRef(&t.SimpleUnion.Variants[i])
		}
	}

	return deps
}

// collectBaseFields converts a list of resolved base schemas into IR fields,
// suitable for merging into discriminated union variant structs.
func collectBaseFields(root *jsonschema.Schema, baseSchemas []*jsonschema.Schema, inlineTypes *[]ir.IRType) []ir.IRField {
	var fields []ir.IRField
	seen := make(map[string]bool)

	for _, base := range baseSchemas {
		// Recursively resolve allOf within each base schema
		resolved := merge.AllOf(root, base)
		if resolved == nil {
			resolved = base
		}
		if resolved.Properties == nil {
			continue
		}

		requiredSet := make(map[string]bool)
		for _, r := range resolved.Required {
			requiredSet[r] = true
		}

		propNames := make([]string, 0, len(resolved.Properties))
		for propName := range resolved.Properties {
			propNames = append(propNames, propName)
		}
		sort.Strings(propNames)

		for _, propName := range propNames {
			if seen[propName] {
				continue
			}
			seen[propName] = true

			propSchema := resolved.Properties[propName]
			fieldName := symbolName(propName)
			var fieldDesc string
			if propSchema != nil {
				fieldDesc = propSchema.Description
			}
			fields = append(fields, ir.IRField{
				Name:        fieldName,
				Description: fieldDesc,
				JSONName:    propName,
				Type:        schemaToIRTypeRefWithContext(root, propSchema, fieldName, inlineTypes),
				Required:    requiredSet[propName],
			})
		}
	}

	return fields
}

// mergeBaseFieldsIntoUnion finds the already-generated discriminated union IR type by name
// and prepends baseFields to each variant's field list (if not already present).
func mergeBaseFieldsIntoUnion(unionName string, baseFields []ir.IRField, types []ir.IRType) {
	if len(baseFields) == 0 {
		return
	}
	for i, t := range types {
		if t.Kind != ir.IRKindDiscriminatedUnion || t.Name != unionName {
			continue
		}
		if t.Union == nil {
			continue
		}
		for j := range t.Union.Variants {
			existing := make(map[string]bool)
			for _, f := range t.Union.Variants[j].Type.Fields {
				existing[f.JSONName] = true
			}
			// Prepend base fields that are not already present
			var merged []ir.IRField
			for _, bf := range baseFields {
				if !existing[bf.JSONName] {
					merged = append(merged, bf)
				}
			}
			merged = append(merged, t.Union.Variants[j].Type.Fields...)
			// Re-sort by name for deterministic output
			sort.Slice(merged, func(a, b int) bool {
				return merged[a].Name < merged[b].Name
			})
			types[i].Union.Variants[j].Type.Fields = merged
		}
		return
	}
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

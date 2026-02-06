package typescriptzod

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in Zod
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "z.string()"}, // Base64 encoded
	ir.IRFormatDateTime: {Type: "z.iso.datetime()"},
	ir.IRFormatDate:     {Type: "z.iso.date()"},
	ir.IRFormatUUID:     {Type: "z.string().uuid()"},
	ir.IRFormatEmail:    {Type: "z.string().email()"},
	ir.IRFormatURI:      {Type: "z.string().url()"},
}

// config holds TypeScript Zod-specific generator configuration
type config struct {
	// Whether to export all types (default: true)
	exportTypes bool
}

// Option is a TypeScript Zod-specific generator option
type Option struct {
	apply func(*config)
}

// OptionValue implements generators.GeneratorOption
func (Option) OptionValue() string { return "typescript-zod" }

// WithExportTypes sets whether to export all types
func WithExportTypes(export bool) Option {
	return Option{apply: func(c *config) {
		c.exportTypes = export
	}}
}

type Generator struct{}

func (g *Generator) getFormatMappings(opts generators.GeneratorOptions) map[ir.IRFormat]generators.FormatTypeMapping {
	result := make(map[ir.IRFormat]generators.FormatTypeMapping)
	for k, v := range DefaultFormatMappings {
		result[k] = v
	}
	for k, v := range opts.FormatMappings {
		result[k] = v
	}
	return result
}

func (g *Generator) Generate(data *ir.IR, opts generators.GeneratorOptions, genOpts ...generators.GeneratorOption) ([]generators.GeneratedFile, error) {
	cfg := &config{
		exportTypes: true,
	}
	for _, opt := range genOpts {
		if zodOpt, ok := opt.(Option); ok {
			zodOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)
	recursiveFields := findRecursiveFields(data.Types)
	recursiveTypes := findRecursiveTypes(data.Types)

	// Identify recursive union types where "typeof XSchema" in a getter
	// return type annotation would create a circular type reference.
	// Discriminated unions directly reference their variants (non-lazily),
	// so TypeScript can't break the cycle through getter deferred resolution.
	unsafeTypeofTypes := make(map[string]bool)
	for _, t := range data.Types {
		if (t.Kind == ir.IRKindDiscriminatedUnion || t.Kind == ir.IRKindUnion) && recursiveTypes[t.Name] {
			unsafeTypeofTypes[t.Name] = true
		}
	}

	funcs := template.FuncMap{
		"pascal":    casing.ToPascalCase,
		"camel":     casing.ToCamelCase,
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"zodType":       makeZodTypeFunc(formatMappings),
		"zodReturnType": makeZodReturnTypeFunc(formatMappings, unsafeTypeofTypes),
		"tsType":        makeTsTypeFunc(),
		"comment":       formatComment,
		"hasPrefix": strings.HasPrefix,
		"export":    func() string { return exportKeyword(cfg.exportTypes) },
		"isIntEnum": isIntEnum,
		"isRecursiveField": func(typeName, fieldName string) bool {
			if fields, ok := recursiveFields[typeName]; ok {
				return fields[fieldName]
			}
			return false
		},
		"isRecursiveType": func(typeName string) bool {
			return recursiveTypes[typeName]
		},
		"hasRecursiveFields": func(typeName string) bool {
			fields, ok := recursiveFields[typeName]
			return ok && len(fields) > 0
		},
	}

	tmpl, err := template.New("typescript-zod").Funcs(funcs).Parse(zodTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return []generators.GeneratedFile{{
		Filename: "schema.ts",
		Content:  buf.Bytes(),
	}}, nil
}

func exportKeyword(export bool) string {
	if export {
		return "export "
	}
	return ""
}

func formatComment(description string) string {
	if description == "" {
		return ""
	}
	description = strings.TrimRight(description, "\n")
	lines := strings.Split(description, "\n")
	var result []string
	for _, line := range lines {
		result = append(result, "// "+line)
	}
	return strings.Join(result, "\n")
}

// isIntEnum returns true if the enum has an integer type
func isIntEnum(t ir.IRType) bool {
	return t.EnumType == ir.IRBuiltinInt
}

// collectNamedRefs extracts all named type references from an IRTypeRef.
func collectNamedRefs(ref *ir.IRTypeRef, refs map[string]bool) {
	if ref == nil {
		return
	}
	if ref.Name != "" {
		refs[ref.Name] = true
	}
	if ref.Array != nil {
		collectNamedRefs(ref.Array, refs)
	}
	if ref.Map != nil {
		collectNamedRefs(ref.Map, refs)
	}
}

// canReach checks if 'from' can reach 'to' in the dependency graph using DFS.
func canReach(deps map[string]map[string]bool, from, to string) bool {
	visited := make(map[string]bool)
	var dfs func(current string) bool
	dfs = func(current string) bool {
		if current == to {
			return true
		}
		if visited[current] {
			return false
		}
		visited[current] = true
		for dep := range deps[current] {
			if dfs(dep) {
				return true
			}
		}
		return false
	}
	return dfs(from)
}

// buildTypeDependencyGraph builds a dependency graph for ALL type kinds,
// mapping each type name to the set of type names it directly references.
func buildTypeDependencyGraph(types []ir.IRType) map[string]map[string]bool {
	deps := make(map[string]map[string]bool)
	for _, t := range types {
		refs := make(map[string]bool)
		switch t.Kind {
		case ir.IRKindStruct:
			for _, f := range t.Fields {
				collectNamedRefs(&f.Type, refs)
			}
		case ir.IRKindAlias:
			if t.Element != nil {
				collectNamedRefs(t.Element, refs)
			}
		case ir.IRKindDiscriminatedUnion:
			if t.Union != nil {
				// The union type itself depends on each variant
				for _, v := range t.Union.Variants {
					refs[v.Name] = true
				}
				// Each variant is also a pseudo-type in the graph
				for _, v := range t.Union.Variants {
					variantRefs := make(map[string]bool)
					for _, f := range v.Type.Fields {
						collectNamedRefs(&f.Type, variantRefs)
					}
					deps[v.Name] = variantRefs
				}
			}
		case ir.IRKindUnion:
			if t.SimpleUnion != nil {
				for i := range t.SimpleUnion.Variants {
					collectNamedRefs(&t.SimpleUnion.Variants[i], refs)
				}
			}
		}
		deps[t.Name] = refs
	}
	return deps
}

// findRecursiveFields identifies fields in struct types that participate in
// reference cycles. Returns a map of type name -> set of field JSON names
// that need getter syntax for Zod v4 recursive schemas.
func findRecursiveFields(types []ir.IRType) map[string]map[string]bool {
	deps := buildTypeDependencyGraph(types)

	// For each struct type, check which fields reference types that can
	// reach back to this type (forming a cycle)
	result := make(map[string]map[string]bool)
	for _, t := range types {
		if t.Kind == ir.IRKindStruct {
			for _, f := range t.Fields {
				fieldRefs := make(map[string]bool)
				collectNamedRefs(&f.Type, fieldRefs)
				for ref := range fieldRefs {
					if canReach(deps, ref, t.Name) {
						if result[t.Name] == nil {
							result[t.Name] = make(map[string]bool)
						}
						result[t.Name][f.JSONName] = true
						break
					}
				}
			}
		}
		// Also check discriminated union variant fields
		if t.Kind == ir.IRKindDiscriminatedUnion && t.Union != nil {
			for _, v := range t.Union.Variants {
				for _, f := range v.Type.Fields {
					fieldRefs := make(map[string]bool)
					collectNamedRefs(&f.Type, fieldRefs)
					for ref := range fieldRefs {
						if canReach(deps, ref, v.Name) {
							if result[v.Name] == nil {
								result[v.Name] = make(map[string]bool)
							}
							result[v.Name][f.JSONName] = true
							break
						}
					}
				}
			}
		}
	}
	return result
}

// findRecursiveTypes identifies types that participate in reference cycles.
// Returns a set of type names that are part of at least one cycle.
func findRecursiveTypes(types []ir.IRType) map[string]bool {
	deps := buildTypeDependencyGraph(types)
	result := make(map[string]bool)
	for _, t := range types {
		// A type is in a cycle if any of its direct dependencies can
		// reach back to it (requiring at least one edge traversal).
		for dep := range deps[t.Name] {
			if canReach(deps, dep, t.Name) {
				result[t.Name] = true
				break
			}
		}
	}
	return result
}

func makeZodTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping) func(*ir.IRTypeRef) string {
	var zodType func(*ir.IRTypeRef) string
	zodType = func(ref *ir.IRTypeRef) string {
		var baseType string
		isNumeric := false
		isString := false
		isArray := false

		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			baseType = mapping.Type
			// Format types that are strings underneath
			isString = ref.Format == ir.IRFormatUUID || ref.Format == ir.IRFormatEmail ||
				ref.Format == ir.IRFormatURI || ref.Format == ir.IRFormatByte
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "z.string()"
					isString = true
				case ir.IRBuiltinInt:
					baseType = "z.number().int()"
					isNumeric = true
				case ir.IRBuiltinFloat:
					baseType = "z.number()"
					isNumeric = true
				case ir.IRBuiltinBool:
					baseType = "z.boolean()"
				case ir.IRBuiltinAny:
					baseType = "z.unknown()"
				}
			} else if ref.Array != nil {
				baseType = "z.array(" + zodType(ref.Array) + ")"
				isArray = true
			} else if ref.Map != nil {
				baseType = "z.record(z.string(), " + zodType(ref.Map) + ")"
			} else if ref.Name != "" {
				baseType = ref.Name + "Schema"
			} else {
				baseType = "z.unknown()"
			}
		}

		// Apply constraints
		if c := ref.Constraints; c != nil {
			// String constraints
			if isString {
				if c.MinLength != nil {
					baseType += fmt.Sprintf(".min(%d)", *c.MinLength)
				}
				if c.MaxLength != nil {
					baseType += fmt.Sprintf(".max(%d)", *c.MaxLength)
				}
				if c.Pattern != "" {
					// Escape backslashes for JavaScript regex
					pattern := strings.ReplaceAll(c.Pattern, "\\", "\\\\")
					baseType += fmt.Sprintf(".regex(/%s/)", pattern)
				}
			}

			// Numeric constraints
			if isNumeric {
				if c.Minimum != nil {
					baseType += fmt.Sprintf(".min(%v)", *c.Minimum)
				}
				if c.Maximum != nil {
					baseType += fmt.Sprintf(".max(%v)", *c.Maximum)
				}
				if c.ExclusiveMinimum != nil {
					baseType += fmt.Sprintf(".gt(%v)", *c.ExclusiveMinimum)
				}
				if c.ExclusiveMaximum != nil {
					baseType += fmt.Sprintf(".lt(%v)", *c.ExclusiveMaximum)
				}
				if c.MultipleOf != nil {
					baseType += fmt.Sprintf(".multipleOf(%v)", *c.MultipleOf)
				}
			}

			// Array constraints
			if isArray {
				if c.MinItems != nil {
					baseType += fmt.Sprintf(".min(%d)", *c.MinItems)
				}
				if c.MaxItems != nil {
					baseType += fmt.Sprintf(".max(%d)", *c.MaxItems)
				}
			}
		}

		if ref.Nullable {
			baseType += ".nullable()"
		}

		return baseType
	}
	return zodType
}

// makeTsTypeFunc creates a function that converts IRTypeRef to TypeScript type strings.
// Used to generate explicit interface/type declarations for recursive types
// where z.infer would yield unknown.
func makeTsTypeFunc() func(*ir.IRTypeRef) string {
	var tsType func(*ir.IRTypeRef) string
	tsType = func(ref *ir.IRTypeRef) string {
		var baseType string

		if ref.Array != nil {
			inner := tsType(ref.Array)
			if strings.Contains(inner, " | ") {
				baseType = "(" + inner + ")[]"
			} else {
				baseType = inner + "[]"
			}
		} else if ref.Map != nil {
			baseType = "Record<string, " + tsType(ref.Map) + ">"
		} else if ref.Name != "" {
			baseType = ref.Name
		} else if ref.Format != "" {
			// All standard JSON Schema formats map to string in TypeScript
			baseType = "string"
		} else {
			switch ref.Builtin {
			case ir.IRBuiltinString:
				baseType = "string"
			case ir.IRBuiltinInt, ir.IRBuiltinFloat:
				baseType = "number"
			case ir.IRBuiltinBool:
				baseType = "boolean"
			default:
				baseType = "unknown"
			}
		}

		if ref.Nullable {
			baseType += " | null"
		}

		return baseType
	}
	return tsType
}

// makeZodReturnTypeFunc creates a function that generates Zod type annotations
// for getter return types. unsafeTypeofTypes is a set of type names (recursive
// unions) where using "typeof XSchema" in a return type annotation would create
// a circular type reference that TypeScript cannot resolve. For those types,
// we fall back to "z.ZodType" to break the cycle.
func makeZodReturnTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping, unsafeTypeofTypes map[string]bool) func(*ir.IRTypeRef, bool) string {
	var inner func(*ir.IRTypeRef) string
	inner = func(ref *ir.IRTypeRef) string {
		var baseType string

		// Check format first
		if _, ok := formatMappings[ref.Format]; ok {
			isStringFormat := ref.Format == ir.IRFormatUUID || ref.Format == ir.IRFormatEmail ||
				ref.Format == ir.IRFormatURI || ref.Format == ir.IRFormatByte
			if isStringFormat {
				baseType = "z.ZodString"
			} else {
				baseType = "z.ZodType"
			}
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "z.ZodString"
				case ir.IRBuiltinInt, ir.IRBuiltinFloat:
					baseType = "z.ZodNumber"
				case ir.IRBuiltinBool:
					baseType = "z.ZodBoolean"
				case ir.IRBuiltinAny:
					baseType = "z.ZodUnknown"
				}
			} else if ref.Array != nil {
				baseType = "z.ZodArray<" + inner(ref.Array) + ">"
			} else if ref.Map != nil {
				baseType = "z.ZodRecord<z.ZodString, " + inner(ref.Map) + ">"
			} else if ref.Name != "" {
				if unsafeTypeofTypes[ref.Name] {
					baseType = "z.ZodType"
				} else {
					baseType = "typeof " + ref.Name + "Schema"
				}
			} else {
				baseType = "z.ZodType"
			}
		}

		if ref.Nullable {
			baseType = "z.ZodNullable<" + baseType + ">"
		}

		return baseType
	}

	return func(ref *ir.IRTypeRef, required bool) string {
		result := inner(ref)
		if !required {
			result = "z.ZodOptional<" + result + ">"
		}
		return result
	}
}

const zodTemplate = `import { z } from "zod";
{{range $i, $t := .Types}}
{{- if eq .Kind "struct"}}
{{template "struct" .}}
{{- else if eq .Kind "alias"}}
{{template "alias" .}}
{{- else if eq .Kind "enum"}}
{{template "enum" .}}
{{- else if eq .Kind "discriminated_union"}}
{{template "union" .}}
{{- else if eq .Kind "union"}}
{{template "simpleunion" .}}
{{- end}}
{{end}}
{{- define "struct" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{export}}const {{.Name}}Schema = z.object({
{{- range $i, $f := .Fields}}
{{- if isRecursiveField $.Name $f.JSONName}}
  get {{$f.JSONName}}(): {{zodReturnType $f.Type $f.Required}} { return {{zodType $f.Type}}{{if not $f.Required}}.optional(){{end}}; },
{{- else}}
  {{$f.JSONName}}: {{zodType $f.Type}}{{if not $f.Required}}.optional(){{end}},
{{- end}}
{{- end}}
});
{{- if hasRecursiveFields .Name}}
{{export}}interface {{.Name}} {
{{- range $i, $f := .Fields}}
  {{$f.JSONName}}{{if not $f.Required}}?{{end}}: {{tsType $f.Type}};
{{- end}}
}
{{- else}}
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
{{- end -}}

{{- define "alias" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if isRecursiveType .Name -}}
{{export}}const {{.Name}}Schema = z.lazy(() => {{if .Element}}{{zodType .Element}}{{else}}z.unknown(){{end}});
{{export}}type {{.Name}} = {{if .Element}}{{tsType .Element}}{{else}}unknown{{end}};
{{- else -}}
{{export}}const {{.Name}}Schema = {{if .Element}}{{zodType .Element}}{{else}}z.unknown(){{end}};
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
{{- end -}}

{{- define "enum" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if isIntEnum . -}}
{{export}}const {{.Name}}Schema = z.union([{{range $i, $v := .EnumValues}}{{if $i}}, {{end}}{{if $v.IsNull}}z.null(){{else}}z.literal({{$v.IntValue}}){{end}}{{end}}]);
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- else -}}
{{export}}const {{.Name}}Schema = z.enum([{{range $i, $v := .Enum}}{{if $i}}, {{end}}"{{$v}}"{{end}}]);
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
{{- end -}}

{{- define "union" -}}
{{- range $i, $v := .Union.Variants}}
{{- if .Type.Description}}
{{comment .Type.Description}}
{{end -}}
{{export}}const {{.Name}}Schema = z.object({
  {{$.Union.DiscriminatorJSON}}: z.literal("{{.ConstValue}}"),
{{- range .Type.Fields}}
{{- if ne .JSONName $.Union.DiscriminatorJSON}}
{{- if isRecursiveField $v.Name .JSONName}}
  get {{.JSONName}}(): {{zodReturnType .Type .Required}} { return {{zodType .Type}}{{if not .Required}}.optional(){{end}}; },
{{- else}}
  {{.JSONName}}: {{zodType .Type}}{{if not .Required}}.optional(){{end}},
{{- end}}
{{- end}}
{{- end}}
});
{{- if hasRecursiveFields .Name}}
{{export}}interface {{.Name}} {
  {{$.Union.DiscriminatorJSON}}: '{{.ConstValue}}';
{{- range .Type.Fields}}
{{- if ne .JSONName $.Union.DiscriminatorJSON}}
  {{.JSONName}}{{if not .Required}}?{{end}}: {{tsType .Type}};
{{- end}}
{{- end}}
}
{{- else}}
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end}}

{{end -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if isRecursiveType .Name -}}
{{export}}const {{.Name}}Schema = z.union([
{{- else -}}
{{export}}const {{.Name}}Schema = z.discriminatedUnion("{{.Union.DiscriminatorJSON}}", [
{{- end}}
{{- range $i, $v := .Union.Variants}}
  {{$v.Name}}Schema,
{{- end}}
]);
{{- if isRecursiveType .Name}}
{{export}}type {{.Name}} = {{range $i, $v := .Union.Variants}}{{if $i}} | {{end}}{{$v.Name}}{{end}};
{{- else}}
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
{{- end -}}

{{- define "simpleunion" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{export}}const {{.Name}}Schema = z.union([
{{- range $i, $v := .SimpleUnion.Variants}}
  {{zodType $v}},
{{- end}}
]);
{{- if isRecursiveType .Name}}
{{export}}type {{.Name}} = {{range $i, $v := .SimpleUnion.Variants}}{{if $i}} | {{end}}{{tsType $v}}{{end}};
{{- else}}
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
{{- end -}}
`

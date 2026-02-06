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

	funcs := template.FuncMap{
		"pascal":    casing.ToPascalCase,
		"camel":     casing.ToCamelCase,
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"zodType":   makeZodTypeFunc(formatMappings),
		"comment":   formatComment,
		"hasPrefix": strings.HasPrefix,
		"export":    func() string { return exportKeyword(cfg.exportTypes) },
		"isIntEnum": isIntEnum,
		"isRecursiveField": func(typeName, fieldName string) bool {
			if fields, ok := recursiveFields[typeName]; ok {
				return fields[fieldName]
			}
			return false
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

// findRecursiveFields identifies fields in struct types that participate in
// reference cycles. Returns a map of type name -> set of field JSON names
// that need getter syntax for Zod v4 recursive schemas.
func findRecursiveFields(types []ir.IRType) map[string]map[string]bool {
	// Build dependency graph: type -> set of types it references
	deps := make(map[string]map[string]bool)
	for _, t := range types {
		if t.Kind != ir.IRKindStruct {
			continue
		}
		refs := make(map[string]bool)
		for _, f := range t.Fields {
			collectNamedRefs(&f.Type, refs)
		}
		deps[t.Name] = refs
	}

	// For each struct type, check which fields reference types that can
	// reach back to this type (forming a cycle)
	result := make(map[string]map[string]bool)
	for _, t := range types {
		if t.Kind != ir.IRKindStruct {
			continue
		}
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
  get {{$f.JSONName}}() { return {{zodType $f.Type}}{{if not $f.Required}}.optional(){{end}}; },
{{- else}}
  {{$f.JSONName}}: {{zodType $f.Type}}{{if not $f.Required}}.optional(){{end}},
{{- end}}
{{- end}}
});
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}

{{- define "alias" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{export}}const {{.Name}}Schema = {{if .Element}}{{zodType .Element}}{{else}}z.unknown(){{end}};
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
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
  {{.JSONName}}: {{zodType .Type}}{{if not .Required}}.optional(){{end}},
{{- end}}
{{- end}}
});
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;

{{end -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{export}}const {{.Name}}Schema = z.discriminatedUnion("{{.Union.DiscriminatorJSON}}", [
{{- range $i, $v := .Union.Variants}}
  {{$v.Name}}Schema,
{{- end}}
]);
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
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
{{export}}type {{.Name}} = z.infer<typeof {{.Name}}Schema>;
{{- end -}}
`

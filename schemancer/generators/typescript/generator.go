package typescript

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in TypeScript
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "string"}, // Base64 encoded
	ir.IRFormatDateTime: {Type: "Date"},
	ir.IRFormatDate:     {Type: "Date"},
	ir.IRFormatUUID:     {Type: "string"},
	ir.IRFormatEmail:    {Type: "string"},
	ir.IRFormatURI:      {Type: "string"},
}

// config holds TypeScript-specific generator configuration
type config struct {
	// Whether to export all types (default: true)
	exportTypes bool
	// Whether to use strict null checks style (T | null vs T | undefined)
	useNullForOptional bool
	// Whether to use branded types for primitive type aliases
	brandedTypes bool
}

// Option is a TypeScript-specific generator option
type Option struct {
	apply func(*config)
}

// OptionValue implements generators.GeneratorOption
func (Option) OptionValue() string { return "typescript" }

// WithExportTypes sets whether to export all types
func WithExportTypes(export bool) Option {
	return Option{apply: func(c *config) {
		c.exportTypes = export
	}}
}

// WithNullForOptional sets whether to use null instead of undefined for optional fields
func WithNullForOptional(useNull bool) Option {
	return Option{apply: func(c *config) {
		c.useNullForOptional = useNull
	}}
}

// WithBrandedTypes enables branded types for primitive type aliases.
// This provides nominal typing for types like `type UserID = Branded<string, "UserID">`
// instead of just `type UserID = string`, preventing accidental mixing of ID types.
func WithBrandedTypes(enabled bool) Option {
	return Option{apply: func(c *config) {
		c.brandedTypes = enabled
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
		exportTypes:        true,
		useNullForOptional: false,
	}
	for _, opt := range genOpts {
		if tsOpt, ok := opt.(Option); ok {
			tsOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)

	funcs := template.FuncMap{
		"pascal":           casing.ToPascalCase,
		"camel":            casing.ToCamelCase,
		"snake":            casing.ToSnakeCase,
		"kebab":            casing.ToKebabCase,
		"lower":            strings.ToLower,
		"upper":            strings.ToUpper,
		"tsType":           makeTsTypeFunc(formatMappings),
		"comment":          formatComment,
		"fieldComment":     formatFieldComment,
		"hasPrefix":        strings.HasPrefix,
		"export":           func() string { return exportKeyword(cfg.exportTypes) },
		"isPrimitiveAlias": isPrimitiveAlias,
		"isIntEnum":        isIntEnum,
		"toEnumKey":        toEnumKey,
	}

	tmpl, err := template.New("typescript").Funcs(funcs).Parse(tsTemplate)
	if err != nil {
		return nil, err
	}

	tplData := prepareTemplateData(data, cfg)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tplData); err != nil {
		return nil, err
	}

	return []generators.GeneratedFile{{
		Filename: "types.ts",
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
	return formatCommentWithIndent(description, "")
}

func formatCommentWithIndent(description, indent string) string {
	if description == "" {
		return ""
	}
	description = strings.TrimRight(description, "\n")
	lines := strings.Split(description, "\n")
	var result []string
	for _, line := range lines {
		result = append(result, indent+"// "+line)
	}
	return strings.Join(result, "\n")
}

func formatFieldComment(description string) string {
	return formatCommentWithIndent(description, "  ")
}

type templateData struct {
	Types      []ir.IRType
	HasBranded bool
	UseBranded bool
}

func prepareTemplateData(data *ir.IR, cfg *config) templateData {
	hasBranded := false
	if cfg.brandedTypes {
		for _, t := range data.Types {
			if t.Kind == ir.IRKindAlias && isPrimitiveAlias(t) {
				hasBranded = true
				break
			}
		}
	}
	return templateData{
		Types:      data.Types,
		HasBranded: hasBranded,
		UseBranded: cfg.brandedTypes,
	}
}

// isIntEnum returns true if the enum has an integer type
func isIntEnum(t ir.IRType) bool {
	return t.EnumType == ir.IRBuiltinInt
}

// toEnumKey converts an enum value to a valid TypeScript enum key
func toEnumKey(v ir.IREnumValue) string {
	if v.IntValue != nil {
		// For integer enums, use VALUE_N format
		return "VALUE_" + v.StringValue
	}
	// For string enums, convert to UPPER_SNAKE_CASE
	return strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(v.StringValue, "-", "_"), " ", "_"))
}

// isPrimitiveAlias returns true if the type is an alias to a primitive type
// (string, number, boolean) that should be branded for nominal typing.
func isPrimitiveAlias(t ir.IRType) bool {
	if t.Element == nil {
		return false
	}
	// Only brand simple primitives, not arrays, maps, or named types
	if t.Element.Array != nil || t.Element.Map != nil || t.Element.Name != "" {
		return false
	}
	// Check if it's a builtin primitive
	switch t.Element.Builtin {
	case ir.IRBuiltinString, ir.IRBuiltinInt, ir.IRBuiltinFloat, ir.IRBuiltinBool:
		return true
	}
	// Also brand formatted strings (uuid, email, etc.)
	if t.Element.Format != ir.IRFormatNone {
		return true
	}
	return false
}

func makeTsTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping) func(*ir.IRTypeRef) string {
	var tsType func(*ir.IRTypeRef) string
	tsType = func(ref *ir.IRTypeRef) string {
		var baseType string

		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			baseType = mapping.Type
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "string"
				case ir.IRBuiltinInt, ir.IRBuiltinFloat:
					baseType = "number"
				case ir.IRBuiltinBool:
					baseType = "boolean"
				case ir.IRBuiltinAny:
					baseType = "unknown"
				}
			} else if ref.Array != nil {
				baseType = tsType(ref.Array) + "[]"
			} else if ref.Map != nil {
				baseType = "Record<string, " + tsType(ref.Map) + ">"
			} else if ref.Name != "" {
				baseType = ref.Name
			} else {
				baseType = "unknown"
			}
		}

		return baseType
	}
	return tsType
}


const tsTemplate = `{{- define "brand" -}}
declare const __brand: unique symbol;
type Brand<B> = {[__brand]: B};
{{export}}type Branded<T, B> = T & Brand<B>;
{{- end -}}

{{- define "interface" -}}
{{- if .Description -}}
{{comment .Description}}
{{end -}}
{{export}}interface {{.Name}} {
{{- range .Fields}}
{{- if .Description}}
{{fieldComment .Description}}
{{- end}}
  {{.JSONName}}{{if not .Required}}?{{end}}: {{tsType .Type}};
{{- end}}
}
{{- end -}}

{{- define "alias" -}}
{{- if .Description -}}
{{comment .Description}}
{{end -}}
{{export}}type {{.Name}} = {{if .Element}}{{tsType .Element}}{{else}}unknown{{end}};
{{- end -}}

{{- define "branded_alias" -}}
{{- if .Description -}}
{{comment .Description}}
{{end -}}
{{export}}type {{.Name}} = Branded<{{if .Element}}{{tsType .Element}}{{else}}unknown{{end}}, "{{.Name}}">;
{{- end -}}

{{- define "enum" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if isIntEnum . -}}
{{export}}enum {{.Name}} {
{{- range .EnumValues}}
{{- if not .IsNull}}
  {{toEnumKey .}} = {{.IntValue}},
{{- end}}
{{- end}}
}
{{- else -}}
{{export}}type {{.Name}} =
{{- range $i, $v := .Enum}}
  | "{{$v}}"
{{- end}};
{{- end -}}
{{- end -}}

{{- define "union" -}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
{{- range $i, $v := .Union.Variants}}
{{- if $i}}
{{end}}
{{- if .Type.Description}}
{{comment .Type.Description}}
{{- end}}
{{export}}interface {{.Name}} {
  {{$.Union.DiscriminatorJSON}}: "{{.ConstValue}}";
{{- range .Type.Fields}}
{{- if ne .JSONName $.Union.DiscriminatorJSON}}
{{- if .Description}}
{{fieldComment .Description}}
{{- end}}
  {{.JSONName}}{{if not .Required}}?{{end}}: {{tsType .Type}};
{{- end}}
{{- end}}
}
{{- end}}

{{export}}type {{.Name}} =
{{- range $i, $v := .Union.Variants}}
  | {{$v.Name}}
{{- end}};
{{- range .Union.Variants}}

{{export}}function is{{.Name}}(value: {{$.Name}}): value is {{.Name}} {
  return value.{{$.Union.DiscriminatorJSON}} === "{{.ConstValue}}";
}
{{- end}}
{{- end -}}

{{- define "simpleunion" -}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{export}}type {{.Name}} = {{range $i, $v := .SimpleUnion.Variants}}{{if $i}} | {{end}}{{tsType $v}}{{end}};
{{- end -}}

{{- if .HasBranded -}}
{{template "brand" .}}

{{end -}}
{{- range $i, $t := .Types -}}
{{- if $i}}

{{end -}}
{{if eq .Kind "struct" -}}
{{template "interface" .}}
{{- else if eq .Kind "alias" -}}
{{- if and $.UseBranded (isPrimitiveAlias .) -}}
{{template "branded_alias" .}}
{{- else -}}
{{template "alias" .}}
{{- end -}}
{{- else if eq .Kind "enum" -}}
{{template "enum" .}}
{{- else if eq .Kind "discriminated_union" -}}
{{template "union" .}}
{{- else if eq .Kind "union" -}}
{{template "simpleunion" .}}
{{- end -}}
{{- end}}
`

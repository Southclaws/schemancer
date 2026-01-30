package golang

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in Go
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "[]byte"},
	ir.IRFormatDateTime: {Type: "time.Time", Import: "time"},
	ir.IRFormatDate:     {Type: "time.Time", Import: "time"},
	ir.IRFormatUUID:     {Type: "uuid.UUID", Import: "github.com/google/uuid"},
	ir.IRFormatEmail:    {Type: "mail.Address", Import: "net/mail"},
	ir.IRFormatURI:      {Type: "url.URL", Import: "net/url"},
}

// OptionalStyle determines how optional fields are represented
type OptionalStyle string

const (
	// OptionalStylePointer uses pointers for optional fields (default): *string
	OptionalStylePointer OptionalStyle = "pointer"
	// OptionalStyleOpt uses Southclaws/opt library: opt.Optional[string]
	OptionalStyleOpt OptionalStyle = "opt"
)

// config holds Go-specific generator configuration
type config struct {
	packageName   string
	optionalStyle OptionalStyle
}

// Option is a Go-specific generator option
type Option struct {
	apply func(*config)
}

// OptionValue implements generators.GeneratorOption
func (Option) OptionValue() string { return "golang" }

// WithPackageName sets the Go package name for generated code
func WithPackageName(name string) Option {
	return Option{apply: func(c *config) {
		c.packageName = name
	}}
}

// WithOptionalStyle sets how optional fields are represented
func WithOptionalStyle(style OptionalStyle) Option {
	return Option{apply: func(c *config) {
		c.optionalStyle = style
	}}
}

type Generator struct{}

func (g *Generator) getFormatMappings(opts generators.GeneratorOptions) map[ir.IRFormat]generators.FormatTypeMapping {
	// Start with defaults, allow overrides from options
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
	// Process Go-specific options
	cfg := &config{
		packageName:   "generated",
		optionalStyle: OptionalStylePointer,
	}
	for _, opt := range genOpts {
		if goOpt, ok := opt.(Option); ok {
			goOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)

	funcs := template.FuncMap{
		"pascal":     casing.ToPascalCase,
		"camel":      casing.ToCamelCase,
		"snake":      casing.ToSnakeCase,
		"kebab":      casing.ToKebabCase,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"goType":     makeGoTypeFunc(formatMappings, cfg.optionalStyle),
		"jsonTag":    jsonTag,
		"hasPrefix":  strings.HasPrefix,
		"trimPrefix": strings.TrimPrefix,
		"comment":    formatComment,
		"isIntEnum":  isIntEnum,
		"toEnumKey":  toEnumKey,
	}

	tmpl, err := template.New("go").Funcs(funcs).Parse(goTemplate)
	if err != nil {
		return nil, err
	}

	tplData := prepareTemplateData(cfg.packageName, cfg.optionalStyle, data, formatMappings)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Return unformatted if formatting fails
		return []generators.GeneratedFile{{
			Filename: cfg.packageName + ".go",
			Content:  buf.Bytes(),
		}}, err
	}

	return []generators.GeneratedFile{{
		Filename: cfg.packageName + ".go",
		Content:  formatted,
	}}, nil
}

func formatComment(description string) string {
	if description == "" {
		return ""
	}
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

// toEnumKey converts an enum value to a valid Go const name
func toEnumKey(typeName string, v ir.IREnumValue) string {
	if v.IntValue != nil {
		// For integer enums, use TypeName + Value format
		return typeName + v.StringValue
	}
	// For string enums, use TypeName + PascalCase format
	return typeName + casing.ToPascalCase(v.StringValue)
}

type templateData struct {
	Package  string
	HasUnion bool
	Imports  []string
	Types    []ir.IRType
}

func prepareTemplateData(packageName string, optStyle OptionalStyle, data *ir.IR, formatMappings map[ir.IRFormat]generators.FormatTypeMapping) templateData {
	hasUnion := false
	hasOptional := false
	importSet := make(map[string]bool)

	for _, t := range data.Types {
		if t.Kind == ir.IRKindDiscriminatedUnion {
			hasUnion = true
			collectImportsFromUnion(t, formatMappings, importSet)
			hasOptional = hasOptional || hasOptionalFields(t.Union)
		} else {
			collectImportsFromType(t, formatMappings, importSet)
			hasOptional = hasOptional || hasOptionalFieldsInType(t)
		}
	}

	// Add opt import if using opt style and there are optional fields
	if optStyle == OptionalStyleOpt && hasOptional {
		importSet["github.com/Southclaws/opt"] = true
	}

	var imports []string
	for imp := range importSet {
		imports = append(imports, imp)
	}

	return templateData{
		Package:  packageName,
		HasUnion: hasUnion,
		Imports:  imports,
		Types:    data.Types,
	}
}

func hasOptionalFieldsInType(t ir.IRType) bool {
	for _, f := range t.Fields {
		if !f.Required {
			return true
		}
	}
	return false
}

func hasOptionalFields(union *ir.IRDiscriminatedUnion) bool {
	if union == nil {
		return false
	}
	for _, v := range union.Variants {
		if hasOptionalFieldsInType(v.Type) {
			return true
		}
	}
	return false
}

func collectImportsFromType(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]bool) {
	for _, field := range t.Fields {
		collectImportsFromRef(&field.Type, formatMappings, importSet)
	}
	if t.Element != nil {
		collectImportsFromRef(t.Element, formatMappings, importSet)
	}
}

func collectImportsFromUnion(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]bool) {
	if t.Union != nil {
		for _, v := range t.Union.Variants {
			collectImportsFromType(v.Type, formatMappings, importSet)
		}
	}
}

func collectImportsFromRef(ref *ir.IRTypeRef, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]bool) {
	if ref == nil {
		return
	}
	if mapping, ok := formatMappings[ref.Format]; ok && mapping.Import != "" {
		importSet[mapping.Import] = true
	}
	if ref.Array != nil {
		collectImportsFromRef(ref.Array, formatMappings, importSet)
	}
	if ref.Map != nil {
		collectImportsFromRef(ref.Map, formatMappings, importSet)
	}
}

func makeGoTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping, optStyle OptionalStyle) func(*ir.IRTypeRef, bool) string {
	var goType func(*ir.IRTypeRef, bool) string
	goType = func(ref *ir.IRTypeRef, required bool) string {
		var baseType string
		var isSlice bool

		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			baseType = mapping.Type
			isSlice = strings.HasPrefix(baseType, "[]")
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "string"
				case ir.IRBuiltinInt:
					baseType = "int"
				case ir.IRBuiltinFloat:
					baseType = "float64"
				case ir.IRBuiltinBool:
					baseType = "bool"
				case ir.IRBuiltinAny:
					baseType = "interface{}"
				}
			} else if ref.Array != nil {
				baseType = "[]" + goType(ref.Array, true)
				isSlice = true
			} else if ref.Map != nil {
				baseType = "map[string]" + goType(ref.Map, true)
			} else if ref.Name != "" {
				baseType = ref.Name
			} else {
				baseType = "interface{}"
			}
		}

		if !required && baseType != "interface{}" && !isSlice && !strings.HasPrefix(baseType, "map") {
			switch optStyle {
			case OptionalStyleOpt:
				return "opt.Optional[" + baseType + "]"
			default:
				return "*" + baseType
			}
		}

		return baseType
	}
	return goType
}

func jsonTag(field ir.IRField) string {
	tag := field.JSONName
	if !field.Required {
		tag += ",omitempty"
	}
	return tag
}


const goTemplate = `package {{.Package}}
{{if or .HasUnion .Imports}}
import (
{{- if .HasUnion}}
	"bytes"
	"encoding/json"
	"fmt"
{{- end}}
{{- range .Imports}}
	"{{.}}"
{{- end}}
)
{{end}}
{{range .Types}}
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

{{define "struct"}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
type {{.Name}} struct {
{{- range .Fields}}
{{- if .Description}}
	{{comment .Description}}
{{- end}}
	{{.Name}} {{goType .Type .Required}} ` + "`" + `json:"{{jsonTag .}}"` + "`" + `
{{- end}}
}
{{end}}

{{define "alias"}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
{{- if .Element}}
type {{.Name}} = {{goType .Element true}}
{{- else}}
type {{.Name}} = interface{}
{{- end}}
{{end}}

{{define "enum"}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
{{- if isIntEnum .}}
type {{.Name}} int

const (
{{- range .EnumValues}}
{{- if not .IsNull}}
	{{toEnumKey $.Name .}} {{$.Name}} = {{.IntValue}}
{{- end}}
{{- end}}
)
{{- else}}
type {{.Name}} string

const (
{{- range .EnumValues}}
{{- if not .IsNull}}
	{{toEnumKey $.Name .}} {{$.Name}} = "{{.StringValue}}"
{{- end}}
{{- end}}
)
{{- end}}
{{end}}

{{define "union"}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
type {{.Union.InterfaceName}} interface {
	{{.Union.WrapperName}}Type() string
	is{{.Union.WrapperName}}()
}

type {{.Union.WrapperName}} struct {
	{{.Union.InterfaceName}}
}

func (w {{.Union.WrapperName}}) MarshalJSON() ([]byte, error) {
	if w.{{.Union.InterfaceName}} == nil {
		return []byte("null"), nil
	}
	return json.Marshal(w.{{.Union.InterfaceName}})
}

func (w *{{.Union.WrapperName}}) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		w.{{.Union.InterfaceName}} = nil
		return nil
	}

	var peek struct {
		Type string ` + "`" + `json:"{{.Union.DiscriminatorJSON}}"` + "`" + `
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("{{.Union.WrapperName}}: invalid JSON: %w", err)
	}
	if peek.Type == "" {
		return fmt.Errorf("{{.Union.WrapperName}}: missing discriminator field %q", "{{.Union.DiscriminatorJSON}}")
	}

	var v {{.Union.InterfaceName}}
	switch peek.Type {
{{- range .Union.Variants}}
	case "{{.ConstValue}}":
		v = &{{.Name}}{}
{{- end}}
	default:
		return fmt.Errorf("{{.Union.WrapperName}}: unknown type %q", peek.Type)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("{{.Union.WrapperName}}: invalid %q payload: %w", peek.Type, err)
	}

	w.{{.Union.InterfaceName}} = v
	return nil
}
{{range .Union.Variants}}
{{- if .Type.Description}}
{{comment .Type.Description}}
{{- end}}
type {{.Name}} struct {
{{- range .Type.Fields}}
{{- if .Description}}
	{{comment .Description}}
{{- end}}
	{{.Name}} {{goType .Type .Required}} ` + "`" + `json:"{{jsonTag .}}"` + "`" + `
{{- end}}
}

func ({{.Name}}) is{{$.Union.WrapperName}}() {}

func ({{.Name}}) {{$.Union.WrapperName}}Type() string { return "{{.ConstValue}}" }
{{end}}
{{end}}

{{define "simpleunion"}}
{{- if .Description}}
{{comment .Description}}
{{- end}}
type {{.Name}} = interface{}
{{end}}
`

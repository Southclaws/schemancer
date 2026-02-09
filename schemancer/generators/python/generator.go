package python

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// pythonReservedWords is the set of Python keywords that cannot be used as identifiers.
// Per PEP 8, the convention is to append a trailing underscore: class_ instead of class.
var pythonReservedWords = map[string]bool{
	"False": true, "None": true, "True": true, "and": true, "as": true,
	"assert": true, "async": true, "await": true, "break": true, "class": true,
	"continue": true, "def": true, "del": true, "elif": true, "else": true,
	"except": true, "finally": true, "for": true, "from": true, "global": true,
	"if": true, "import": true, "in": true, "is": true, "lambda": true,
	"nonlocal": true, "not": true, "or": true, "pass": true, "raise": true,
	"return": true, "try": true, "while": true, "with": true, "yield": true,
}

// safeSnake converts a name to snake_case and appends _ if it's a Python reserved word.
func safeSnake(name string) string {
	s := casing.ToSnakeCase(name)
	if pythonReservedWords[s] {
		return s + "_"
	}
	return s
}

// fieldAlias returns the JSON name as an alias if the snake_case field name is a
// Python reserved word, otherwise returns empty string.
func fieldAlias(name, jsonName string) string {
	s := casing.ToSnakeCase(name)
	if pythonReservedWords[s] {
		return jsonName
	}
	return ""
}

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in Python
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "bytes"},
	ir.IRFormatDateTime: {Type: "datetime", Import: "datetime"},
	ir.IRFormatDate:     {Type: "date", Import: "datetime"},
	ir.IRFormatUUID:     {Type: "UUID", Import: "uuid"},
	ir.IRFormatEmail:    {Type: "EmailStr", Import: "pydantic"},
	ir.IRFormatURI:      {Type: "AnyUrl", Import: "pydantic"},
}

// config holds Python-specific generator configuration
type config struct {
	// No specific options yet, but structure is ready for future options
}

// Option is a Python-specific generator option
type Option struct {
	apply func(*config)
}

// OptionValue implements generators.GeneratorOption
func (Option) OptionValue() string { return "python" }

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
	cfg := &config{}
	for _, opt := range genOpts {
		if pyOpt, ok := opt.(Option); ok {
			pyOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)
	inheritedFields := computeInheritedFields(data.Types)

	funcs := template.FuncMap{
		"snake":          casing.ToSnakeCase,
		"safeSnake":      safeSnake,
		"fieldAlias":     fieldAlias,
		"isInheritedField": func(typeName, fieldJSONName string) bool {
			if fields, ok := inheritedFields[typeName]; ok {
				return fields[fieldJSONName]
			}
			return false
		},
		"hasNonInheritedFields": func(typeName string, fields []ir.IRField) bool {
			inherited, ok := inheritedFields[typeName]
			if !ok {
				return len(fields) > 0
			}
			for _, f := range fields {
				if !inherited[f.JSONName] {
					return true
				}
			}
			return false
		},
		"pascal":         casing.ToPascalCase,
		"camel":          casing.ToCamelCase,
		"lower":          strings.ToLower,
		"upper":          strings.ToUpper,
		"pythonType":     makePythonTypeFunc(formatMappings),
		"pythonDefault":  makePythonDefaultFunc(),
		"pythonField":    makePythonFieldFunc(),
		"hasConstraints": hasConstraints,
		"comment":        formatComment,
		"fieldComment":   formatFieldComment,
		"hasPrefix":      strings.HasPrefix,
		"literalValue":   literalValue,
		"isIntEnum":      isIntEnum,
		"toEnumKey":      toEnumKey,
	}

	tmpl, err := template.New("python").Funcs(funcs).Parse(pythonTemplate)
	if err != nil {
		return nil, err
	}

	tplData := prepareTemplateData(data, formatMappings)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tplData); err != nil {
		return nil, err
	}

	return []generators.GeneratedFile{{
		Filename: "models.py",
		Content:  buf.Bytes(),
	}}, nil
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
	if len(lines) == 1 {
		return indent + `"""` + lines[0] + `"""`
	}
	var result []string
	result = append(result, indent+`"""`)
	for _, line := range lines {
		result = append(result, indent+line)
	}
	result = append(result, indent+`"""`)
	return strings.Join(result, "\n")
}

func formatFieldComment(description string) string {
	return formatCommentWithIndent(description, "    ")
}

func literalValue(s string) string {
	return `"` + s + `"`
}

type templateData struct {
	Imports      []importGroup
	Types        []ir.IRType
	HasUnion     bool
	HasOptional  bool
	HasList      bool
	HasDict      bool
	HasLiteral   bool
	HasAnnotated bool
}

type importGroup struct {
	Module string
	Names  []string
}

func prepareTemplateData(data *ir.IR, formatMappings map[ir.IRFormat]generators.FormatTypeMapping) templateData {
	importSet := make(map[string]map[string]bool)
	hasUnion := false
	hasOptional := false
	hasList := false
	hasDict := false
	hasLiteral := false
	hasAnnotated := false

	// Always need pydantic BaseModel
	addImport(importSet, "pydantic", "BaseModel")
	addImport(importSet, "pydantic", "ConfigDict")

	for _, t := range data.Types {
		switch t.Kind {
		case ir.IRKindDiscriminatedUnion:
			hasUnion = true
			hasAnnotated = true
			hasLiteral = true
			addImport(importSet, "pydantic", "Field")
			collectImportsFromUnion(t, formatMappings, importSet, &hasOptional, &hasList, &hasDict, &hasLiteral)
		case ir.IRKindAlias:
			if t.Element != nil {
				// Check if it's a primitive alias (needs RootModel)
				if t.Element.Builtin != ir.IRBuiltinNone || t.Element.Format != ir.IRFormatNone {
					addImport(importSet, "pydantic", "RootModel")
				}
				if t.Element.Array != nil {
					hasList = true
					addImport(importSet, "pydantic", "RootModel")
				}
				collectImportsFromRef(t.Element, formatMappings, importSet, &hasOptional, &hasList, &hasDict, &hasLiteral)
			}
		case ir.IRKindEnum:
			addImport(importSet, "enum", "Enum")
		default:
			collectImportsFromType(t, formatMappings, importSet, &hasOptional, &hasList, &hasDict, &hasLiteral)
		}
	}

	// Build import groups
	var imports []importGroup

	// typing imports
	var typingNames []string
	if hasUnion {
		typingNames = append(typingNames, "Union")
	}
	if hasOptional {
		// Using T | None syntax in Python 3.10+, but we still might need Optional for complex cases
	}
	if hasList {
		typingNames = append(typingNames, "List")
	}
	if hasDict {
		typingNames = append(typingNames, "Dict", "Any")
	}
	if hasLiteral {
		typingNames = append(typingNames, "Literal")
	}
	if hasAnnotated {
		typingNames = append(typingNames, "Annotated")
	}
	if len(typingNames) > 0 {
		sort.Strings(typingNames)
		imports = append(imports, importGroup{Module: "typing", Names: typingNames})
	}

	// Standard library imports (datetime, uuid, etc.)
	stdLibModules := []string{"datetime", "uuid", "ipaddress", "enum"}
	for _, mod := range stdLibModules {
		if names, ok := importSet[mod]; ok && len(names) > 0 {
			var nameList []string
			for name := range names {
				nameList = append(nameList, name)
			}
			sort.Strings(nameList)
			imports = append(imports, importGroup{Module: mod, Names: nameList})
		}
	}

	// Pydantic imports
	if names, ok := importSet["pydantic"]; ok && len(names) > 0 {
		var nameList []string
		for name := range names {
			nameList = append(nameList, name)
		}
		sort.Strings(nameList)
		imports = append(imports, importGroup{Module: "pydantic", Names: nameList})
	}

	return templateData{
		Imports:      imports,
		Types:        data.Types,
		HasUnion:     hasUnion,
		HasOptional:  hasOptional,
		HasList:      hasList,
		HasDict:      hasDict,
		HasLiteral:   hasLiteral,
		HasAnnotated: hasAnnotated,
	}
}

func addImport(importSet map[string]map[string]bool, module, name string) {
	if importSet[module] == nil {
		importSet[module] = make(map[string]bool)
	}
	importSet[module][name] = true
}

func collectImportsFromType(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]map[string]bool, hasOptional, hasList, hasDict, hasLiteral *bool) {
	for _, field := range t.Fields {
		if !field.Required {
			*hasOptional = true
		}
		if pythonReservedWords[casing.ToSnakeCase(field.Name)] {
			addImport(importSet, "pydantic", "Field")
		}
		collectImportsFromRef(&field.Type, formatMappings, importSet, hasOptional, hasList, hasDict, hasLiteral)
	}
	if t.Element != nil {
		collectImportsFromRef(t.Element, formatMappings, importSet, hasOptional, hasList, hasDict, hasLiteral)
	}
}

func collectImportsFromUnion(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]map[string]bool, hasOptional, hasList, hasDict, hasLiteral *bool) {
	if t.Union != nil {
		for _, v := range t.Union.Variants {
			*hasLiteral = true
			collectImportsFromType(v.Type, formatMappings, importSet, hasOptional, hasList, hasDict, hasLiteral)
		}
	}
}

func collectImportsFromRef(ref *ir.IRTypeRef, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, importSet map[string]map[string]bool, hasOptional, hasList, hasDict, hasLiteral *bool) {
	if ref == nil {
		return
	}

	// Check for format-specific imports
	if mapping, ok := formatMappings[ref.Format]; ok && mapping.Import != "" {
		addImport(importSet, mapping.Import, mapping.Type)
	}

	// Add Field import if constraints are present
	if ref.Constraints != nil {
		addImport(importSet, "pydantic", "Field")
	}

	if ref.Array != nil {
		*hasList = true
		collectImportsFromRef(ref.Array, formatMappings, importSet, hasOptional, hasList, hasDict, hasLiteral)
	}

	if ref.Map != nil {
		*hasDict = true
		collectImportsFromRef(ref.Map, formatMappings, importSet, hasOptional, hasList, hasDict, hasLiteral)
	}
}

func makePythonTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping) func(*ir.IRTypeRef, bool) string {
	var pythonType func(*ir.IRTypeRef, bool) string
	pythonType = func(ref *ir.IRTypeRef, required bool) string {
		var baseType string

		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			baseType = mapping.Type
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "str"
				case ir.IRBuiltinInt:
					baseType = "int"
				case ir.IRBuiltinFloat:
					baseType = "float"
				case ir.IRBuiltinBool:
					baseType = "bool"
				case ir.IRBuiltinAny:
					baseType = "Any"
				}
			} else if ref.Array != nil {
				baseType = "List[" + pythonType(ref.Array, true) + "]"
			} else if ref.Map != nil {
				baseType = "Dict[str, " + pythonType(ref.Map, true) + "]"
			} else if ref.Name != "" {
				baseType = ref.Name
			} else {
				baseType = "Any"
			}
		}

		if !required {
			return baseType + " | None"
		}

		return baseType
	}
	return pythonType
}

func makePythonDefaultFunc() func(bool) string {
	return func(required bool) string {
		if required {
			return ""
		}
		return " = None"
	}
}

// computeInheritedFields builds a map of type name -> set of field JSONNames
// that are inherited from a base type (i.e., exist in both the type and its BaseType).
func computeInheritedFields(types []ir.IRType) map[string]map[string]bool {
	// Index all type fields by name (including union variant types)
	typeFieldMap := make(map[string][]ir.IRField)
	for _, t := range types {
		typeFieldMap[t.Name] = t.Fields
		if t.Union != nil {
			for _, v := range t.Union.Variants {
				typeFieldMap[v.Name] = v.Type.Fields
			}
		}
	}

	result := make(map[string]map[string]bool)
	for _, t := range types {
		if t.BaseType != "" {
			markInherited(result, t.Name, t.BaseType, typeFieldMap)
		}
		if t.Union != nil {
			for _, v := range t.Union.Variants {
				if v.Type.BaseType != "" {
					markInherited(result, v.Name, v.Type.BaseType, typeFieldMap)
				}
			}
		}
	}
	return result
}

func markInherited(result map[string]map[string]bool, typeName, baseTypeName string, typeFieldMap map[string][]ir.IRField) {
	baseFieldNames := make(map[string]bool)
	for _, f := range typeFieldMap[baseTypeName] {
		baseFieldNames[f.JSONName] = true
	}
	inherited := make(map[string]bool)
	for _, f := range typeFieldMap[typeName] {
		if baseFieldNames[f.JSONName] {
			inherited[f.JSONName] = true
		}
	}
	result[typeName] = inherited
}

// hasConstraints returns true if the type ref has any validation constraints
func hasConstraints(ref *ir.IRTypeRef) bool {
	return ref != nil && ref.Constraints != nil
}

// isIntEnum returns true if the enum has an integer type
func isIntEnum(t ir.IRType) bool {
	return t.EnumType == ir.IRBuiltinInt
}

// toEnumKey converts an enum value to a valid Python enum key
func toEnumKey(v ir.IREnumValue) string {
	if v.IntValue != nil {
		// For integer enums, use VALUE_N format
		return "VALUE_" + v.StringValue
	}
	// For string enums, convert to UPPER_SNAKE_CASE
	return strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(v.StringValue, "-", "_"), " ", "_"))
}

// makePythonFieldFunc returns a function that generates Pydantic Field() calls with constraints.
// The alias parameter is non-empty when the field name is a Python reserved word and needs
// Field(alias="original_json_name") to maintain correct JSON serialization.
func makePythonFieldFunc() func(*ir.IRTypeRef, bool, string) string {
	return func(ref *ir.IRTypeRef, required bool, alias string) string {
		var parts []string

		if alias != "" {
			parts = append(parts, fmt.Sprintf(`alias="%s"`, alias))
		}

		if ref != nil && ref.Constraints != nil {
			c := ref.Constraints

			// String constraints
			if c.MinLength != nil {
				parts = append(parts, fmt.Sprintf("min_length=%d", *c.MinLength))
			}
			if c.MaxLength != nil {
				parts = append(parts, fmt.Sprintf("max_length=%d", *c.MaxLength))
			}
			if c.Pattern != "" {
				// Escape backslashes for Python string
				pattern := strings.ReplaceAll(c.Pattern, "\\", "\\\\")
				pattern = strings.ReplaceAll(pattern, `"`, `\"`)
				parts = append(parts, fmt.Sprintf(`pattern=r"%s"`, pattern))
			}

			// Numeric constraints
			if c.Minimum != nil {
				parts = append(parts, fmt.Sprintf("ge=%v", *c.Minimum))
			}
			if c.Maximum != nil {
				parts = append(parts, fmt.Sprintf("le=%v", *c.Maximum))
			}
			if c.ExclusiveMinimum != nil {
				parts = append(parts, fmt.Sprintf("gt=%v", *c.ExclusiveMinimum))
			}
			if c.ExclusiveMaximum != nil {
				parts = append(parts, fmt.Sprintf("lt=%v", *c.ExclusiveMaximum))
			}
			if c.MultipleOf != nil {
				parts = append(parts, fmt.Sprintf("multiple_of=%v", *c.MultipleOf))
			}

			// Array constraints (min_length/max_length work for lists too)
			if c.MinItems != nil {
				parts = append(parts, fmt.Sprintf("min_length=%d", *c.MinItems))
			}
			if c.MaxItems != nil {
				parts = append(parts, fmt.Sprintf("max_length=%d", *c.MaxItems))
			}
		}

		if len(parts) == 0 {
			if required {
				return ""
			}
			return " = None"
		}

		if !required {
			parts = append(parts, "default=None")
		}

		return " = Field(" + strings.Join(parts, ", ") + ")"
	}
}


const pythonTemplate = `from __future__ import annotations

{{range $i, $imp := .Imports -}}
from {{$imp.Module}} import {{range $j, $n := $imp.Names}}{{if $j}}, {{end}}{{$n}}{{end}}
{{end}}


{{range $i, $t := .Types -}}
{{if $i}}

{{end -}}
{{if eq .Kind "struct" -}}
{{template "class" .}}
{{- else if eq .Kind "alias" -}}
{{template "alias" .}}
{{- else if eq .Kind "enum" -}}
{{template "enum" .}}
{{- else if eq .Kind "discriminated_union" -}}
{{template "union" .}}
{{- else if eq .Kind "union" -}}
{{template "simpleunion" .}}
{{- end -}}
{{end}}
{{- define "class"}}
class {{.Name}}(BaseModel):
    model_config = ConfigDict(extra="forbid")
{{if hasNonInheritedFields .Name .Fields}}
{{- range .Fields}}
{{- if not (isInheritedField $.Name .JSONName)}}
{{- if .Description}}
{{fieldComment .Description}}
{{- end}}
    {{safeSnake .Name}}: {{pythonType .Type .Required}}{{pythonField .Type .Required (fieldAlias .Name .JSONName)}}
{{- end}}
{{- end}}
{{- else}}
    pass
{{- end}}
{{- end}}

{{- define "alias"}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if .Element}}
{{- if or .Element.Builtin .Element.Format .Element.Array}}
class {{.Name}}(RootModel[{{pythonType .Element true}}]):
    pass
{{- else}}
{{.Name}} = {{pythonType .Element true}}
{{- end}}
{{- else}}
{{.Name}} = Any
{{- end}}
{{- end}}

{{- define "enum"}}
{{- if .Description}}
{{comment .Description}}
{{end -}}
{{- if isIntEnum .}}
class {{.Name}}(int, Enum):
{{- range .EnumValues}}
{{- if not .IsNull}}
    {{toEnumKey .}} = {{.IntValue}}
{{- end}}
{{- end}}
{{- else}}
class {{.Name}}(str, Enum):
{{- range .Enum}}
    {{upper .}} = {{literalValue .}}
{{- end}}
{{- end}}
{{- end}}

{{- define "union"}}
{{- range $i, $v := .Union.Variants}}{{if $i}}

{{end -}}
{{$variantName := $v.Name}}
{{- if $v.Type.Description}}
{{comment $v.Type.Description}}
{{end -}}
{{- if $v.Type.BaseType}}
class {{$v.Name}}({{$v.Type.BaseType}}):
{{- else}}
class {{$v.Name}}(BaseModel):
    model_config = ConfigDict(extra="forbid")
{{- end}}
    {{safeSnake $.Union.DiscriminatorField}}: Literal[{{literalValue $v.ConstValue}}]{{with fieldAlias $.Union.DiscriminatorField $.Union.DiscriminatorJSON}} = Field(alias="{{.}}"){{end}}
{{- range $v.Type.Fields}}
{{- if ne .JSONName $.Union.DiscriminatorJSON}}
{{- if not (isInheritedField $variantName .JSONName)}}
{{- if .Description}}
{{fieldComment .Description}}
{{- end}}
    {{safeSnake .Name}}: {{pythonType .Type .Required}}{{pythonField .Type .Required (fieldAlias .Name .JSONName)}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{.Name}} = Annotated[
    Union[{{range $i, $v := .Union.Variants}}{{if $i}}, {{end}}{{$v.Name}}{{end}}],
    Field(discriminator="{{safeSnake .Union.DiscriminatorField}}"),
]
{{end}}
{{define "simpleunion"}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{.Name}} = Union[{{range $i, $v := .SimpleUnion.Variants}}{{if $i}}, {{end}}{{pythonType $v true}}{{end}}]
{{end}}
`

package java

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/casing"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in Java
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "byte[]"},
	ir.IRFormatDateTime: {Type: "java.time.OffsetDateTime"},
	ir.IRFormatDate:     {Type: "java.time.LocalDate"},
	ir.IRFormatUUID:     {Type: "java.util.UUID"},
	ir.IRFormatEmail:    {Type: "String"},
	ir.IRFormatURI:      {Type: "java.net.URI"},
}

// config holds Java-specific generator configuration
type config struct {
	packageName string
}

// Option is a Java-specific generator option
type Option struct {
	apply func(*config)
}

// OptionValue implements generators.GeneratorOption
func (Option) OptionValue() string { return "java" }

// WithPackageName sets the Java package name for generated code
func WithPackageName(name string) Option {
	return Option{apply: func(c *config) {
		c.packageName = name
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
		packageName: "generated",
	}
	for _, opt := range genOpts {
		if javaOpt, ok := opt.(Option); ok {
			javaOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)

	funcs := template.FuncMap{
		"pascal":    casing.ToPascalCase,
		"camel":     casing.ToCamelCase,
		"snake":     casing.ToSnakeCase,
		"kebab":     casing.ToKebabCase,
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"javaType":  makeJavaTypeFunc(formatMappings),
		"comment":   formatComment,
		"hasPrefix": strings.HasPrefix,
		"isIntEnum": isIntEnum,
		"toEnumKey": toEnumKey,
	}

	tmpl, err := template.New("java").Funcs(funcs).Parse(javaPerTypeTemplate)
	if err != nil {
		return nil, err
	}

	var files []generators.GeneratedFile

	for _, t := range data.Types {
		tplData := preparePerTypeData(cfg.packageName, t, formatMappings)

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, tplData); err != nil {
			return nil, err
		}

		files = append(files, generators.GeneratedFile{
			Filename: t.Name + ".java",
			Content:  buf.Bytes(),
		})
	}

	return files, nil
}

func formatComment(description string) string {
	if description == "" {
		return ""
	}
	description = strings.TrimRight(description, "\n")
	lines := strings.Split(description, "\n")
	if len(lines) == 1 {
		return "/** " + lines[0] + " */"
	}
	var result []string
	result = append(result, "/**")
	for _, line := range lines {
		result = append(result, " * "+line)
	}
	result = append(result, " */")
	return strings.Join(result, "\n")
}

// isIntEnum returns true if the enum has an integer type
func isIntEnum(t ir.IRType) bool {
	return t.EnumType == ir.IRBuiltinInt
}

// toEnumKey converts an enum value to a valid Java enum key
func toEnumKey(v ir.IREnumValue) string {
	if v.IntValue != nil {
		return "VALUE_" + v.StringValue
	}
	// For string enums, convert to UPPER_SNAKE_CASE
	return strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(v.StringValue, "-", "_"), " ", "_"))
}

type templateData struct {
	Package  string
	Imports  []string
	Types    []ir.IRType
	HasUnion bool
}

// perTypeData holds data for generating a single Java type/file
type perTypeData struct {
	Package  string
	Imports  []string
	Type     ir.IRType
	HasUnion bool
}

func preparePerTypeData(packageName string, t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping) perTypeData {
	importSet := make(map[string]bool)
	hasUnion := false

	// Always need Jackson annotations
	importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
	importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true

	if t.Kind == ir.IRKindDiscriminatedUnion {
		hasUnion = true
		importSet["com.fasterxml.jackson.annotation.JsonCreator"] = true
		importSet["com.fasterxml.jackson.annotation.JsonSubTypes"] = true
		importSet["com.fasterxml.jackson.annotation.JsonTypeInfo"] = true
		importSet["com.fasterxml.jackson.annotation.JsonTypeName"] = true
		collectImportsFromUnion(t, formatMappings, importSet)
	} else {
		collectImportsFromType(t, formatMappings, importSet)
	}

	var imports []string
	for imp := range importSet {
		imports = append(imports, imp)
	}
	sort.Strings(imports)

	return perTypeData{
		Package:  packageName,
		Imports:  imports,
		Type:     t,
		HasUnion: hasUnion,
	}
}

func prepareTemplateData(packageName string, data *ir.IR, formatMappings map[ir.IRFormat]generators.FormatTypeMapping) templateData {
	importSet := make(map[string]bool)
	hasUnion := false

	// Always need Jackson annotations
	importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
	importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true

	for _, t := range data.Types {
		if t.Kind == ir.IRKindDiscriminatedUnion {
			hasUnion = true
			importSet["com.fasterxml.jackson.annotation.JsonCreator"] = true
			importSet["com.fasterxml.jackson.annotation.JsonSubTypes"] = true
			importSet["com.fasterxml.jackson.annotation.JsonTypeInfo"] = true
			importSet["com.fasterxml.jackson.annotation.JsonTypeName"] = true
			collectImportsFromUnion(t, formatMappings, importSet)
		} else {
			collectImportsFromType(t, formatMappings, importSet)
		}
	}

	var imports []string
	for imp := range importSet {
		imports = append(imports, imp)
	}
	sort.Strings(imports)

	return templateData{
		Package:  packageName,
		Imports:  imports,
		Types:    data.Types,
		HasUnion: hasUnion,
	}
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

	// Check for format-specific imports
	if mapping, ok := formatMappings[ref.Format]; ok {
		addJavaImport(mapping.Type, importSet)
	}

	// Check for List import
	if ref.Array != nil {
		importSet["java.util.List"] = true
		collectImportsFromRef(ref.Array, formatMappings, importSet)
	}

	// Check for Map import
	if ref.Map != nil {
		importSet["java.util.Map"] = true
		collectImportsFromRef(ref.Map, formatMappings, importSet)
	}
}

func addJavaImport(typeName string, importSet map[string]bool) {
	// Add imports for java.time, java.util, java.net types
	if strings.HasPrefix(typeName, "java.time.") ||
		strings.HasPrefix(typeName, "java.util.") ||
		strings.HasPrefix(typeName, "java.net.") {
		importSet[typeName] = true
	}
}

func makeJavaTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping) func(*ir.IRTypeRef, bool) string {
	// javaTypeBoxed returns the boxed version of a type (for use in generics)
	var javaTypeBoxed func(*ir.IRTypeRef) string
	javaTypeBoxed = func(ref *ir.IRTypeRef) string {
		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			return getSimpleTypeName(mapping.Type)
		}

		if ref.Builtin != ir.IRBuiltinNone {
			switch ref.Builtin {
			case ir.IRBuiltinString:
				return "String"
			case ir.IRBuiltinInt:
				return "Integer"
			case ir.IRBuiltinFloat:
				return "Double"
			case ir.IRBuiltinBool:
				return "Boolean"
			case ir.IRBuiltinAny:
				return "Object"
			}
		} else if ref.Array != nil {
			return "List<" + javaTypeBoxed(ref.Array) + ">"
		} else if ref.Map != nil {
			return "Map<String, " + javaTypeBoxed(ref.Map) + ">"
		} else if ref.Name != "" {
			return ref.Name
		}
		return "Object"
	}

	var javaType func(*ir.IRTypeRef, bool) string
	javaType = func(ref *ir.IRTypeRef, required bool) string {
		var baseType string

		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			baseType = getSimpleTypeName(mapping.Type)
		}

		if baseType == "" {
			if ref.Builtin != ir.IRBuiltinNone {
				switch ref.Builtin {
				case ir.IRBuiltinString:
					baseType = "String"
				case ir.IRBuiltinInt:
					if required {
						baseType = "int"
					} else {
						baseType = "Integer"
					}
				case ir.IRBuiltinFloat:
					if required {
						baseType = "double"
					} else {
						baseType = "Double"
					}
				case ir.IRBuiltinBool:
					if required {
						baseType = "boolean"
					} else {
						baseType = "Boolean"
					}
				case ir.IRBuiltinAny:
					baseType = "Object"
				}
			} else if ref.Array != nil {
				// Use boxed types for generic type parameters
				baseType = "List<" + javaTypeBoxed(ref.Array) + ">"
			} else if ref.Map != nil {
				// Use boxed types for generic type parameters
				baseType = "Map<String, " + javaTypeBoxed(ref.Map) + ">"
			} else if ref.Name != "" {
				baseType = ref.Name
			} else {
				baseType = "Object"
			}
		}

		return baseType
	}
	return javaType
}

// getSimpleTypeName extracts the simple class name from a fully qualified name
func getSimpleTypeName(fqn string) string {
	if idx := strings.LastIndex(fqn, "."); idx != -1 {
		return fqn[idx+1:]
	}
	return fqn
}


const javaTemplate = `package {{.Package}};
{{range .Imports}}
import {{.}};
{{- end}}
{{- range $i, $t := .Types}}
{{- if eq .Kind "struct"}}
{{template "class" .}}
{{- else if eq .Kind "alias"}}
{{template "alias" .}}
{{- else if eq .Kind "enum"}}
{{template "enum" .}}
{{- else if eq .Kind "discriminated_union"}}
{{template "union" .}}
{{- else if eq .Kind "union"}}
{{template "simpleunion" .}}
{{- end}}
{{- end}}

{{- define "class"}}
{{if .Description}}
{{comment .Description}}
{{- end}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} {
{{- range .Fields}}

{{- if .Description}}
    {{comment .Description}}
{{- end}}
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}})
    public {{javaType .Type .Required}} {{camel .Name}};
{{- end}}
}
{{- end}}

{{- define "alias" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{- if .Element}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} extends {{javaType .Element true}} {}
{{- else}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} {}
{{- end}}
{{- end}}

{{- define "enum" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{- if isIntEnum .}}
public enum {{.Name}} {
{{- range $i, $v := .EnumValues}}
{{- if not $v.IsNull}}
{{- if $i}},{{end}}
    {{toEnumKey $v}}({{$v.IntValue}})
{{- end}}
{{- end}};

    private final int value;

    {{.Name}}(int value) {
        this.value = value;
    }

    @JsonProperty
    public int getValue() {
        return value;
    }
}
{{- else}}
public enum {{.Name}} {
{{- range $i, $v := .Enum}}
{{- if $i}},{{end}}
    {{upper $v}}("{{$v}}")
{{- end}};

    private final String value;

    {{.Name}}(String value) {
        this.value = value;
    }

    @JsonProperty
    public String getValue() {
        return value;
    }
}
{{- end}}
{{- end}}

{{- define "union" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonTypeInfo(
    use = JsonTypeInfo.Id.NAME,
    include = JsonTypeInfo.As.PROPERTY,
    property = "{{.Union.DiscriminatorJSON}}",
    visible = true
)
@JsonSubTypes({
{{- range $i, $v := .Union.Variants}}
{{- if $i}},{{end}}
    @JsonSubTypes.Type(value = {{$v.Name}}.class, name = "{{$v.ConstValue}}")
{{- end}}
})
public sealed interface {{.Name}} permits {{range $i, $v := .Union.Variants}}{{if $i}}, {{end}}{{$v.Name}}{{end}} {
    String {{camel .Union.DiscriminatorField}}();
}
{{range .Union.Variants}}
{{if .Type.Description}}
{{comment .Type.Description}}
{{end}}
@JsonTypeName("{{.ConstValue}}")
public record {{.Name}}(
    @JsonProperty(value = "{{$.Union.DiscriminatorJSON}}") String {{camel $.Union.DiscriminatorField}}{{range .Type.Fields}}{{if ne .JSONName $.Union.DiscriminatorJSON}},
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}}) {{javaType .Type .Required}} {{camel .Name}}{{end}}{{end}}
) implements {{$.Name}} {
    @JsonCreator
    public {{.Name}} {}
}
{{- end}}
{{- end}}

{{- define "simpleunion" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} extends Object {}
{{- end}}
`

// javaPerTypeTemplate generates a single Java file for one type
const javaPerTypeTemplate = `package {{.Package}};
{{range .Imports}}
import {{.}};
{{- end}}
{{- with .Type}}
{{- if eq .Kind "struct"}}
{{template "class" .}}
{{- else if eq .Kind "alias"}}
{{template "alias" .}}
{{- else if eq .Kind "enum"}}
{{template "enum" .}}
{{- else if eq .Kind "discriminated_union"}}
{{template "union" .}}
{{- else if eq .Kind "union"}}
{{template "simpleunion" .}}
{{- end}}
{{- end}}

{{- define "class"}}
{{if .Description}}
{{comment .Description}}
{{- end}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} {
{{- range .Fields}}

{{- if .Description}}
    {{comment .Description}}
{{- end}}
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}})
    public {{javaType .Type .Required}} {{camel .Name}};
{{- end}}
}
{{- end}}

{{- define "alias" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{- if .Element}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} extends {{javaType .Element true}} {}
{{- else}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} {}
{{- end}}
{{- end}}

{{- define "enum" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
{{- if isIntEnum .}}
public enum {{.Name}} {
{{- range $i, $v := .EnumValues}}
{{- if not $v.IsNull}}
{{- if $i}},{{end}}
    {{toEnumKey $v}}({{$v.IntValue}})
{{- end}}
{{- end}};

    private final int value;

    {{.Name}}(int value) {
        this.value = value;
    }

    @JsonProperty
    public int getValue() {
        return value;
    }
}
{{- else}}
public enum {{.Name}} {
{{- range $i, $v := .Enum}}
{{- if $i}},{{end}}
    {{upper $v}}("{{$v}}")
{{- end}};

    private final String value;

    {{.Name}}(String value) {
        this.value = value;
    }

    @JsonProperty
    public String getValue() {
        return value;
    }
}
{{- end}}
{{- end}}

{{- define "union" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonTypeInfo(
    use = JsonTypeInfo.Id.NAME,
    include = JsonTypeInfo.As.PROPERTY,
    property = "{{.Union.DiscriminatorJSON}}",
    visible = true
)
@JsonSubTypes({
{{- range $i, $v := .Union.Variants}}
{{- if $i}},{{end}}
    @JsonSubTypes.Type(value = {{$v.Name}}.class, name = "{{$v.ConstValue}}")
{{- end}}
})
public sealed interface {{.Name}} permits {{range $i, $v := .Union.Variants}}{{if $i}}, {{end}}{{$v.Name}}{{end}} {
    String {{camel .Union.DiscriminatorField}}();
}
{{range .Union.Variants}}
{{if .Type.Description}}
{{comment .Type.Description}}
{{end}}
@JsonTypeName("{{.ConstValue}}")
public record {{.Name}}(
    @JsonProperty(value = "{{$.Union.DiscriminatorJSON}}") String {{camel $.Union.DiscriminatorField}}{{range .Type.Fields}}{{if ne .JSONName $.Union.DiscriminatorJSON}},
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}}) {{javaType .Type .Required}} {{camel .Name}}{{end}}{{end}}
) implements {{$.Name}} {
    @JsonCreator
    public {{.Name}} {}
}
{{- end}}
{{- end}}

{{- define "simpleunion" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} extends Object {}
{{- end}}
`

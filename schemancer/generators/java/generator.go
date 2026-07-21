package java

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

// javaReservedWords contains all Java language keywords that cannot be used as identifiers.
var javaReservedWords = map[string]bool{
	"abstract": true, "assert": true, "boolean": true, "break": true,
	"byte": true, "case": true, "catch": true, "char": true,
	"class": true, "const": true, "continue": true, "default": true,
	"do": true, "double": true, "else": true, "enum": true,
	"extends": true, "final": true, "finally": true, "float": true,
	"for": true, "goto": true, "if": true, "implements": true,
	"import": true, "instanceof": true, "int": true, "interface": true,
	"long": true, "native": true, "new": true, "package": true,
	"private": true, "protected": true, "public": true, "return": true,
	"short": true, "static": true, "strictfp": true, "super": true,
	"switch": true, "synchronized": true, "this": true, "throw": true,
	"throws": true, "transient": true, "try": true, "void": true,
	"volatile": true, "while": true,
}

// javaConflictingGetters contains getter method names that conflict with
// methods on java.lang.Object and would cause compilation errors.
var javaConflictingGetters = map[string]bool{
	"getClass": true,
}

// DefaultFormatMappings provides sensible defaults for JSON Schema formats in Java
var DefaultFormatMappings = map[ir.IRFormat]generators.FormatTypeMapping{
	ir.IRFormatByte:     {Type: "byte[]"},
	ir.IRFormatDateTime: {Type: "java.time.OffsetDateTime"},
	ir.IRFormatDate:     {Type: "java.time.LocalDate"},
	ir.IRFormatUUID:     {Type: "java.util.UUID"},
	ir.IRFormatEmail:    {Type: "String"},
	ir.IRFormatURI:      {Type: "java.net.URI"},
}

// PropertyInclusion controls Jackson @JsonInclude behavior on generated classes.
type PropertyInclusion string

const (
	PropertyInclusionNonNull  PropertyInclusion = "non_null"
	PropertyInclusionNonEmpty PropertyInclusion = "non_empty"
	PropertyInclusionAlways   PropertyInclusion = "always"
)

// config holds Java-specific generator configuration
type config struct {
	packageName        string
	accessors          bool
	propertyInclusion  PropertyInclusion
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

// WithAccessors enables getter/setter generation (fields become private)
func WithAccessors(enabled bool) Option {
	return Option{apply: func(c *config) {
		c.accessors = enabled
	}}
}

// WithPropertyInclusion sets Jackson @JsonInclude behavior for generated classes.
// Valid values: "non_null" (default), "non_empty", "always".
func WithPropertyInclusion(inclusion PropertyInclusion) Option {
	return Option{apply: func(c *config) {
		c.propertyInclusion = inclusion
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
		packageName:       "generated",
		propertyInclusion: PropertyInclusionNonNull,
	}
	for _, opt := range genOpts {
		if javaOpt, ok := opt.(Option); ok {
			javaOpt.apply(cfg)
		}
	}

	formatMappings := g.getFormatMappings(opts)
	typeIndex := buildTypeIndex(data.Types)

	// Build type-kind lookup for constructor generation
	typeKinds := make(map[string]ir.IRTypeKind)
	for _, t := range data.Types {
		typeKinds[t.Name] = t.Kind
	}

	funcs := template.FuncMap{
		"pascal":           casing.ToPascalCase,
		"camel":            casing.ToCamelCase,
		"snake":            casing.ToSnakeCase,
		"kebab":            casing.ToKebabCase,
		"lower":            strings.ToLower,
		"upper":            strings.ToUpper,
		"javaType":         makeJavaTypeFunc(formatMappings, typeIndex),
		"javaInit":         makeJavaInitFunc(typeIndex),
		"javaDefault":      makeJavaDefaultFunc(),
		"javaFieldName":    makeJavaFieldNameFunc(),
		"javaConstructor":  makeJavaConstructorFunc(typeKinds),
		"javaJsonInclude":  makeJavaJsonIncludeFunc(),
		"javaGetter":       makeJavaGetterFunc(formatMappings, typeIndex),
		"javaSetter":       makeJavaSetterFunc(formatMappings, typeIndex),
		"comment":          formatComment,
		"hasPrefix":        strings.HasPrefix,
		"isIntEnum":        isIntEnum,
		"toEnumKey":        toEnumKey,
		"hasDefault":       hasDefault,
		"hasFieldDefaults": hasFieldDefaults,
	}

	tmpl, err := template.New("java").Funcs(funcs).Parse(javaPerTypeTemplate)
	if err != nil {
		return nil, err
	}

	variantTmpl, err := template.New("java-variant").Funcs(funcs).Parse(javaVariantTemplate)
	if err != nil {
		return nil, err
	}

	var files []generators.GeneratedFile

	for _, t := range data.Types {
		if t.Kind == ir.IRKindAlias && isInlinableAlias(t, typeIndex) {
			continue
		}

		if t.Kind == ir.IRKindDiscriminatedUnion && t.Union != nil {
			// Interface file
			tplData := preparePerTypeData(cfg.packageName, t, formatMappings, typeIndex, cfg.accessors, cfg.propertyInclusion)
			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, tplData); err != nil {
				return nil, err
			}
			files = append(files, generators.GeneratedFile{
				Filename: t.Name + ".java",
				Content:  buf.Bytes(),
			})

			// One file per variant
			for _, v := range t.Union.Variants {
				vData := prepareVariantData(cfg.packageName, v, t, formatMappings, typeIndex)
				var vBuf bytes.Buffer
				if err := variantTmpl.Execute(&vBuf, vData); err != nil {
					return nil, err
				}
				files = append(files, generators.GeneratedFile{
					Filename: v.Name + ".java",
					Content:  vBuf.Bytes(),
				})
			}
			continue
		}

		tplData := preparePerTypeData(cfg.packageName, t, formatMappings, typeIndex, cfg.accessors, cfg.propertyInclusion)

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

// hasDefault returns true if the field has a default value
func hasDefault(f ir.IRField) bool {
	return f.Default != nil
}

// hasFieldDefaults returns true if any field in the type has a default value
func hasFieldDefaults(t ir.IRType) bool {
	for _, f := range t.Fields {
		if f.Default != nil {
			return true
		}
	}
	return false
}

type templateData struct {
	Package  string
	Imports  []string
	Types    []ir.IRType
	HasUnion bool
}

// perTypeData holds data for generating a single Java type/file
type perTypeData struct {
	Package           string
	Imports           []string
	Type              ir.IRType
	HasUnion          bool
	Accessors         bool
	PropertyInclusion PropertyInclusion
}

// variantData holds data for generating a single discriminated union variant file
type variantData struct {
	Package            string
	Imports            []string
	Variant            ir.IRVariant
	UnionName          string
	DiscriminatorField string
	DiscriminatorJSON  string
}

func preparePerTypeData(packageName string, t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, accessors bool, propertyInclusion PropertyInclusion) perTypeData {
	importSet := make(map[string]bool)
	hasUnion := false

	if t.Kind == ir.IRKindEnum {
		// Enums only need JsonValue for serialization/deserialization
		importSet["com.fasterxml.jackson.annotation.JsonValue"] = true
	} else if t.Kind == ir.IRKindDiscriminatedUnion {
		// Interface file only needs interface-level annotations
		hasUnion = true
		importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
		importSet["com.fasterxml.jackson.annotation.JsonSubTypes"] = true
		importSet["com.fasterxml.jackson.annotation.JsonTypeInfo"] = true
	} else {
		// Structs need JsonIgnoreProperties and JsonProperty
		importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
		importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true
		if propertyInclusion != PropertyInclusionAlways {
			importSet["com.fasterxml.jackson.annotation.JsonInclude"] = true
		}
		collectImportsFromType(t, formatMappings, typeIndex, importSet)

		// Check if any field has a default value — if so, add JsonSetter and Nulls imports
		for _, field := range t.Fields {
			if field.Default != nil {
				importSet["com.fasterxml.jackson.annotation.JsonSetter"] = true
				importSet["com.fasterxml.jackson.annotation.Nulls"] = true
				break
			}
		}
	}

	var imports []string
	for imp := range importSet {
		imports = append(imports, imp)
	}
	sort.Strings(imports)

	return perTypeData{
		Package:           packageName,
		Imports:           imports,
		Type:              t,
		HasUnion:          hasUnion,
		Accessors:         accessors,
		PropertyInclusion: propertyInclusion,
	}
}

func prepareVariantData(packageName string, v ir.IRVariant, union ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType) variantData {
	importSet := make(map[string]bool)
	importSet["com.fasterxml.jackson.annotation.JsonCreator"] = true
	importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true
	importSet["com.fasterxml.jackson.annotation.JsonTypeName"] = true
	// Records don't initialize fields, so skip ArrayList/HashMap imports
	for _, field := range v.Type.Fields {
		collectImportsFromRefForAlias(&field.Type, formatMappings, typeIndex, importSet)
	}

	var imports []string
	for imp := range importSet {
		imports = append(imports, imp)
	}
	sort.Strings(imports)

	return variantData{
		Package:            packageName,
		Imports:            imports,
		Variant:            v,
		UnionName:          union.Name,
		DiscriminatorField: union.Union.DiscriminatorField,
		DiscriminatorJSON:  union.Union.DiscriminatorJSON,
	}
}

func prepareTemplateData(packageName string, data *ir.IR, formatMappings map[ir.IRFormat]generators.FormatTypeMapping) templateData {
	importSet := make(map[string]bool)
	hasUnion := false
	typeIndex := buildTypeIndex(data.Types)

	for _, t := range data.Types {
		if t.Kind == ir.IRKindEnum {
			importSet["com.fasterxml.jackson.annotation.JsonValue"] = true
		} else if t.Kind == ir.IRKindDiscriminatedUnion {
			hasUnion = true
			importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
			importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true
			importSet["com.fasterxml.jackson.annotation.JsonCreator"] = true
			importSet["com.fasterxml.jackson.annotation.JsonSubTypes"] = true
			importSet["com.fasterxml.jackson.annotation.JsonTypeInfo"] = true
			importSet["com.fasterxml.jackson.annotation.JsonTypeName"] = true
			collectImportsFromUnion(t, formatMappings, typeIndex, importSet)
		} else {
			importSet["com.fasterxml.jackson.annotation.JsonIgnoreProperties"] = true
			importSet["com.fasterxml.jackson.annotation.JsonProperty"] = true
			collectImportsFromType(t, formatMappings, typeIndex, importSet)

			for _, field := range t.Fields {
				if field.Default != nil {
					importSet["com.fasterxml.jackson.annotation.JsonSetter"] = true
					importSet["com.fasterxml.jackson.annotation.Nulls"] = true
					break
				}
			}
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

func collectImportsFromType(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, importSet map[string]bool) {
	for _, field := range t.Fields {
		collectImportsFromRefInner(&field.Type, formatMappings, typeIndex, importSet, field.Required)
	}
	if t.Element != nil {
		// Alias types don't initialize fields, so skip ArrayList/HashMap imports
		collectImportsFromRefForAlias(t.Element, formatMappings, typeIndex, importSet)
	}
}

func collectImportsFromUnion(t ir.IRType, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, importSet map[string]bool) {
	if t.Union != nil {
		for _, v := range t.Union.Variants {
			collectImportsFromType(v.Type, formatMappings, typeIndex, importSet)
		}
	}
}

func collectImportsFromRef(ref *ir.IRTypeRef, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, importSet map[string]bool) {
	collectImportsFromRefInner(ref, formatMappings, typeIndex, importSet, true)
}

// collectImportsFromRefForAlias collects imports without ArrayList/HashMap (aliases don't initialize fields)
func collectImportsFromRefForAlias(ref *ir.IRTypeRef, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, importSet map[string]bool) {
	collectImportsFromRefInner(ref, formatMappings, typeIndex, importSet, false)
}

func collectImportsFromRefInner(ref *ir.IRTypeRef, formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType, importSet map[string]bool, includeInitImports bool) {
	if ref == nil {
		return
	}

	ref = resolveInlinedRef(ref, typeIndex)

	// Check for format-specific imports
	if mapping, ok := formatMappings[ref.Format]; ok {
		if mapping.Import != "" {
			importSet[mapping.Import] = true
		} else {
			addJavaImport(mapping.Type, importSet)
		}
	}

	// Check for List import
	if ref.Array != nil {
		importSet["java.util.List"] = true
		if includeInitImports {
			importSet["java.util.ArrayList"] = true
		}
		collectImportsFromRefInner(ref.Array, formatMappings, typeIndex, importSet, includeInitImports)
	}

	// Check for Map import
	if ref.Map != nil {
		importSet["java.util.Map"] = true
		if includeInitImports {
			importSet["java.util.HashMap"] = true
		}
		collectImportsFromRefInner(ref.Map, formatMappings, typeIndex, importSet, includeInitImports)
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

func buildTypeIndex(types []ir.IRType) map[string]ir.IRType {
	index := make(map[string]ir.IRType, len(types))
	for _, t := range types {
		index[t.Name] = t
	}
	return index
}

func isInlinableElement(ref *ir.IRTypeRef) bool {
	if ref == nil {
		return false
	}
	if ref.Builtin == ir.IRBuiltinAny {
		return false
	}
	if ref.Builtin != ir.IRBuiltinNone {
		return true
	}
	if ref.Format != ir.IRFormatNone {
		return true
	}
	if ref.Array != nil || ref.Map != nil {
		return true
	}
	return false
}

func resolvedAliasElement(ref *ir.IRTypeRef, index map[string]ir.IRType, visited map[string]bool) *ir.IRTypeRef {
	if ref == nil {
		return nil
	}
	if isInlinableElement(ref) {
		return ref
	}
	if ref.Name == "" {
		return nil
	}
	if visited[ref.Name] {
		return nil
	}
	visited[ref.Name] = true

	t, ok := index[ref.Name]
	if !ok || t.Kind != ir.IRKindAlias || t.Element == nil {
		return nil
	}
	return resolvedAliasElement(t.Element, index, visited)
}

func isInlinableAlias(t ir.IRType, index map[string]ir.IRType) bool {
	if t.Kind != ir.IRKindAlias || t.Element == nil {
		return false
	}
	return resolvedAliasElement(t.Element, index, make(map[string]bool)) != nil
}

func resolveInlinedRef(ref *ir.IRTypeRef, index map[string]ir.IRType) *ir.IRTypeRef {
	if ref == nil {
		return nil
	}

	resolved := resolvedAliasElement(ref, index, make(map[string]bool))
	if resolved == nil {
		return ref
	}

	constraints := ref.Constraints
	if constraints == nil {
		constraints = resolved.Constraints
	}

	return &ir.IRTypeRef{
		Builtin:     resolved.Builtin,
		Format:      resolved.Format,
		Array:       resolved.Array,
		Map:         resolved.Map,
		Name:        resolved.Name,
		Nullable:    ref.Nullable,
		Constraints: constraints,
	}
}

func makeJavaTypeFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType) func(*ir.IRTypeRef, bool) string {
	// javaTypeBoxed returns the boxed version of a type (for use in generics)
	var javaTypeBoxed func(*ir.IRTypeRef) string
	javaTypeBoxed = func(ref *ir.IRTypeRef) string {
		ref = resolveInlinedRef(ref, typeIndex)
		// Check format first
		if mapping, ok := formatMappings[ref.Format]; ok {
			return getSimpleTypeName(mapping.Type)
		}

		if ref.Builtin != ir.IRBuiltinNone {
			switch ref.Builtin {
			case ir.IRBuiltinString:
				return "String"
			case ir.IRBuiltinInt:
				return "Long"
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
		ref = resolveInlinedRef(ref, typeIndex)
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
						baseType = "long"
					} else {
						baseType = "Long"
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

func makeJavaInitFunc(typeIndex map[string]ir.IRType) func(ir.IRField) string {
	return func(field ir.IRField) string {
		if !field.Required {
			return ""
		}
		ref := resolveInlinedRef(&field.Type, typeIndex)
		if ref.Array != nil {
			return " = new ArrayList<>()"
		}
		if ref.Map != nil {
			return " = new HashMap<>()"
		}
		return ""
	}
}

// makeJavaDefaultFunc returns a template function that renders a field's default value as a Java literal.
// When a field has a default, this takes priority over javaInit (collection initialization).
func makeJavaDefaultFunc() func(ir.IRField) string {
	return func(field ir.IRField) string {
		if field.Default == nil {
			return ""
		}

		raw := field.Default.RawValue

		switch field.Default.Builtin {
		case ir.IRBuiltinInt:
			return " = " + raw + "L"
		case ir.IRBuiltinFloat:
			return " = " + raw
		case ir.IRBuiltinBool:
			return " = " + raw
		case ir.IRBuiltinString:
			// raw is already JSON-quoted (e.g., "\"hello\"")
			return " = " + raw
		default:
			return ""
		}
	}
}

// safeJavaFieldName resolves the Java field name for a field.
// Uses x-java-name extension if present, otherwise falls back to camelCase
// with reserved word escaping (appends _ suffix).
func safeJavaFieldName(field ir.IRField) string {
	if name, ok := field.Extensions["x-java-name"]; ok {
		return name
	}
	name := casing.ToCamelCase(field.Name)
	if javaReservedWords[name] {
		return name + "_"
	}
	return name
}

// makeJavaFieldNameFunc returns a template function that resolves the Java field name.
func makeJavaFieldNameFunc() func(ir.IRField) string {
	return safeJavaFieldName
}

// makeJavaConstructorFunc returns a template function that generates a no-arg constructor
// initializing required struct-typed fields with new instances.
func makeJavaConstructorFunc(typeKinds map[string]ir.IRTypeKind) func(ir.IRType) string {
	return func(t ir.IRType) string {
		var refs []ir.IRField
		for _, f := range t.Fields {
			if !f.Required {
				continue
			}
			if f.Type.Name != "" && f.Type.Array == nil && f.Type.Map == nil && f.Type.Builtin == ir.IRBuiltinNone {
				if kind, ok := typeKinds[f.Type.Name]; ok && kind == ir.IRKindStruct {
					refs = append(refs, f)
				}
			}
		}
		if len(refs) == 0 {
			return ""
		}

		var sb strings.Builder
		sb.WriteString("\n\n    public " + t.Name + "() {\n")
		for _, f := range refs {
			fieldName := safeJavaFieldName(f)
			sb.WriteString("        this." + fieldName + " = new " + f.Type.Name + "();\n")
		}
		sb.WriteString("    }")
		return sb.String()
	}
}

func makeJavaJsonIncludeFunc() func(PropertyInclusion) string {
	return func(inclusion PropertyInclusion) string {
		switch inclusion {
		case PropertyInclusionNonEmpty:
			return "\n@JsonInclude(JsonInclude.Include.NON_EMPTY)"
		case PropertyInclusionAlways:
			return ""
		default:
			return "\n@JsonInclude(JsonInclude.Include.NON_NULL)"
		}
	}
}

// makeJavaGetterFunc returns a template function that generates a getter method for a field.
func makeJavaGetterFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType) func(ir.IRField) string {
	javaType := makeJavaTypeFunc(formatMappings, typeIndex)
	return func(field ir.IRField) string {
		fieldName := safeJavaFieldName(field)
		typeName := javaType(&field.Type, field.Required)
		pascalName := casing.ToPascalCase(field.Name)
		prefix := "get"
		if typeName == "boolean" || typeName == "Boolean" {
			prefix = "is"
		}
		methodName := prefix + pascalName
		if javaConflictingGetters[methodName] {
			methodName = methodName + "_"
		}
		return fmt.Sprintf("    public %s %s() {\n        return %s;\n    }", typeName, methodName, fieldName)
	}
}

// makeJavaSetterFunc returns a template function that generates a setter method for a field.
func makeJavaSetterFunc(formatMappings map[ir.IRFormat]generators.FormatTypeMapping, typeIndex map[string]ir.IRType) func(ir.IRField) string {
	javaType := makeJavaTypeFunc(formatMappings, typeIndex)
	return func(field ir.IRField) string {
		fieldName := safeJavaFieldName(field)
		typeName := javaType(&field.Type, field.Required)
		pascalName := casing.ToPascalCase(field.Name)
		setterName := "set" + pascalName
		if javaConflictingGetters["get"+pascalName] {
			setterName = setterName + "_"
		}
		return fmt.Sprintf("    public void %s(%s %s) {\n        this.%s = %s;\n    }", setterName, typeName, fieldName, fieldName, fieldName)
	}
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
{{- if hasDefault .}}
    @JsonSetter(nulls = Nulls.SKIP)
{{- end}}
    public {{javaType .Type .Required}} {{javaFieldName .}}{{if hasDefault .}}{{javaDefault .}}{{else}}{{javaInit .}}{{end}};
{{- end}}
{{- javaConstructor .}}
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

    private final long value;

    {{.Name}}(long value) {
        this.value = value;
    }

    @JsonValue
    public long getValue() {
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

    @JsonValue
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
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}}) {{javaType .Type .Required}} {{javaFieldName .}}{{end}}{{end}}
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
{{template "class" $}}
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
{{- if .Type.Description}}
{{comment .Type.Description}}
{{- end}}
@JsonIgnoreProperties(ignoreUnknown = true){{javaJsonInclude $.PropertyInclusion}}
public class {{.Type.Name}} {
{{- range .Type.Fields}}

{{- if .Description}}
    {{comment .Description}}
{{- end}}
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}})
{{- if hasDefault .}}
    @JsonSetter(nulls = Nulls.SKIP)
{{- end}}
    {{if $.Accessors}}private{{else}}public{{end}} {{javaType .Type .Required}} {{javaFieldName .}}{{if hasDefault .}}{{javaDefault .}}{{else}}{{javaInit .}}{{end}};
{{- end}}
{{- if $.Accessors}}
{{- range .Type.Fields}}

{{javaGetter .}}

{{javaSetter .}}
{{- end}}
{{- end}}
{{- javaConstructor .Type}}
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

    private final long value;

    {{.Name}}(long value) {
        this.value = value;
    }

    @JsonValue
    public long getValue() {
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

    @JsonValue
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
{{- end}}

{{- define "simpleunion" -}}
{{if .Description}}
{{comment .Description}}
{{end}}
@JsonIgnoreProperties(ignoreUnknown = true)
public class {{.Name}} extends Object {}
{{- end}}
`

// javaVariantTemplate generates a single variant record file
const javaVariantTemplate = `package {{.Package}};
{{range .Imports}}
import {{.}};
{{- end}}
{{if .Variant.Type.Description}}
{{comment .Variant.Type.Description}}
{{end}}
@JsonTypeName("{{.Variant.ConstValue}}")
public record {{.Variant.Name}}(
    @JsonProperty(value = "{{.DiscriminatorJSON}}") String {{camel .DiscriminatorField}}{{range .Variant.Type.Fields}}{{if ne .JSONName $.DiscriminatorJSON}},
    @JsonProperty(value = "{{.JSONName}}"{{if .Required}}, required = true{{end}}) {{javaType .Type .Required}} {{javaFieldName .}}{{end}}{{end}}
) implements {{.UnionName}} {
    @JsonCreator
    public {{.Variant.Name}} {}
}
`

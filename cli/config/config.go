package config

// Describes how to map a JSON Schema "format" value to a concrete type in a target language. This allows customizing the type used for fields with a "format" keyword (e.g. "uuid", "date-time", "email", "uri") or defining mappings for custom application-specific formats.
type FormatMapping struct {
	// The import or package path needed to use the specified type. This is added to the import block of the generated file. For example, "github.com/google/uuid" in Go, "java.util.UUID" in Java, or "uuid" in Python. Leave empty if the type is a builtin that requires no import.
	Import string `json:"import"`
	// The fully qualified type name in the target language. For example, "uuid.UUID" in Go, "string" in TypeScript, "java.util.UUID" in Java, or "uuid.UUID" in Python. This is the type that will be used in the generated code for any field with the corresponding JSON Schema format.
	Type string `json:"type"`
}

// Configuration for Go code generation. Controls the output directory, package name, how optional fields are represented, and custom type mappings for JSON Schema format values.
type GolangConfig struct {
	// Custom type mappings for JSON Schema "format" values. By default, schemancer maps common formats to standard Go types (e.g. "uuid" to github.com/google/uuid.UUID, "date-time" to time.Time). Use this to override defaults or add mappings for custom formats. The map key is the JSON Schema format string (e.g. "uuid", "date-time", "email") and the value describes the Go type and import path to use.
	FormatMappings map[string]FormatMapping `json:"format_mappings,omitempty"`
	// Controls how optional (non-required) fields are represented in the generated Go structs. Supported values are "pointer" (the default), which uses Go pointer types (e.g. *string, *int), and "opt", which uses the github.com/Southclaws/opt library's Optional[T] generic type. Can be overridden by the --optional-style CLI flag.
	OptionalStyle *string `json:"optional_style,omitempty"`
	// The output directory path where generated Go files will be written. The directory will be created if it does not exist. The generated file will be named after the package (e.g. "models.go" for package "models"). This field is required for the language to be included in multi-language generation mode.
	Output *string `json:"output,omitempty"`
	// The Go package name for the generated source file. This appears in the "package" declaration at the top of the generated file. Defaults to "generated" if not specified. Can be overridden by the --package CLI flag.
	Package *string `json:"package,omitempty"`
}

// Configuration for Java code generation. Controls the output directory, package name, accessor method generation, and custom format type mappings. Each top-level type is generated as a separate Java file with Jackson annotations.
type JavaConfig struct {
	// When true, generates getter and setter methods for all fields instead of using public fields. The fields become private and are accessed through getFieldName()/setFieldName() methods following standard JavaBean conventions. Defaults to false.
	Accessors *bool `json:"accessors,omitempty"`
	// Custom type mappings for JSON Schema "format" values. By default, schemancer maps common formats to standard Java types (e.g. "uuid" to java.util.UUID, "date-time" to java.time.OffsetDateTime). Use this to override defaults or add mappings for custom formats. The map key is the JSON Schema format string and the value describes the Java type and import path.
	FormatMappings map[string]FormatMapping `json:"format_mappings,omitempty"`
	// The output directory path where generated Java files will be written. The directory will be created if it does not exist. Each top-level type produces a separate .java file. This field is required for the language to be included in multi-language generation mode.
	Output *string `json:"output,omitempty"`
	// The Java package name for the generated classes. This appears in the "package" declaration at the top of each generated file. Defaults to "generated" if not specified. Can be overridden by the --package CLI flag.
	Package *string `json:"package,omitempty"`
}

// Configuration for Python code generation. Controls the output directory and custom format type mappings. The generated code uses Pydantic v2 BaseModel classes with full type annotations.
type PythonConfig struct {
	// Custom type mappings for JSON Schema "format" values. By default, schemancer maps common formats to standard Python types (e.g. "uuid" to uuid.UUID, "date-time" to datetime.datetime). Use this to override defaults or add mappings for custom formats. The map key is the JSON Schema format string and the value describes the Python type and import path.
	FormatMappings map[string]FormatMapping `json:"format_mappings,omitempty"`
	// The output directory path where the generated Python file will be written. The directory will be created if it does not exist. This field is required for the language to be included in multi-language generation mode.
	Output *string `json:"output,omitempty"`
}

// Configuration for TypeScript code generation. Controls the output directory, output filename, optional field representation, branded primitive types, and custom format type mappings.
type TypeScriptConfig struct {
	// When true, primitive type aliases are generated as branded types instead of plain type aliases. For example, instead of "type UserId = string", it generates a branded type that prevents accidental assignment between different string-based types. This provides stronger type safety at the cost of slightly more verbose usage. Defaults to false. Can be overridden by the --branded-primitives CLI flag.
	BrandedPrimitives *bool `json:"branded_primitives,omitempty"`
	// The filename for the generated TypeScript file. Defaults to "types.ts" if not specified. Use this to customize the output filename, for example "models.ts" or "schema.ts".
	Filename *string `json:"filename,omitempty"`
	// Custom type mappings for JSON Schema "format" values. By default, schemancer maps common formats to standard TypeScript types. Use this to override defaults or add mappings for custom formats. The map key is the JSON Schema format string and the value describes the TypeScript type and optional import path.
	FormatMappings map[string]FormatMapping `json:"format_mappings,omitempty"`
	// When true, optional fields use "| null" instead of "?" (undefined) for their optional representation. For example, a non-required string field becomes "field: string | null" instead of "field?: string". This is useful for APIs that distinguish between missing fields and null values. Defaults to false. Can be overridden by the --null-optional CLI flag.
	NullOptional *bool `json:"null_optional,omitempty"`
	// The output directory path where the generated TypeScript file will be written. The directory will be created if it does not exist. This field is required for the language to be included in multi-language generation mode.
	Output *string `json:"output,omitempty"`
}

// Configuration for TypeScript Zod code generation. Controls the output directory and custom format type mappings. The generated code produces Zod v4 schemas with z.infer<> type exports and full constraint support.
type TypeScriptZodConfig struct {
	// Custom type mappings for JSON Schema "format" values. Use this to override the default Zod type mappings or add custom format handlers. The map key is the JSON Schema format string and the value describes the Zod type and optional import path.
	FormatMappings map[string]FormatMapping `json:"format_mappings,omitempty"`
	// The output directory path where the generated Zod schema file will be written. The directory will be created if it does not exist. This field is required for the language to be included in multi-language generation mode.
	Output *string `json:"output,omitempty"`
}

// The schemancer configuration file structure. This file is typically named schemancer.yaml and placed in the root of your project alongside your JSON Schema definitions. It controls how code is generated for each target language, including output paths, package names, language-specific options, and custom format type mappings. Each top-level key corresponds to a supported target language. Only languages with a configuration block will be included when running schemancer in multi-language mode (i.e. without explicit language and output arguments on the command line).
type Config struct {
	// Go-specific generation options. When present with an output path set, schemancer will generate Go source files. The generated code uses standard encoding/json struct tags and idiomatic Go naming conventions.
	Golang *GolangConfig `json:"golang,omitempty"`
	// Java-specific generation options. When present with an output path set, schemancer will generate Java class files. Each top-level type is emitted as a separate .java file with Jackson annotations for JSON serialization/deserialization.
	Java *JavaConfig `json:"java,omitempty"`
	// Python-specific generation options. When present with an output path set, schemancer will generate Python source files using Pydantic v2 BaseModel classes with full type annotations and validation support.
	Python *PythonConfig `json:"python,omitempty"`
	// TypeScript-specific generation options. When present with an output path set, schemancer will generate TypeScript type definitions. The generated code produces interfaces and type aliases suitable for use with any TypeScript project.
	Typescript *TypeScriptConfig `json:"typescript,omitempty"`
	// TypeScript Zod-specific generation options. When present with an output path set, schemancer will generate Zod v4 schema definitions with inferred TypeScript types. The generated code produces z.object() schemas with full runtime validation support, including constraints like min/max length, numeric bounds, and array limits.
	TypescriptZod *TypeScriptZodConfig `json:"typescript-zod,omitempty"`
}

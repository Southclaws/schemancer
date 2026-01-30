package generators

import (
	"github.com/Southclaws/schemancer/schemancer/ir"
)

type Language string

const (
	LanguageGo            Language = "golang"
	LanguageTypeScript    Language = "typescript"
	LanguageTypeScriptZod Language = "typescript-zod"
	LanguageJava          Language = "java"
	LanguagePython        Language = "python"
)

// GeneratedFile represents a single generated output file
type GeneratedFile struct {
	// Filename is the suggested filename for this output (e.g., "User.java", "types.ts")
	Filename string
	// Content is the generated file content
	Content []byte
}

type Generator interface {
	Generate(ir *ir.IR, opts GeneratorOptions, genOpts ...GeneratorOption) ([]GeneratedFile, error)
}

type GlobalOptions struct {
	Language          Language
	FormatTypeMapping map[ir.IRFormat]FormatTypeMapping
}

// GeneratorOptions provides generic options for all code generators
type GeneratorOptions struct {
	// FormatMappings allows customizing how JSON Schema formats map to target types
	// Key is the IRFormat, value is a target-specific type configuration
	FormatMappings map[ir.IRFormat]FormatTypeMapping
}

// FormatTypeMapping describes how to map a JSON Schema format to a target language type
type FormatTypeMapping struct {
	Type   string // The type name in the target language
	Import string // The import/package path needed (if any)
}

// GeneratorOption is a marker interface for language-specific generator options.
// Each generator implementation defines its own option types that satisfy this interface.
type GeneratorOption interface {
	OptionValue() string
}

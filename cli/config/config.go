package config

import (
	"os"

	"github.com/goccy/go-yaml"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

// Config represents the schemancer.yaml configuration file structure.
type Config struct {
	Golang        *GolangConfig        `yaml:"golang,omitempty"`
	TypeScript    *TypeScriptConfig    `yaml:"typescript,omitempty"`
	TypeScriptZod *TypeScriptZodConfig `yaml:"typescript-zod,omitempty"`
	Java          *JavaConfig          `yaml:"java,omitempty"`
	Python        *PythonConfig        `yaml:"python,omitempty"`
}

// GolangConfig contains Go-specific generation options.
type GolangConfig struct {
	Output         string                    `yaml:"output,omitempty"`
	Package        string                    `yaml:"package,omitempty"`
	OptionalStyle  string                    `yaml:"optional_style,omitempty"`
	FormatMappings map[string]*FormatMapping `yaml:"format_mappings,omitempty"`
}

// TypeScriptConfig contains TypeScript-specific generation options.
type TypeScriptConfig struct {
	Output            string                    `yaml:"output,omitempty"`
	NullOptional      bool                      `yaml:"null_optional,omitempty"`
	BrandedPrimitives bool                      `yaml:"branded_primitives,omitempty"`
	FormatMappings    map[string]*FormatMapping `yaml:"format_mappings,omitempty"`
}

// TypeScriptZodConfig contains TypeScript Zod-specific generation options.
type TypeScriptZodConfig struct {
	Output         string                    `yaml:"output,omitempty"`
	FormatMappings map[string]*FormatMapping `yaml:"format_mappings,omitempty"`
}

// JavaConfig contains Java-specific generation options.
type JavaConfig struct {
	Output         string                    `yaml:"output,omitempty"`
	Package        string                    `yaml:"package,omitempty"`
	FormatMappings map[string]*FormatMapping `yaml:"format_mappings,omitempty"`
}

// PythonConfig contains Python-specific generation options.
type PythonConfig struct {
	Output         string                    `yaml:"output,omitempty"`
	FormatMappings map[string]*FormatMapping `yaml:"format_mappings,omitempty"`
}

// LanguageOutput represents a language and its output path from config.
type LanguageOutput struct {
	Language generators.Language
	Output   string
}

// GetConfiguredLanguages returns all languages that have an output path configured.
func (c *Config) GetConfiguredLanguages() []LanguageOutput {
	if c == nil {
		return nil
	}

	var languages []LanguageOutput

	if c.Golang != nil && c.Golang.Output != "" {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageGo,
			Output:   c.Golang.Output,
		})
	}

	if c.TypeScript != nil && c.TypeScript.Output != "" {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageTypeScript,
			Output:   c.TypeScript.Output,
		})
	}

	if c.TypeScriptZod != nil && c.TypeScriptZod.Output != "" {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageTypeScriptZod,
			Output:   c.TypeScriptZod.Output,
		})
	}

	if c.Java != nil && c.Java.Output != "" {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageJava,
			Output:   c.Java.Output,
		})
	}

	if c.Python != nil && c.Python.Output != "" {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguagePython,
			Output:   c.Python.Output,
		})
	}

	return languages
}

// FormatMapping describes how to map a JSON Schema format to a target language type.
type FormatMapping struct {
	Type   string `yaml:"type"`
	Import string `yaml:"import,omitempty"`
}

// Load reads and parses a config file from the given path.
// Returns nil config (not an error) if the file doesn't exist.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// GetFormatMappings converts the config's format mappings to the generator's format.
func (c *Config) GetFormatMappings(language generators.Language) map[ir.IRFormat]generators.FormatTypeMapping {
	if c == nil {
		return nil
	}

	var mappings map[string]*FormatMapping

	switch language {
	case generators.LanguageGo:
		if c.Golang != nil {
			mappings = c.Golang.FormatMappings
		}
	case generators.LanguageTypeScript:
		if c.TypeScript != nil {
			mappings = c.TypeScript.FormatMappings
		}
	case generators.LanguageTypeScriptZod:
		if c.TypeScriptZod != nil {
			mappings = c.TypeScriptZod.FormatMappings
		}
	case generators.LanguageJava:
		if c.Java != nil {
			mappings = c.Java.FormatMappings
		}
	case generators.LanguagePython:
		if c.Python != nil {
			mappings = c.Python.FormatMappings
		}
	}

	if len(mappings) == 0 {
		return nil
	}

	result := make(map[ir.IRFormat]generators.FormatTypeMapping, len(mappings))
	for format, mapping := range mappings {
		result[ir.IRFormat(format)] = generators.FormatTypeMapping{
			Type:   mapping.Type,
			Import: mapping.Import,
		}
	}

	return result
}

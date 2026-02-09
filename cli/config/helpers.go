package config

import (
	"os"

	"github.com/goccy/go-yaml"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

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

	if c.Golang != nil && c.Golang.Output != nil {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageGo,
			Output:   *c.Golang.Output,
		})
	}

	if c.Typescript != nil && c.Typescript.Output != nil {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageTypeScript,
			Output:   *c.Typescript.Output,
		})
	}

	if c.TypescriptZod != nil && c.TypescriptZod.Output != nil {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageTypeScriptZod,
			Output:   *c.TypescriptZod.Output,
		})
	}

	if c.Java != nil && c.Java.Output != nil {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguageJava,
			Output:   *c.Java.Output,
		})
	}

	if c.Python != nil && c.Python.Output != nil {
		languages = append(languages, LanguageOutput{
			Language: generators.LanguagePython,
			Output:   *c.Python.Output,
		})
	}

	return languages
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

	var mappings map[string]FormatMapping

	switch language {
	case generators.LanguageGo:
		if c.Golang != nil {
			mappings = c.Golang.FormatMappings
		}
	case generators.LanguageTypeScript:
		if c.Typescript != nil {
			mappings = c.Typescript.FormatMappings
		}
	case generators.LanguageTypeScriptZod:
		if c.TypescriptZod != nil {
			mappings = c.TypescriptZod.FormatMappings
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

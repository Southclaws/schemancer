package schemancer

import (
	"fmt"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/golang"
	"github.com/Southclaws/schemancer/schemancer/generators/java"
	"github.com/Southclaws/schemancer/schemancer/generators/python"
	"github.com/Southclaws/schemancer/schemancer/generators/typescript"
	typescriptzod "github.com/Southclaws/schemancer/schemancer/generators/typescript-zod"

	"github.com/google/jsonschema-go/jsonschema"
)

var Generators = map[generators.Language]generators.Generator{
	generators.LanguageGo:            &golang.Generator{},
	generators.LanguageTypeScript:    &typescript.Generator{},
	generators.LanguageTypeScriptZod: &typescriptzod.Generator{},
	generators.LanguageJava:          &java.Generator{},
	generators.LanguagePython:        &python.Generator{},
}

func Generate(schema *jsonschema.Schema, opts generators.GlobalOptions, genOpts ...generators.GeneratorOption) ([]generators.GeneratedFile, error) {
	if opts.Language == "" {
		opts.Language = generators.LanguageGo
	}

	gen, ok := Generators[opts.Language]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", opts.Language)
	}

	irData, err := SchemaToIR(schema)
	if err != nil {
		return nil, err
	}

	return gen.Generate(irData, generators.GeneratorOptions{
		FormatMappings: opts.FormatTypeMapping,
	}, genOpts...)
}

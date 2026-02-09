package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/spf13/cobra"

	"github.com/Southclaws/schemancer/cli/config"
	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/golang"
	"github.com/Southclaws/schemancer/schemancer/generators/java"
	"github.com/Southclaws/schemancer/schemancer/generators/typescript"
	"github.com/Southclaws/schemancer/schemancer/loader"
)

var (
	// Global options
	configFile string

	// Go options
	goPackage       string
	goOptionalStyle string

	// TypeScript options
	tsNullOptional bool
	tsBranded      bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "schemancer <schema> [language] [output]",
		Short: "JSON Schema Code Generator",
		Long: `Generate type-safe code from JSON Schema definitions.

When called with just a schema file, generates code for all languages
configured with output paths in schemancer.yaml.

When called with schema, language, and output, generates code for that
specific language to that output path.`,
		Args: cobra.RangeArgs(1, 3),
		RunE: run,
	}

	// Global flags
	rootCmd.Flags().StringVar(&configFile, "config", "schemancer.yaml", "Path to config file")

	// Go flags
	rootCmd.Flags().StringVar(&goPackage, "package", "", "Go: package name for generated code")
	rootCmd.Flags().StringVar(&goOptionalStyle, "optional-style", "", "Go: how to represent optional fields (pointer, opt)")

	// TypeScript flags
	rootCmd.Flags().BoolVar(&tsNullOptional, "null-optional", false, "TypeScript: use null instead of undefined for optional fields")
	rootCmd.Flags().BoolVar(&tsBranded, "branded-primitives", false, "TypeScript: use branded types for primitive type aliases")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	schemaFile := args[0]

	// Load config file
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	schema, err := loader.FromFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	// Determine mode: single language or multi-language from config
	if len(args) == 3 {
		// Explicit: schemancer <schema> <language> <output>
		return generateSingle(cmd, cfg, schema, args[1], args[2])
	}

	if len(args) == 1 {
		// Config-driven: schemancer <schema>
		return generateFromConfig(cmd, cfg, schema)
	}

	return fmt.Errorf("invalid arguments: provide either 1 argument (schema) or 3 arguments (schema language output)")
}

func generateFromConfig(cmd *cobra.Command, cfg *config.Config, schema *jsonschema.Schema) error {
	languages := cfg.GetConfiguredLanguages()
	if len(languages) == 0 {
		return fmt.Errorf("no languages configured with output paths in config file")
	}

	for _, lang := range languages {
		if err := generateSingle(cmd, cfg, schema, string(lang.Language), lang.Output); err != nil {
			return err
		}
	}

	return nil
}

func generateSingle(cmd *cobra.Command, cfg *config.Config, schema *jsonschema.Schema, language, outputPath string) error {
	genOpts, err := getGeneratorOptions(cmd, cfg, language)
	if err != nil {
		return err
	}

	formatMappings := cfg.GetFormatMappings(generators.Language(language))

	files, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language:          generators.Language(language),
		FormatTypeMapping: formatMappings,
	}, genOpts...)
	if err != nil {
		return fmt.Errorf("failed to generate %s: %w", language, err)
	}

	if outputPath == "-" {
		// Output to stdout - concatenate all files
		for i, f := range files {
			if i > 0 {
				fmt.Println() // Separator between files
			}
			fmt.Print(string(f.Content))
		}
	} else {
		// Always treat outputPath as a directory
		if err := os.MkdirAll(outputPath, 0o755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		for _, f := range files {
			outFile := filepath.Join(outputPath, f.Filename)
			if err := os.WriteFile(outFile, f.Content, 0o644); err != nil {
				return fmt.Errorf("failed to write %s: %w", outFile, err)
			}
			fmt.Fprintf(os.Stderr, "Generated %s\n", outFile)
		}
	}

	return nil
}

func getGeneratorOptions(cmd *cobra.Command, cfg *config.Config, language string) ([]generators.GeneratorOption, error) {
	var genOpts []generators.GeneratorOption

	switch language {
	case "golang":
		// Resolve package name: CLI flag > config > default
		pkg := "generated"
		if cfg != nil && cfg.Golang != nil && cfg.Golang.Package != nil {
			pkg = *cfg.Golang.Package
		}
		if goPackage != "" {
			pkg = goPackage
		}
		genOpts = append(genOpts, golang.WithPackageName(pkg))

		// Resolve optional style: CLI flag > config > default
		optStyle := "pointer"
		if cfg != nil && cfg.Golang != nil && cfg.Golang.OptionalStyle != nil {
			optStyle = *cfg.Golang.OptionalStyle
		}
		if goOptionalStyle != "" {
			optStyle = goOptionalStyle
		}
		genOpts = append(genOpts, golang.WithOptionalStyle(golang.OptionalStyle(optStyle)))

	case "typescript":
		// Resolve null_optional: CLI flag > config > default
		nullOpt := false
		if cfg != nil && cfg.Typescript != nil {
			nullOpt = *cfg.Typescript.NullOptional
		}
		if cmd.Flags().Changed("null-optional") {
			nullOpt = tsNullOptional
		}
		genOpts = append(genOpts, typescript.WithNullForOptional(nullOpt))

		// Resolve branded_primitives: CLI flag > config > default
		branded := false
		if cfg != nil && cfg.Typescript != nil {
			branded = *cfg.Typescript.BrandedPrimitives
		}
		if cmd.Flags().Changed("branded-primitives") {
			branded = tsBranded
		}
		genOpts = append(genOpts, typescript.WithBrandedTypes(branded))

		// Resolve filename: config > default ("types.ts")
		if cfg != nil && cfg.Typescript != nil && cfg.Typescript.Filename != nil {
			genOpts = append(genOpts, typescript.WithFilename(*cfg.Typescript.Filename))
		}

	case "java":
		// Resolve package name: CLI flag > config > default
		pkg := "generated"
		if cfg != nil && cfg.Java != nil && cfg.Java.Package != nil {
			pkg = *cfg.Java.Package
		}
		if goPackage != "" {
			pkg = goPackage
		}
		genOpts = append(genOpts, java.WithPackageName(pkg))

		// Resolve accessors: config > default (false)
		if cfg != nil && cfg.Java != nil && cfg.Java.Accessors != nil && *cfg.Java.Accessors {
			genOpts = append(genOpts, java.WithAccessors(true))
		}

	case "python":
		// Python has no special options yet

	case "typescript-zod":
		// TypeScript Zod has no special options yet

	default:
		return nil, fmt.Errorf("unsupported language: %s (supported: golang, typescript, typescript-zod, java, python)", language)
	}

	return genOpts, nil
}

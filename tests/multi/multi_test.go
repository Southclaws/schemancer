package multi_test

import (
	"testing"

	"github.com/Southclaws/schemancer/cli/config"
	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/golang"
	"github.com/Southclaws/schemancer/schemancer/generators/java"
	"github.com/Southclaws/schemancer/schemancer/generators/typescript"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"

	"github.com/stretchr/testify/require"
)

func TestMultiLanguageGeneration(t *testing.T) {
	cfg, err := config.Load("schemancer.yaml")
	require.NoError(t, err, "failed to load config")

	schema, err := loader.FromFile("schema.yaml")
	require.NoError(t, err, "failed to load schema")

	languages := cfg.GetConfiguredLanguages()
	require.Len(t, languages, 5, "expected 5 configured languages")

	// Generate Go
	goFiles, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language:          generators.LanguageGo,
		FormatTypeMapping: cfg.GetFormatMappings(generators.LanguageGo),
	}, golang.WithPackageName("multi"))
	require.NoError(t, err, "failed to generate Go")
	testutil.WriteAndCompareMultipleFiles(t, goFiles, "generated/golang", "expected/golang")

	// Generate TypeScript
	tsFiles, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language:          generators.LanguageTypeScript,
		FormatTypeMapping: cfg.GetFormatMappings(generators.LanguageTypeScript),
	}, typescript.WithNullForOptional(false))
	require.NoError(t, err, "failed to generate TypeScript")
	testutil.WriteAndCompareMultipleFiles(t, tsFiles, "generated/typescript", "expected/typescript")

	// Generate TypeScript Zod
	tsZodFiles, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language:          generators.LanguageTypeScriptZod,
		FormatTypeMapping: cfg.GetFormatMappings(generators.LanguageTypeScriptZod),
	})
	require.NoError(t, err, "failed to generate TypeScript Zod")
	testutil.WriteAndCompareMultipleFiles(t, tsZodFiles, "generated/typescript-zod", "expected/typescript-zod")

	// Generate Java
	javaFiles, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language: generators.LanguageJava,
	}, java.WithPackageName("multi"))
	require.NoError(t, err, "failed to generate Java")
	testutil.WriteAndCompareMultipleFiles(t, javaFiles, "generated/java", "expected/java")

	// Generate Python
	pythonFiles, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language: generators.LanguagePython,
	})
	require.NoError(t, err, "failed to generate Python")
	testutil.WriteAndCompareMultipleFiles(t, pythonFiles, "generated/python", "expected/python")
}

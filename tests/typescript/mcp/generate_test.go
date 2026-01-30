package mcp_test

import (
	"os"
	"testing"

	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"
)

func TestMCP(t *testing.T) {
	schema, err := loader.FromFile("schema.json")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	files, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language: generators.LanguageTypeScript,
	})
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}
	generated := testutil.GetSingleFile(t, files)

	if err := os.WriteFile("output.ts", generated, 0o644); err != nil {
		t.Fatalf("failed to write output: %v", err)
	}

	expected, err := os.ReadFile("expected.ts")
	if err != nil {
		t.Fatalf("failed to read expected output: %v", err)
	}

	testutil.CompareGenerated(t, generated, expected)
}

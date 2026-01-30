package integer_enums_test

import (
	"os"
	"testing"

	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/golang"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"
)

func TestIntegerEnums(t *testing.T) {
	schema, err := loader.FromFile("schema.yaml")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	files, err := schemancer.Generate(schema, generators.GlobalOptions{}, golang.WithPackageName("integer_enums"))
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}
	generated := testutil.GetSingleFile(t, files)

	if err := os.WriteFile("output.go", generated, 0o644); err != nil {
		t.Fatalf("failed to write output: %v", err)
	}

	expected, err := os.ReadFile("expected_test.go")
	if err != nil {
		t.Fatalf("failed to read expected output: %v", err)
	}

	testutil.CompareGenerated(t, generated, expected)
}

package discriminated_union_test

import (
	"os"
	"testing"

	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/detect"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/generators/golang"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"
)

func TestDiscriminatedUnion(t *testing.T) {
	schema, err := loader.FromFile("schema.yaml")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	files, err := schemancer.Generate(schema, generators.GlobalOptions{}, golang.WithPackageName("discriminated_union"))
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

func TestDiscriminatedUnionDetection(t *testing.T) {
	schema, err := loader.FromFile("schema.yaml")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	result, err := detect.DiscriminatedUnion(schema)
	if err != nil {
		t.Fatalf("failed to detect discriminated union: %v", err)
	}

	if result == nil {
		t.Fatal("expected discriminated union to be detected")
	}

	if result.DiscriminatorField != "type" {
		t.Errorf("expected discriminator field 'type', got %q", result.DiscriminatorField)
	}

	if len(result.Variants) != 3 {
		t.Errorf("expected 3 variants, got %d", len(result.Variants))
	}
}

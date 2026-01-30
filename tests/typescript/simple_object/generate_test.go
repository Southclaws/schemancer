package simple_object_test

import (
	"os"
	"testing"

	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleObject(t *testing.T) {
	schema, err := loader.FromFile("schema.yaml")
	require.NoError(t, err, "failed to load schema")

	files, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language: generators.LanguageTypeScript,
	})
	require.NoError(t, err, "failed to generate")
	generated := testutil.GetSingleFile(t, files)

	if err := os.WriteFile("output.ts", generated, 0o644); err != nil {
		t.Fatalf("failed to write output: %v", err)
	}

	expected, err := os.ReadFile("expected.ts")
	require.NoError(t, err, "failed to read expected output")

	assert.Equal(t, string(expected), string(generated), "generated code does not match expected")
}

package testutil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Southclaws/schemancer/schemancer/generators"
)

// GetSingleFile extracts the content from a single-file generation result.
// Fails the test if there isn't exactly one file.
func GetSingleFile(t *testing.T, files []generators.GeneratedFile) []byte {
	t.Helper()
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
	return files[0].Content
}

// CompareGenerated compares generated code with expected code, stripping the
// package line since generated code uses a different package name than tests.
func CompareGenerated(t *testing.T, generated, expected []byte) {
	t.Helper()

	genBody := stripPackageLine(string(generated))
	expBody := stripPackageLine(string(expected))

	assert.Equal(t, expBody, genBody, "generated code does not match expected")
}

func stripPackageLine(s string) string {
	parts := strings.SplitN(s, "\n", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return s
}

// WriteAndCompareMultipleFiles writes all generated files to the output directory
// and compares each one to the corresponding expected file.
// Expected files should be in the expectedDir with matching filenames.
func WriteAndCompareMultipleFiles(t *testing.T, files []generators.GeneratedFile, outputDir, expectedDir string) {
	t.Helper()

	// Write all generated files
	for _, f := range files {
		outPath := filepath.Join(outputDir, f.Filename)
		if err := os.WriteFile(outPath, f.Content, 0o644); err != nil {
			t.Fatalf("failed to write %s: %v", outPath, err)
		}
	}

	// Compare each generated file to expected
	for _, f := range files {
		expectedPath := filepath.Join(expectedDir, f.Filename)
		expected, err := os.ReadFile(expectedPath)
		if err != nil {
			t.Fatalf("failed to read expected file %s: %v", expectedPath, err)
		}

		assert.Equal(t, string(expected), string(f.Content), "generated %s does not match expected", f.Filename)
	}
}

// ConcatenateFiles combines all generated files into a single byte slice.
// Useful for comparing multi-file output against a single concatenated expected file.
func ConcatenateFiles(files []generators.GeneratedFile) []byte {
	var result []byte
	for i, f := range files {
		result = append(result, f.Content...)
		if i < len(files)-1 {
			result = append(result, '\n')
		}
	}
	return result
}

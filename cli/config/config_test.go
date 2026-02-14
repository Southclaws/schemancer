package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Southclaws/schemancer/schemancer/generators"
	"github.com/Southclaws/schemancer/schemancer/ir"
)

func TestLoad_FullConfig(t *testing.T) {
	cfg, err := Load("testdata/full.yaml")
	require.NoError(t, err)
	require.NotNil(t, cfg)

	// Check golang config
	require.NotNil(t, cfg.Golang)
	assert.Equal(t, "mymodels", *cfg.Golang.Package)
	assert.Equal(t, "opt", *cfg.Golang.OptionalStyle)
	require.Len(t, cfg.Golang.FormatMappings, 2)
	assert.Equal(t, "CustomUUID", cfg.Golang.FormatMappings["uuid"].Type)
	assert.Equal(t, "example.com/uuid", cfg.Golang.FormatMappings["uuid"].Import)

	// Check typescript config
	require.NotNil(t, cfg.Typescript)
	assert.True(t, *cfg.Typescript.NullOptional)
	assert.True(t, *cfg.Typescript.BrandedPrimitives)
	require.Len(t, cfg.Typescript.FormatMappings, 2)
	assert.Equal(t, "Date", cfg.Typescript.FormatMappings["date-time"].Type)
}

func TestLoad_GolangOnly(t *testing.T) {
	cfg, err := Load("testdata/golang_only.yaml")
	require.NoError(t, err)
	require.NotNil(t, cfg)

	require.NotNil(t, cfg.Golang)
	assert.Equal(t, "generated", *cfg.Golang.Package)
	assert.Equal(t, "pointer", *cfg.Golang.OptionalStyle)

	assert.Nil(t, cfg.Typescript)
}

func TestLoad_FileNotFound(t *testing.T) {
	cfg, err := Load("testdata/nonexistent.yaml")
	require.NoError(t, err)
	assert.Nil(t, cfg)
}

func TestGetFormatMappings_Golang(t *testing.T) {
	cfg, err := Load("testdata/full.yaml")
	require.NoError(t, err)

	mappings := cfg.GetFormatMappings(generators.LanguageGo)
	require.NotNil(t, mappings)
	require.Len(t, mappings, 2)

	assert.Equal(t, generators.FormatTypeMapping{
		Type:   "CustomUUID",
		Import: "example.com/uuid",
	}, mappings[ir.IRFormat("uuid")])

	assert.Equal(t, generators.FormatTypeMapping{
		Type:   "time.Time",
		Import: "time",
	}, mappings[ir.IRFormat("date-time")])
}

func TestGetFormatMappings_TypeScript(t *testing.T) {
	cfg, err := Load("testdata/full.yaml")
	require.NoError(t, err)

	mappings := cfg.GetFormatMappings(generators.LanguageTypeScript)
	require.NotNil(t, mappings)
	require.Len(t, mappings, 2)

	assert.Equal(t, generators.FormatTypeMapping{
		Type: "Date",
	}, mappings[ir.IRFormat("date-time")])

	assert.Equal(t, generators.FormatTypeMapping{
		Type: "UUID",
	}, mappings[ir.IRFormat("uuid")])
}

func TestGetFormatMappings_NilConfig(t *testing.T) {
	var cfg *Config
	mappings := cfg.GetFormatMappings(generators.LanguageGo)
	assert.Nil(t, mappings)
}

func TestGetFormatMappings_NoMappings(t *testing.T) {
	cfg, err := Load("testdata/golang_only.yaml")
	require.NoError(t, err)

	// Golang config exists but has no format mappings
	mappings := cfg.GetFormatMappings(generators.LanguageGo)
	assert.Nil(t, mappings)

	// Typescript config doesn't exist
	mappings = cfg.GetFormatMappings(generators.LanguageTypeScript)
	assert.Nil(t, mappings)
}

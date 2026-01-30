package deref_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Southclaws/schemancer/schemancer/deref"
	"github.com/Southclaws/schemancer/schemancer/loader"
)

func TestSchema_ExternalRefs(t *testing.T) {
	f, err := os.Open("testdata/user.yaml")
	require.NoError(t, err)
	defer f.Close()

	schema, err := loader.FromReader(f)
	require.NoError(t, err)

	// Before dereferencing, the address property should have a $ref
	require.NotNil(t, schema.Properties["address"])
	assert.Equal(t, "./address.yaml", schema.Properties["address"].Ref)

	// Dereference
	err = deref.Schema(schema, "testdata")
	require.NoError(t, err)

	// After dereferencing, external refs without fragments should be inlined
	address := schema.Properties["address"]
	require.NotNil(t, address)
	assert.Empty(t, address.Ref, "$ref should be cleared after inlining")
	assert.Equal(t, "Address", address.Title)
	assert.Equal(t, "object", address.Type)
	assert.Contains(t, address.Required, "street")
	assert.Contains(t, address.Required, "city")

	// Check that nested ref (country) was also dereferenced and inlined
	country := address.Properties["country"]
	require.NotNil(t, country)
	assert.Empty(t, country.Ref, "nested $ref should be cleared")
	assert.Equal(t, "Country", country.Title)
	assert.Equal(t, "object", country.Type)
}

func TestSchema_NoExternalRefs(t *testing.T) {
	schema, err := loader.FromFile("testdata/country.yaml")
	require.NoError(t, err)

	// Country has no external refs
	err = deref.Schema(schema, "testdata")
	require.NoError(t, err)

	// Schema should remain unchanged
	assert.Equal(t, "Country", schema.Title)
	assert.Equal(t, "object", schema.Type)
	assert.NotNil(t, schema.Properties["code"])
	assert.NotNil(t, schema.Properties["name"])
}

func TestSchema_InternalRefs(t *testing.T) {
	schema, err := loader.FromFile("testdata/with_defs.yaml")
	require.NoError(t, err)

	err = deref.Schema(schema, "testdata")
	require.NoError(t, err)

	// Internal refs should not be modified
	assert.Equal(t, "#/$defs/Name", schema.Properties["name"].Ref)

	// Definitions should be preserved
	require.NotNil(t, schema.Defs)
	assert.NotNil(t, schema.Defs["Name"])
}

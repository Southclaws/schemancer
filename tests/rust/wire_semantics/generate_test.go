package wire_semantics_test

import (
	"strings"
	"testing"

	"github.com/Southclaws/schemancer/schemancer"
	"github.com/Southclaws/schemancer/schemancer/generators"
	rustgen "github.com/Southclaws/schemancer/schemancer/generators/rust"
	"github.com/Southclaws/schemancer/schemancer/ir"
	"github.com/Southclaws/schemancer/schemancer/loader"
	"github.com/Southclaws/schemancer/tests/testutil"
)

func TestRustWireSemantics(t *testing.T) {
	schema, err := loader.FromFile("schema.yaml")
	if err != nil {
		t.Fatalf("load schema: %v", err)
	}

	files, err := schemancer.Generate(schema, generators.GlobalOptions{
		Language: generators.LanguageRust,
		FormatTypeMapping: map[ir.IRFormat]generators.FormatTypeMapping{
			ir.IRFormat("raw"): {
				Type:   "RawJson",
				Import: "crate::RawJson",
			},
		},
	}, rustgen.WithFilename("wire.rs"))
	if err != nil {
		t.Fatalf("generate Rust: %v", err)
	}

	generated := string(testutil.GetSingleFile(t, files))
	for _, expected := range []string{
		"pub required_value: String",
		"pub required_nullable: Nullable<String>",
		"pub optional_value: OptionalField<String>",
		"pub optional_nullable: OptionalField<Nullable<String>>",
		"pub mode: OptionalField<EnvelopeMode>",
		"#[serde(deny_unknown_fields)]",
		"let raw = Box::<RawValue>::deserialize(deserializer)?",
		"pub payload: RawJson",
		"pub enum Event",
	} {
		if !strings.Contains(generated, expected) {
			t.Errorf("generated Rust missing %q", expected)
		}
	}

	if strings.Contains(generated, "#[serde(tag = \"kind\")]") {
		t.Error("derived internally tagged decoding would destroy nested RawValue payloads")
	}
}

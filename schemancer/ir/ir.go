package ir

import "github.com/google/jsonschema-go/jsonschema"

type IR struct {
	Schema *jsonschema.Schema
	Types  []IRType
}

type IRType struct {
	Name        string
	Description string
	Kind        IRTypeKind
	BaseType    string // Name of the base type this extends (from allOf $ref composition)
	Fields      []IRField
	Element     *IRTypeRef
	KeyType     *IRTypeRef
	Enum        []string              // String enum values (backwards compatible)
	EnumValues  []IREnumValue         // Typed enum values (supports int, string, null)
	EnumType    IRBuiltin             // The underlying type of the enum (string, int)
	Union       *IRDiscriminatedUnion // For discriminated unions (oneOf with discriminator)
	SimpleUnion *IRUnion              // For non-discriminated unions (oneOf/anyOf without discriminator)
}

// IREnumValue represents a single enum value with type information
type IREnumValue struct {
	StringValue string // The string representation (for string enums or as name for int enums)
	IntValue    *int   // The integer value (for int enums)
	IsNull      bool   // Whether this is a null value
}

type IRTypeKind string

const (
	IRKindStruct             IRTypeKind = "struct"
	IRKindAlias              IRTypeKind = "alias"
	IRKindEnum               IRTypeKind = "enum"
	IRKindDiscriminatedUnion IRTypeKind = "discriminated_union"
	IRKindUnion              IRTypeKind = "union" // Non-discriminated union (oneOf/anyOf without discriminator)
)

type IRField struct {
	Name        string
	Description string
	JSONName    string
	Type        IRTypeRef
	Required    bool
	Default     *IRDefault        // Default value from JSON Schema "default" keyword
	Extensions  map[string]string // Language-specific extensions (x-java-name, x-go-name, etc.)
}

// IRDefault represents a default value for a field from the JSON Schema "default" keyword.
type IRDefault struct {
	RawValue string    // The raw JSON value as a string (e.g., "1", "\"hello\"", "true")
	Builtin  IRBuiltin // The type of the default value for type-aware rendering
}

type IRTypeRef struct {
	Name        string
	Builtin     IRBuiltin
	Format      IRFormat
	Array       *IRTypeRef
	Map         *IRTypeRef
	Nullable    bool
	Constraints *IRConstraints
}

// IRConstraints holds JSON Schema validation constraints.
// These are captured in the IR so generators can emit validation code
// (e.g., Zod .min()/.max(), Pydantic Field constraints).
type IRConstraints struct {
	// String constraints
	MinLength *int
	MaxLength *int
	Pattern   string

	// Numeric constraints
	Minimum          *float64
	Maximum          *float64
	ExclusiveMinimum *float64
	ExclusiveMaximum *float64
	MultipleOf       *float64

	// Array constraints
	MinItems    *int
	MaxItems    *int
	UniqueItems bool
}

type IRFormat string

const (
	IRFormatNone     IRFormat = ""
	IRFormatByte     IRFormat = "byte"
	IRFormatDateTime IRFormat = "date-time"
	IRFormatDate     IRFormat = "date"
	IRFormatUUID     IRFormat = "uuid"
	IRFormatEmail    IRFormat = "email"
	IRFormatURI      IRFormat = "uri"
)

type IRBuiltin string

const (
	IRBuiltinNone   IRBuiltin = ""
	IRBuiltinString IRBuiltin = "string"
	IRBuiltinInt    IRBuiltin = "int"
	IRBuiltinFloat  IRBuiltin = "float"
	IRBuiltinBool   IRBuiltin = "bool"
	IRBuiltinAny    IRBuiltin = "any"
)

type IRDiscriminatedUnion struct {
	InterfaceName      string
	WrapperName        string
	DiscriminatorField string
	DiscriminatorJSON  string
	Variants           []IRVariant
}

type IRVariant struct {
	Name       string
	ConstValue string
	Type       IRType
}

// IRUnion represents a non-discriminated union (oneOf/anyOf without a discriminator field).
// This is used when the schema defines multiple possible types without a way to distinguish them.
type IRUnion struct {
	Variants []IRTypeRef // The possible types in the union
}

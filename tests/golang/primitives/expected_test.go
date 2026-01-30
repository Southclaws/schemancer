package primitives_test

type AllPrimitives struct {
	BoolField   bool    `json:"boolField"`
	IntField    int     `json:"intField"`
	NumberField float64 `json:"numberField"`
	StringField string  `json:"stringField"`
}

type Amount = float64

type MixedRequired struct {
	OptionalInt    *int    `json:"optionalInt,omitempty"`
	OptionalString *string `json:"optionalString,omitempty"`
	RequiredInt    int     `json:"requiredInt"`
	RequiredString string  `json:"requiredString"`
}

type OptionalPrimitives struct {
	MaybeBool   *bool    `json:"maybeBool,omitempty"`
	MaybeInt    *int     `json:"maybeInt,omitempty"`
	MaybeNumber *float64 `json:"maybeNumber,omitempty"`
	MaybeString *string  `json:"maybeString,omitempty"`
}

type Timestamp = int

type UserId = string

type TypeAliases struct {
	Amount    *Amount    `json:"amount,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	UserID    *UserId    `json:"userId,omitempty"`
}

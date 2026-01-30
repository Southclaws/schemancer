package empty_minimal_test

type ArrayNoItems = []interface{}

type EmptyEnum = string

type EmptyObject = interface{}

type EmptyObjectWithTitle = interface{}

type JustABoolean = bool

type JustANumber = float64

type JustAString = string

type JustAnArray = []interface{}

type JustAnInteger = int

type ObjectNoType struct {
	Field *string `json:"field,omitempty"`
}

type ObjectWithEmptyProps struct {
}

type SingleEnum string

const (
	SingleEnumOnlyValue SingleEnum = "only_value"
)

type TrueSchema = interface{}

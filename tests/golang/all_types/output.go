package all_types

import (
	"github.com/google/uuid"
	"net/mail"
	"net/url"
	"time"
)

type AllPrimitives struct {
	BooleanField *bool       `json:"booleanField,omitempty"`
	IntegerField *int        `json:"integerField,omitempty"`
	NullField    interface{} `json:"nullField,omitempty"`
	NumberField  *float64    `json:"numberField,omitempty"`
	StringField  *string     `json:"stringField,omitempty"`
}

type ArrayVariantsObjectArrayItem struct {
	ID *string `json:"id,omitempty"`
}

type ArrayVariants struct {
	BoundedArray []string                       `json:"boundedArray,omitempty"`
	IntArray     []int                          `json:"intArray,omitempty"`
	MixedArray   []interface{}                  `json:"mixedArray,omitempty"`
	NestedArray  [][]float64                    `json:"nestedArray,omitempty"`
	ObjectArray  []ArrayVariantsObjectArrayItem `json:"objectArray,omitempty"`
	StringArray  []string                       `json:"stringArray,omitempty"`
	UniqueArray  []string                       `json:"uniqueArray,omitempty"`
}

type MapVariants struct {
	AnyMap    map[string]interface{} `json:"anyMap,omitempty"`
	IntMap    map[string]interface{} `json:"intMap,omitempty"`
	ObjectMap map[string]interface{} `json:"objectMap,omitempty"`
	StringMap map[string]interface{} `json:"stringMap,omitempty"`
}

type NumberConstraints struct {
	ExclusiveRange *float64 `json:"exclusiveRange,omitempty"`
	MultipleOf     *int     `json:"multipleOf,omitempty"`
	NegativeInt    *int     `json:"negativeInt,omitempty"`
	PositiveInt    *int     `json:"positiveInt,omitempty"`
	RangeInt       *int     `json:"rangeInt,omitempty"`
}

type StringConstraints struct {
	FixedLength   *string `json:"fixedLength,omitempty"`
	LongString    *string `json:"longString,omitempty"`
	PatternString *string `json:"patternString,omitempty"`
	ShortString   *string `json:"shortString,omitempty"`
}

type StringFormats struct {
	Date                *time.Time    `json:"date,omitempty"`
	DateTime            *time.Time    `json:"dateTime,omitempty"`
	Duration            *string       `json:"duration,omitempty"`
	Email               *mail.Address `json:"email,omitempty"`
	Hostname            *string       `json:"hostname,omitempty"`
	IdnEmail            *string       `json:"idnEmail,omitempty"`
	IdnHostname         *string       `json:"idnHostname,omitempty"`
	Ipv4                *string       `json:"ipv4,omitempty"`
	Ipv6                *string       `json:"ipv6,omitempty"`
	Iri                 *string       `json:"iri,omitempty"`
	IriReference        *string       `json:"iriReference,omitempty"`
	JSONPointer         *string       `json:"jsonPointer,omitempty"`
	Regex               *string       `json:"regex,omitempty"`
	RelativeJSONPointer *string       `json:"relativeJsonPointer,omitempty"`
	Time                *string       `json:"time,omitempty"`
	URI                 *url.URL      `json:"uri,omitempty"`
	URIReference        *string       `json:"uriReference,omitempty"`
	URITemplate         *string       `json:"uriTemplate,omitempty"`
	UUIDFormat          *uuid.UUID    `json:"uuidFormat,omitempty"`
}

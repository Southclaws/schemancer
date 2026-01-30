package special_chars_test

type NumberPrefixed struct {
	N1field   *string `json:"1field,omitempty"`
	N2ndField *string `json:"2nd_field,omitempty"`
	N3rdField *int    `json:"3rdField,omitempty"`
}

type WeirdNames struct {
	HashTag                *string `json:"#hashTag,omitempty"`
	DoubleDollar           *string `json:"$$doubleDollar,omitempty"`
	DollarSign             *string `json:"$dollarSign,omitempty"`
	N123startsWithNumber   *string `json:"123startsWithNumber,omitempty"`
	AtSign                 *string `json:"@atSign,omitempty"`
	ABC                    *string `json:"ABC,omitempty"`
	ALLCAPS                *string `json:"ALLCAPS,omitempty"`
	MixedWith123Numbers456 *string `json:"MixedWith123Numbers456,omitempty"`
	PascalCase             *string `json:"PascalCase,omitempty"`
	XMLHttpRequest         *string `json:"XMLHttpRequest,omitempty"`
	A                      *string `json:"a,omitempty"`
	Ab                     *string `json:"ab,omitempty"`
	CamelCase              *string `json:"camelCase,omitempty"`
	GetHttpresponse        *string `json:"getHTTPResponse,omitempty"`
	KebabCaseName          *string `json:"kebab-case-name,omitempty"`
	ParseJSON              *string `json:"parseJSON,omitempty"`
	SnakeCaseName          *string `json:"snake_case_name,omitempty"`
	WithSpaces             *string `json:"with spaces,omitempty"`
	WithDashes             *string `json:"with-dashes,omitempty"`
	WithUnderscores        *string `json:"with_underscores,omitempty"`
}

package reserved_keywords_test

type GoKeywords struct {
	Break       *string                `json:"break,omitempty"`
	Case        *string                `json:"case,omitempty"`
	Chan        *string                `json:"chan,omitempty"`
	Const       *string                `json:"const,omitempty"`
	Continue    *string                `json:"continue,omitempty"`
	Default     *string                `json:"default,omitempty"`
	Defer       *string                `json:"defer,omitempty"`
	Else        *string                `json:"else,omitempty"`
	Fallthrough *string                `json:"fallthrough,omitempty"`
	For         *string                `json:"for,omitempty"`
	Func        string                 `json:"func"`
	Go          *string                `json:"go,omitempty"`
	If          *string                `json:"if,omitempty"`
	Import      *string                `json:"import,omitempty"`
	Interface   *string                `json:"interface,omitempty"`
	Map         map[string]interface{} `json:"map"`
	Package     *string                `json:"package,omitempty"`
	Range       *string                `json:"range,omitempty"`
	Return      *string                `json:"return,omitempty"`
	Select      *string                `json:"select,omitempty"`
	Struct      *string                `json:"struct,omitempty"`
	Switch      *string                `json:"switch,omitempty"`
	Type        string                 `json:"type"`
	Var         *string                `json:"var,omitempty"`
}

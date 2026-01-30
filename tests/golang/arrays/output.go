package arrays

type IntArray = []int

type NestedArray = [][]float64

type ObjectArrayItem struct {
	ID    string `json:"id"`
	Value *int   `json:"value,omitempty"`
}

type ObjectArray = []ObjectArrayItem

type StringArray = []string

type MixedContainer struct {
	Nested  *NestedArray `json:"nested,omitempty"`
	Numbers *IntArray    `json:"numbers,omitempty"`
	Objects *ObjectArray `json:"objects,omitempty"`
	Strings *StringArray `json:"strings,omitempty"`
}

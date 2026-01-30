package deep_nesting

type DeepArrayItemItemItem struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type DeepArray = [][][]DeepArrayItemItemItem

type Level1Level2Level3Level4Level5DataItem struct {
	Key   *string `json:"key,omitempty"`
	Value *int    `json:"value,omitempty"`
}

type Level1Level2Level3Level4Level5 struct {
	Data []Level1Level2Level3Level4Level5DataItem `json:"data,omitempty"`
	Name string                                   `json:"name"`
}

type Level1Level2Level3Level4 struct {
	Level5 *Level1Level2Level3Level4Level5 `json:"level5,omitempty"`
	Name   string                          `json:"name"`
}

type Level1Level2Level3 struct {
	Level4 *Level1Level2Level3Level4 `json:"level4,omitempty"`
	Name   string                    `json:"name"`
}

type Level1Level2 struct {
	Level3 *Level1Level2Level3 `json:"level3,omitempty"`
	Name   string              `json:"name"`
}

type Level1 struct {
	Level2 *Level1Level2 `json:"level2,omitempty"`
	Name   string        `json:"name"`
}

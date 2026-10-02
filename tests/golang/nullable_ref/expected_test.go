package nullable_ref_test

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Container struct {
	Address *Address `json:"address"`
	Name    string   `json:"name"`
}

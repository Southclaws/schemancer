package refs_test

type Address struct {
	City    string  `json:"city"`
	Country *string `json:"country,omitempty"`
	Street  string  `json:"street"`
}

type Person struct {
	Friends     []Person `json:"friends,omitempty"`
	HomeAddress *Address `json:"homeAddress,omitempty"`
	Name        string   `json:"name"`
	WorkAddress *Address `json:"workAddress,omitempty"`
}

type Company struct {
	Ceo          *Person  `json:"ceo,omitempty"`
	Employees    []Person `json:"employees,omitempty"`
	Headquarters *Address `json:"headquarters,omitempty"`
	Name         string   `json:"name"`
}

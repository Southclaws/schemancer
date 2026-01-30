package nested_objects_test

type Customer struct {
	Email *string `json:"email,omitempty"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
}

type LineItem struct {
	Price     float64 `json:"price"`
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	Customer Customer   `json:"customer"`
	ID       string     `json:"id"`
	Items    []LineItem `json:"items"`
	Total    *float64   `json:"total,omitempty"`
}

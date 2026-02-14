package multi

import (
	"github.com/google/uuid"
)

type Address struct {
	City    string  `json:"city"`
	Country *string `json:"country,omitempty"`
	Street  string  `json:"street"`
}

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
)

var StatusValues = []Status{
	StatusActive,
	StatusInactive,
	StatusPending,
}

type Person struct {
	Address *Address  `json:"address,omitempty"`
	Age     *int      `json:"age,omitempty"`
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Status  Status    `json:"status"`
}

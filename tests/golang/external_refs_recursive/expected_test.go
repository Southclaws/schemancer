package external_refs_recursive_test

import (
	"github.com/google/uuid"
	"net/mail"
)

type Address struct {
	City    string  `json:"city"`
	Country *string `json:"country,omitempty"`
	Street  string  `json:"street"`
}

type ContactInfo struct {
	Address *Address     `json:"address,omitempty"`
	Email   mail.Address `json:"email"`
	Phone   *string      `json:"phone,omitempty"`
}

type Author struct {
	ContactInfo *ContactInfo `json:"contactInfo,omitempty"`
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
}

type Article struct {
	Author    Author    `json:"author"`
	CoAuthors []Author  `json:"coAuthors,omitempty"`
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
}

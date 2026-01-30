package simple_object_test

import (
	"github.com/google/uuid"
	"net/mail"
)

type User struct {
	Active *bool        `json:"active,omitempty"`
	Age    *int         `json:"age,omitempty"`
	Email  mail.Address `json:"email"`
	ID     uuid.UUID    `json:"id"`
	Name   string       `json:"name"`
}

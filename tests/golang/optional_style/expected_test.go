package optional_style_test

import (
	"github.com/Southclaws/opt"
	"github.com/google/uuid"
	"time"
)

type PersonAddress struct {
	City   string               `json:"city"`
	Street string               `json:"street"`
	Zip    opt.Optional[string] `json:"zip,omitempty"`
}

type PersonFriendsItem struct {
	Name  string                  `json:"name"`
	Since opt.Optional[time.Time] `json:"since,omitempty"`
}

type Person struct {
	Active       opt.Optional[bool]          `json:"active,omitempty"`
	Address      opt.Optional[PersonAddress] `json:"address,omitempty"`
	Age          opt.Optional[int]           `json:"age,omitempty"`
	Birthday     opt.Optional[time.Time]     `json:"birthday,omitempty"`
	Email        opt.Optional[string]        `json:"email,omitempty"`
	Friends      []PersonFriendsItem         `json:"friends,omitempty"`
	Metadata     map[string]interface{}      `json:"metadata,omitempty"`
	Name         string                      `json:"name"`
	ReferralCode opt.Optional[uuid.UUID]     `json:"referralCode,omitempty"`
	Roles        []string                    `json:"roles"`
	Tags         []string                    `json:"tags,omitempty"`
}

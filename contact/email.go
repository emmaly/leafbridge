package contact

import (
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Email is an email type of Contact
type Email struct {
	leafbridge.Common
	ID           id.Contact
	Type         Type    `firestore:",omitempty"`
	Context      Context `firestore:",omitempty"`
	EmailAddress string  `firestore:",omitempty"`
}

// AsEmail converts a Contact to an Email
func (c Contact) AsEmail() Email {
	if c.Type != TypeEmail {
		return Email{}
	}
	x := Email{
		ID:           c.ID,
		Type:         c.Type,
		Context:      c.Context,
		EmailAddress: c.EmailAddress,
		Common: leafbridge.Common{
			Notes: c.Notes,
		},
	}
	return x
}

// AsContact converts an Email into a Contact
func (x Email) AsContact() Contact {
	c := Contact{
		ID:           x.ID,
		Type:         x.Type,
		Context:      x.Context,
		EmailAddress: x.EmailAddress,
		Common: leafbridge.Common{
			Notes: x.Notes,
		},
	}
	return c
}

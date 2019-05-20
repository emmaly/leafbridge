package contact

import (
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Email is an email type of Contact
type Email struct {
	ID           id.Contact
	Type         Type        `firestore:",omitempty"`
	Context      Context     `firestore:",omitempty"`
	Notes        []note.Note `firestore:",omitempty"`
	EmailAddress string      `firestore:",omitempty"`
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
		Notes:        c.Notes,
		EmailAddress: c.EmailAddress,
	}
	return x
}

// AsContact converts an Email into a Contact
func (x Email) AsContact() Contact {
	c := Contact{
		ID:           x.ID,
		Type:         x.Type,
		Context:      x.Context,
		Notes:        x.Notes,
		EmailAddress: x.EmailAddress,
	}
	return c
}

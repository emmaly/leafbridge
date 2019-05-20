package person

import (
	"github.com/emmaly/leafbridge/id"
)

// EmailContact is an email type of Contact
type EmailContact struct {
	ID           id.Contact
	Type         ContactType    `firestore:",omitempty"`
	Context      ContactContext `firestore:",omitempty"`
	Notes        []id.Note      `firestore:",omitempty"`
	EmailAddress string         `firestore:",omitempty"`
}

// EmailContact converts a Contact to an EmailContact
func (c Contact) EmailContact() EmailContact {
	if c.Type != Email {
		return EmailContact{}
	}
	ec := EmailContact{
		ID:           c.ID,
		Type:         c.Type,
		Context:      c.Context,
		Notes:        c.Notes,
		EmailAddress: c.EmailAddress,
	}
	return ec
}

// Contact converts an EmailContact into a Contact
func (ec EmailContact) Contact() Contact {
	c := Contact{
		ID:           ec.ID,
		Type:         ec.Type,
		Context:      ec.Context,
		Notes:        ec.Notes,
		EmailAddress: ec.EmailAddress,
	}
	return c
}

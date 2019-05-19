package person

import "github.com/emmaly/leafbridge/note"

// EmailContact is an email type of Contact
type EmailContact struct {
	Type         ContactType
	Context      ContactContext
	Notes        []note.Note
	EmailAddress string
}

// EmailContact converts a Contact to an EmailContact
func (c Contact) EmailContact() EmailContact {
	if c.Type != Email {
		return EmailContact{}
	}
	ec := EmailContact{
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
		Type:         ec.Type,
		Context:      ec.Context,
		Notes:        ec.Notes,
		EmailAddress: ec.EmailAddress,
	}
	return c
}

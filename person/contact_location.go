package person

import (
	"github.com/emmaly/leafbridge/id"
)

// LocationContact is an email type of Contact
type LocationContact struct {
	ID          id.Contact
	Type        ContactType    `firestore:",omitempty"`
	Context     ContactContext `firestore:",omitempty"`
	Notes       []id.Note      `firestore:",omitempty"`
	Address     []string       `firestore:",omitempty"`
	City        string         `firestore:",omitempty"`
	Region      string         `firestore:",omitempty"`
	PostalCode  string         `firestore:",omitempty"`
	CountryCode string         `firestore:",omitempty"`
}

// LocationContact converts a Contact to an LocationContact
func (c Contact) LocationContact() LocationContact {
	if c.Type != PostalMail && c.Type != PhysicalLocation {
		return LocationContact{}
	}
	ec := LocationContact{
		ID:          c.ID,
		Type:        c.Type,
		Context:     c.Context,
		Notes:       c.Notes,
		Address:     c.Address,
		City:        c.City,
		Region:      c.Region,
		PostalCode:  c.PostalCode,
		CountryCode: c.CountryCode,
	}
	return ec
}

// Contact converts an LocationContact into a Contact
func (ec LocationContact) Contact() Contact {
	c := Contact{
		ID:          ec.ID,
		Type:        ec.Type,
		Context:     ec.Context,
		Notes:       ec.Notes,
		Address:     ec.Address,
		City:        ec.City,
		Region:      ec.Region,
		PostalCode:  ec.PostalCode,
		CountryCode: ec.CountryCode,
	}
	return c
}

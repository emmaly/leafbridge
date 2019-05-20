package contact

import (
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Location is an email type of Contact
type Location struct {
	ID          id.Contact
	Type        Type        `firestore:",omitempty"`
	Context     Context     `firestore:",omitempty"`
	Notes       []note.Note `firestore:",omitempty"`
	Address     []string    `firestore:",omitempty"`
	City        string      `firestore:",omitempty"`
	Region      string      `firestore:",omitempty"`
	PostalCode  string      `firestore:",omitempty"`
	CountryCode string      `firestore:",omitempty"`
}

// AsLocation converts a Contact to an Location
func (c Contact) AsLocation() Location {
	if c.Type != TypePostalMail && c.Type != TypePhysicalLocation {
		return Location{}
	}
	x := Location{
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
	return x
}

// AsContact converts an Location into a Contact
func (x Location) AsContact() Contact {
	c := Contact{
		ID:          x.ID,
		Type:        x.Type,
		Context:     x.Context,
		Notes:       x.Notes,
		Address:     x.Address,
		City:        x.City,
		Region:      x.Region,
		PostalCode:  x.PostalCode,
		CountryCode: x.CountryCode,
	}
	return c
}

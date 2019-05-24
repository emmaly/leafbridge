package contact

import (
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Location is an email type of Contact
type Location struct {
	leafbridge.Common
	ID          id.Contact
	Type        Type     `firestore:",omitempty"`
	Context     Context  `firestore:",omitempty"`
	Address     []string `firestore:",omitempty"`
	City        string   `firestore:",omitempty"`
	Region      string   `firestore:",omitempty"`
	PostalCode  string   `firestore:",omitempty"`
	CountryCode string   `firestore:",omitempty"`
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
		Address:     c.Address,
		City:        c.City,
		Region:      c.Region,
		PostalCode:  c.PostalCode,
		CountryCode: c.CountryCode,
		Common: leafbridge.Common{
			Notes: c.Notes,
		},
	}
	return x
}

// AsContact converts an Location into a Contact
func (x Location) AsContact() Contact {
	c := Contact{
		ID:          x.ID,
		Type:        x.Type,
		Context:     x.Context,
		Address:     x.Address,
		City:        x.City,
		Region:      x.Region,
		PostalCode:  x.PostalCode,
		CountryCode: x.CountryCode,
		Common: leafbridge.Common{
			Notes: x.Notes,
		},
	}
	return c
}

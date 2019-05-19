package person

import "github.com/emmaly/leafbridge/note"

// LocationContact is an email type of Contact
type LocationContact struct {
	Type        ContactType
	Context     ContactContext
	Notes       []note.Note
	Address     []string
	City        string
	Region      string
	PostalCode  string
	CountryCode string
}

// LocationContact converts a Contact to an LocationContact
func (c Contact) LocationContact() LocationContact {
	if c.Type != PostalMail && c.Type != PhysicalLocation {
		return LocationContact{}
	}
	ec := LocationContact{
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

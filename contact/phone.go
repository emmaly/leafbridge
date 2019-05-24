package contact

import (
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Phone is an email type of Contact
type Phone struct {
	leafbridge.Common
	ID      id.Contact
	Type    Type    `firestore:",omitempty"`
	Context Context `firestore:",omitempty"`
	Number  string  `firestore:",omitempty"`
}

// AsPhone converts a Contact to an Phone
func (c Contact) AsPhone() Phone {
	if c.Type != TypeMobilePhone && c.Type != TypeVoicePhone && c.Type != TypeVoiceMessage && c.Type != TypeFax && c.Type != TypePager {
		return Phone{}
	}
	x := Phone{
		ID:      c.ID,
		Type:    c.Type,
		Context: c.Context,
		Number:  c.Number,
		Common: leafbridge.Common{
			Notes: c.Notes,
		},
	}
	return x
}

// AsContact converts an Phone into a Contact
func (x Phone) AsContact() Contact {
	c := Contact{
		ID:      x.ID,
		Type:    x.Type,
		Context: x.Context,
		Number:  x.Number,
		Common: leafbridge.Common{
			Notes: x.Notes,
		},
	}
	return c
}

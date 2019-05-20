package contact

import (
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Phone is an email type of Contact
type Phone struct {
	ID      id.Contact
	Type    Type        `firestore:",omitempty"`
	Context Context     `firestore:",omitempty"`
	Notes   []note.Note `firestore:",omitempty"`
	Number  string      `firestore:",omitempty"`
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
		Notes:   c.Notes,
		Number:  c.Number,
	}
	return x
}

// AsContact converts an Phone into a Contact
func (x Phone) AsContact() Contact {
	c := Contact{
		ID:      x.ID,
		Type:    x.Type,
		Context: x.Context,
		Notes:   x.Notes,
		Number:  x.Number,
	}
	return c
}

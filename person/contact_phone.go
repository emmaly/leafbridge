package person

import (
	"github.com/emmaly/leafbridge/id"
)

// PhoneContact is an email type of Contact
type PhoneContact struct {
	ID      id.Contact
	Type    ContactType    `firestore:",omitempty"`
	Context ContactContext `firestore:",omitempty"`
	Notes   []id.Note      `firestore:",omitempty"`
	Number  string         `firestore:",omitempty"`
}

// PhoneContact converts a Contact to an PhoneContact
func (c Contact) PhoneContact() PhoneContact {
	if c.Type != MobilePhone && c.Type != VoicePhone && c.Type != VoiceMessage && c.Type != Fax && c.Type != Pager {
		return PhoneContact{}
	}
	ec := PhoneContact{
		ID:      c.ID,
		Type:    c.Type,
		Context: c.Context,
		Notes:   c.Notes,
		Number:  c.Number,
	}
	return ec
}

// Contact converts an PhoneContact into a Contact
func (ec PhoneContact) Contact() Contact {
	c := Contact{
		ID:      ec.ID,
		Type:    ec.Type,
		Context: ec.Context,
		Notes:   ec.Notes,
		Number:  ec.Number,
	}
	return c
}

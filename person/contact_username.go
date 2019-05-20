package person

import (
	"github.com/emmaly/leafbridge/id"
)

// UsernameContact is a username type of Contact
type UsernameContact struct {
	ID       id.Contact
	Type     ContactType    `firestore:",omitempty"`
	Context  ContactContext `firestore:",omitempty"`
	Notes    []id.Note      `firestore:",omitempty"`
	Username string         `firestore:",omitempty"`
}

// UsernameContact converts a Contact to an UsernameContact
func (c Contact) UsernameContact() UsernameContact {
	if c.Type != GitHub && c.Type != Twitter && c.Type != Tumblr && c.Type != Diaspora && c.Type != Mastodon && c.Type != Facebook {
		return UsernameContact{}
	}
	dc := UsernameContact{
		ID:       c.ID,
		Type:     c.Type,
		Context:  c.Context,
		Notes:    c.Notes,
		Username: c.Username,
	}
	return dc
}

// Contact converts an UsernameContact into a Contact
func (dc UsernameContact) Contact() Contact {
	c := Contact{
		ID:       dc.ID,
		Type:     dc.Type,
		Context:  dc.Context,
		Notes:    dc.Notes,
		Username: dc.Username,
	}
	return c
}

// URL returns the URL to the account, if there is one available
func (dc UsernameContact) URL() string {
	switch dc.Type {
	case GitHub:
		return "https://github.com/" + dc.Username
	case Twitter:
		return "https://twitter.com/" + dc.Username
	case Tumblr:
		return "https://" + dc.Username + "tumblr.com/"
	case Facebook:
		return "https://www.facebook.com/" + dc.Username
	default:
		return ""
	}
}

package contact

import (
	"time"

	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Username is a username type of Contact
type Username struct {
	ID         id.Contact
	Type       Type        `firestore:",omitempty"`
	Context    Context     `firestore:",omitempty"`
	Created    time.Time   `firestore:",omitempty"`
	CreatedBy  id.Person   `firestore:",omitempty"`
	Modified   time.Time   `firestore:",omitempty"`
	ModifiedBy id.Person   `firestore:",omitempty"`
	Notes      []note.Note `firestore:",omitempty"`
	Username   string      `firestore:",omitempty"`
}

// AsUsername converts a Contact to an Username
func (c Contact) AsUsername() Username {
	if c.Type != TypeGitHub && c.Type != TypeTwitter && c.Type != TypeTumblr && c.Type != TypeDiaspora && c.Type != TypeMastodon && c.Type != TypeFacebook {
		return Username{}
	}
	x := Username{
		ID:       c.ID,
		Type:     c.Type,
		Context:  c.Context,
		Notes:    c.Notes,
		Username: c.Username,
	}
	return x
}

// AsContact converts an Username into a Contact
func (x Username) AsContact() Contact {
	c := Contact{
		ID:       x.ID,
		Type:     x.Type,
		Context:  x.Context,
		Notes:    x.Notes,
		Username: x.Username,
	}
	return c
}

// URL returns the URL to the account, if there is one available
func (x Username) URL() string {
	switch x.Type {
	case TypeGitHub:
		return "https://github.com/" + x.Username
	case TypeTwitter:
		return "https://twitter.com/" + x.Username
	case TypeTumblr:
		return "https://" + x.Username + "tumblr.com/"
	case TypeFacebook:
		return "https://www.facebook.com/" + x.Username
	default:
		return ""
	}
}

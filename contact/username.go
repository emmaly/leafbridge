package contact

import (
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Username is a username type of Contact
type Username struct {
	leafbridge.Common
	ID       id.Contact
	Type     Type    `firestore:",omitempty"`
	Context  Context `firestore:",omitempty"`
	Username string  `firestore:",omitempty"`
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
		Username: c.Username,
		Common: leafbridge.Common{
			Notes: c.Notes,
		},
	}
	return x
}

// AsContact converts an Username into a Contact
func (x Username) AsContact() Contact {
	c := Contact{
		ID:       x.ID,
		Type:     x.Type,
		Context:  x.Context,
		Username: x.Username,
		Common: leafbridge.Common{
			Notes: x.Notes,
		},
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

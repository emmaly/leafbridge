package contact

import (
	"time"

	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Contact provides ...
type Contact interface {
	Type() Type
}

// Common is the common set of fields for Contacts
type Common struct {
	leafbridge.Common
	ID      id.Contact `firestore:",omitempty"`
	Type    Type       `firestore:",omitempty"`
	Context Context    `firestore:",omitempty"`
}

// New initializes the common fields
func (c *Common) New() {
	c.ID = id.NewContact()
	c.Created = time.Now()
	c.Modified = c.Created
}

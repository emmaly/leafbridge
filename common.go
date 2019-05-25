package leafbridge

import (
	"time"

	"github.com/emmaly/leafbridge/id"
)

// Common is the set of fields that are used throughout LeafBridge types
type Common struct {
	Created    time.Time `firestore:",omitempty"`
	CreatedBy  id.Person `firestore:",omitempty"`
	Modified   time.Time `firestore:",omitempty"`
	ModifiedBy id.Person `firestore:",omitempty"`
	Notes      []Note    `firestore:",omitempty"`
}

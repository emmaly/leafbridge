package computer

import (
	"time"

	"github.com/emmaly/leafbridge/id"
)

// Computer is a computer
type Computer struct {
	ID          id.Computer
	Name        string    `firestore:",omitempty"`
	PrimaryUser id.Person `firestore:",omitempty"`
	Created     time.Time `firestore:",omitempty"`
	LastSeen    time.Time `firestore:",omitempty"`
	Notes       []id.Note `firestore:",omitempty"`
}

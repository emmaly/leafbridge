package computer

import (
	"time"

	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Computer is a computer
type Computer struct {
	ID          id.Computer
	Name        string      `firestore:",omitempty"`
	PrimaryUser id.Person   `firestore:",omitempty"`
	Created     time.Time   `firestore:",omitempty"`
	LastSeen    time.Time   `firestore:",omitempty"`
	Notes       []note.Note `firestore:",omitempty"`
}

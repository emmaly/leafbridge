package computer

import (
	"time"

	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Computer is a computer
type Computer struct {
	leafbridge.Common
	ID          id.Computer
	Name        string    `firestore:",omitempty"`
	PrimaryUser id.Person `firestore:",omitempty"`
	LastSeen    time.Time `firestore:",omitempty"`
}

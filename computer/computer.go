package computer

import (
	"time"

	"github.com/emmaly/leafbridge/id"
)

// Computer is a computer
type Computer struct {
	ID          id.Computer
	Name        string
	PrimaryUser id.Person
	Created     time.Time
	LastSeen    time.Time
	Notes       []id.Note
}

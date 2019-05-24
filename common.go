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

// Note is a note
type Note struct {
	Common
	ID    id.Note
	Type  NoteType `firestore:",omitempty"`
	Title string   `firestore:",omitempty"`
	Body  string   `firestore:",omitempty"`
}

// NoteType indicates what type of note this Note is.
type NoteType uint8

// NoteType constants
const (
	General NoteType = iota
	Custom
	Purchase
	Install
	Disposition
	Return
	Repair
	Transfer
)

var tNames = map[NoteType]string{
	General:     "General",
	Custom:      "Custom",
	Purchase:    "Purchase",
	Install:     "Install",
	Disposition: "Disposition",
	Return:      "Return",
	Repair:      "Repair",
	Transfer:    "Transfer",
}

func (t NoteType) String() string {
	return tNames[t]
}

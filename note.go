package leafbridge

import (
	"time"

	"github.com/emmaly/leafbridge/id"
)

// Note is a note
type Note struct {
	Common
	ID    id.Note
	Type  NoteType `firestore:",omitempty"`
	Title string   `firestore:",omitempty"`
	Body  string   `firestore:",omitempty"`
}

// NewNote returns a new and initialized Note
func NewNote() Note {
	var n Note
	n.ID = id.NewNote()
	n.Created = time.Now()
	n.Modified = n.Created
	return n
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

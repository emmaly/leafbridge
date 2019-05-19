package person

import (
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Person is a person
type Person struct {
	ID      id.Person
	Name    Name
	Contact []Contact
	Notes   []note.Note
}

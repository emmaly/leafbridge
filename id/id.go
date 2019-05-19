package id

import "github.com/segmentio/ksuid"

//
//
// New returns a new ID string
func New() string {
	return ksuid.New().String()
}

//
//
// Note is an ID for a note
type Note string

// NewNote returns a new Note
func NewNote() Note {
	return Note(New())
}

//
//
// Person is an ID for a person
type Person string

// NewPerson returns a new Person
func NewPerson() Person {
	return Person(New())
}

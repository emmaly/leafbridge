package person

import "github.com/emmaly/leafbridge/id"

// Person is a person
type Person struct {
	ID            id.Person
	Name          Name
	Contacts      []id.Contact
	Notes         []id.Note
	Supervisor    id.Person
	DirectReports []id.Person
}

// LoadPerson fetches a Person from the DB by ID
func LoadPerson(id id.Person) Person {
	return Person{
		ID: id,
	}
}

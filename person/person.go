package person

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Person is a person
type Person struct {
	ID            id.Person
	Name          Name
	Contacts      []id.Contact `firestore:",omitempty"`
	Notes         []id.Note    `firestore:",omitempty"`
	Supervisor    id.Person    `firestore:",omitempty"`
	DirectReports []id.Person  `firestore:",omitempty"`
}

// LoadPerson fetches a Person from the DB by ID
func LoadPerson(ctx context.Context, fs *firestore.Client, id id.Person) (Person, error) {
	doc, err := fs.Collection("people").Doc(string(id)).Get(ctx)
	if err != nil {
		return Person{}, err
	}
	if !doc.Exists() {
		return Person{}, leafbridge.ErrMissingRecord
	}

	var p Person
	err = doc.DataTo(&p)
	if err != nil {
		return Person{}, err
	}

	return p, nil
}

// DeletePerson delete a Person from the DB by ID
func DeletePerson(ctx context.Context, fs *firestore.Client, id id.Person) error {
	_, err := fs.Collection("people").Doc(string(id)).Delete(ctx)
	return err
}

// Save stores a Person to the DB
func (c Person) Save(ctx context.Context, fs *firestore.Client) error {
	_, err := fs.Collection("people").Doc(string(c.ID)).Set(ctx, c)
	return err
}

package person

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/contact"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Person is a person
type Person struct {
	ID              id.Person
	Name            Name
	Created         time.Time         `firestore:",omitempty"`
	CreatedBy       id.Person         `firestore:",omitempty"`
	Modified        time.Time         `firestore:",omitempty"`
	ModifiedBy      id.Person         `firestore:",omitempty"`
	Notes           []note.Note       `firestore:",omitempty"`
	Title           string            `firestore:",omitempty"`
	Contacts        []contact.Contact `firestore:",omitempty"`
	Supervisor      id.Person         `firestore:",omitempty"`
	DirectReports   []id.Person       `firestore:",omitempty"`
	PrimaryLocation id.Location       `firestore:",omitempty"`
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

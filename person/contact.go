package person

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Contact is a contact method for this Contact
type Contact struct {
	ID           id.Contact
	Type         ContactType    `firestore:",omitempty"`
	Context      ContactContext `firestore:",omitempty"`
	Notes        []id.Note      `firestore:",omitempty"`
	EmailAddress string         `firestore:",omitempty"`
	Number       string         `firestore:",omitempty"`
	Username     string         `firestore:",omitempty"`
	URL          string         `firestore:",omitempty"`
	Description  string         `firestore:",omitempty"`
	Value        string         `firestore:",omitempty"`
	Address      []string       `firestore:",omitempty"`
	City         string         `firestore:",omitempty"`
	Region       string         `firestore:",omitempty"`
	PostalCode   string         `firestore:",omitempty"`
	CountryCode  string         `firestore:",omitempty"`
}

// Contact converts a Contact into a Contact (it just returns itself)
func (c Contact) Contact() Contact {
	return c
}

// LoadContact fetches a Contact from the DB by ID
func LoadContact(ctx context.Context, fs *firestore.Client, id id.Contact) (Contact, error) {
	doc, err := fs.Collection("contacts").Doc(string(id)).Get(ctx)
	if err != nil {
		return Contact{}, err
	}
	if !doc.Exists() {
		return Contact{}, leafbridge.ErrMissingRecord
	}

	var p Contact
	err = doc.DataTo(&p)
	if err != nil {
		return Contact{}, err
	}

	return p, nil
}

// DeleteContact delete a Contact from the DB by ID
func DeleteContact(ctx context.Context, fs *firestore.Client, id id.Contact) error {
	_, err := fs.Collection("contacts").Doc(string(id)).Delete(ctx)
	return err
}

// Save stores a Contact to the DB
func (c Contact) Save(ctx context.Context, fs *firestore.Client) error {
	_, err := fs.Collection("contacts").Doc(string(c.ID)).Set(ctx, c)
	return err
}

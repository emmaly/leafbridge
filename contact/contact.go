package contact

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Contact is a contact method for this Contact
type Contact struct {
	ID           id.Contact
	Type         Type        `firestore:",omitempty"`
	Context      Context     `firestore:",omitempty"`
	Created      time.Time   `firestore:",omitempty"`
	CreatedBy    id.Person   `firestore:",omitempty"`
	Modified     time.Time   `firestore:",omitempty"`
	ModifiedBy   id.Person   `firestore:",omitempty"`
	Notes        []note.Note `firestore:",omitempty"`
	EmailAddress string      `firestore:",omitempty"`
	Number       string      `firestore:",omitempty"`
	Username     string      `firestore:",omitempty"`
	URL          string      `firestore:",omitempty"`
	Description  string      `firestore:",omitempty"`
	Value        string      `firestore:",omitempty"`
	Address      []string    `firestore:",omitempty"`
	City         string      `firestore:",omitempty"`
	Region       string      `firestore:",omitempty"`
	PostalCode   string      `firestore:",omitempty"`
	CountryCode  string      `firestore:",omitempty"`
}

// AsContact converts a Contact into a Contact (it just returns itself)
func (c Contact) AsContact() Contact {
	return c
}

// Load fetches a Contact from the DB by ID
func Load(ctx context.Context, fs *firestore.Client, id id.Contact) (Contact, error) {
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

// Delete delete a Contact from the DB by ID
func Delete(ctx context.Context, fs *firestore.Client, id id.Contact) error {
	_, err := fs.Collection("contacts").Doc(string(id)).Delete(ctx)
	return err
}

// Save stores a Contact to the DB
func (c Contact) Save(ctx context.Context, fs *firestore.Client) error {
	_, err := fs.Collection("contacts").Doc(string(c.ID)).Set(ctx, c)
	return err
}

package note

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Note is a note
type Note struct {
	ID         id.Note
	Type       Type      `firestore:",omitempty"`
	Title      string    `firestore:",omitempty"`
	Body       string    `firestore:",omitempty"`
	Created    time.Time `firestore:",omitempty"`
	CreatedBy  id.Person `firestore:",omitempty"`
	Modified   time.Time `firestore:",omitempty"`
	ModifiedBy id.Person `firestore:",omitempty"`
	Notes      []Note    `firestore:",omitempty"`
}

// LoadNote fetches a Note from the DB by ID
func LoadNote(ctx context.Context, fs *firestore.Client, id id.Note) (Note, error) {
	doc, err := fs.Collection("notes").Doc(string(id)).Get(ctx)
	if err != nil {
		return Note{}, err
	}
	if !doc.Exists() {
		return Note{}, leafbridge.ErrMissingRecord
	}

	var p Note
	err = doc.DataTo(&p)
	if err != nil {
		return Note{}, err
	}

	return p, nil
}

// DeleteNote delete a Note from the DB by ID
func DeleteNote(ctx context.Context, fs *firestore.Client, id id.Note) error {
	_, err := fs.Collection("notes").Doc(string(id)).Delete(ctx)
	return err
}

// Save stores a Note to the DB
func (c Note) Save(ctx context.Context, fs *firestore.Client) error {
	_, err := fs.Collection("notes").Doc(string(c.ID)).Set(ctx, c)
	return err
}

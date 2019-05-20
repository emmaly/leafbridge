package location

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
)

// Location is a note
type Location struct {
	ID          id.Location
	Type        Type        `firestore:",omitempty"`
	Name        string      `firestore:",omitempty"`
	Description string      `firestore:",omitempty"`
	Created     time.Time   `firestore:",omitempty"`
	CreatedBy   id.Person   `firestore:",omitempty"`
	Modified    time.Time   `firestore:",omitempty"`
	ModifiedBy  id.Person   `firestore:",omitempty"`
	Notes       []note.Note `firestore:",omitempty"`
}

// LoadLocation fetches a Location from the DB by ID
func LoadLocation(ctx context.Context, fs *firestore.Client, id id.Location) (Location, error) {
	doc, err := fs.Collection("locations").Doc(string(id)).Get(ctx)
	if err != nil {
		return Location{}, err
	}
	if !doc.Exists() {
		return Location{}, leafbridge.ErrMissingRecord
	}

	var p Location
	err = doc.DataTo(&p)
	if err != nil {
		return Location{}, err
	}

	return p, nil
}

// Delete deletes a Location from the DB by ID
func Delete(ctx context.Context, fs *firestore.Client, id id.Location) error {
	_, err := fs.Collection("locations").Doc(string(id)).Delete(ctx)
	return err
}

// Save stores a Location to the DB
func (c Location) Save(ctx context.Context, fs *firestore.Client) error {
	_, err := fs.Collection("locations").Doc(string(c.ID)).Set(ctx, c)
	return err
}

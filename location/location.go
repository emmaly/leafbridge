package location

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/id"
)

// Location is a note
type Location struct {
	leafbridge.Common
	ID          id.Location
	Type        Type   `firestore:",omitempty"`
	Name        string `firestore:",omitempty"`
	Description string `firestore:",omitempty"`
}

// New returns a new and initialized Location
func New() Location {
	var l Location
	l.ID = id.NewLocation()
	l.Created = time.Now()
	l.Modified = l.Created
	return l
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

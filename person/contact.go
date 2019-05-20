package person

import (
	"bytes"
	"encoding/gob"

	"github.com/emmaly/leafbridge/id"
	"github.com/ipfs/go-datastore"
)

// Contact is a contact method for this Person
type Contact struct {
	ID           id.Contact
	Type         ContactType
	Context      ContactContext
	Notes        []id.Note
	EmailAddress string
	Number       string
	Username     string
	URL          string
	Description  string
	Value        string
	Address      []string
	City         string
	Region       string
	PostalCode   string
	CountryCode  string
}

// LoadContact fetches a Contact from the DB by ID
func LoadContact(ds datastore.Datastore, id id.Contact) (Contact, error) {
	k := datastore.NewKey("CONTACT:" + string(id))
	v, err := ds.Get(k)
	if err != nil {
		return Contact{}, err
	}

	var buf bytes.Buffer
	_, err = buf.Write(v)
	if err != nil {
		return Contact{}, err
	}

	var c Contact
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&c)
	if err != nil {
		return Contact{}, err
	}

	return c, nil
}

// DeleteContact delete a Contact from the DB by ID
func DeleteContact(ds datastore.Datastore, id id.Contact) error {
	k := datastore.NewKey("CONTACT:" + string(id))
	return ds.Delete(k)
}

// Save stores a Contact to the DB
func (c Contact) Save(ds datastore.Datastore) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&c)
	if err != nil {
		return err
	}

	k := datastore.NewKey("CONTACT:" + string(c.ID))
	return ds.Put(k, buf.Bytes())
}

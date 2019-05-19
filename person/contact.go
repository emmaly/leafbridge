package person

import "github.com/emmaly/leafbridge/note"

// Contact is a contact method for this Person
type Contact struct {
	Type         ContactType
	Context      ContactContext
	Notes        []note.Note
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

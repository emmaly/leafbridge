package contact

// Location is an email type of Contact
type Location struct {
	Common
	Address     []string `firestore:",omitempty"`
	City        string   `firestore:",omitempty"`
	Region      string   `firestore:",omitempty"`
	PostalCode  string   `firestore:",omitempty"`
	CountryCode string   `firestore:",omitempty"`
}

// Type returns the type of contact
func (l Location) Type() Type {
	return l.Common.Type
}

// NewPhysicalLocation returns a new and initialized Location contact type
func NewPhysicalLocation() Location {
	var e Location
	e.Common.New()
	e.Common.Type = TypePhysicalLocation
	return e
}

// NewPostalMail returns a new and initialized Location contact type
func NewPostalMail() Location {
	var e Location
	e.Common.New()
	e.Common.Type = TypePostalMail
	return e
}

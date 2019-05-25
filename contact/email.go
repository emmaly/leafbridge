package contact

// Email is an email type of Contact
type Email struct {
	Common
	EmailAddress string `firestore:",omitempty"`
}

// Type returns the type of contact
func (e Email) Type() Type {
	return e.Common.Type
}

// NewEmail returns a new and initialized Email contact type
func NewEmail() Email {
	var e Email
	e.Common.New()
	e.Common.Type = TypeEmail
	return e
}

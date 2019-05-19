package person

// ContactContext indicates the type of context for this Contact.
type ContactContext uint8

// ContactContext constants
const (
	Home ContactContext = iota
	Work
)

var ccNames = map[ContactContext]string{
	Home: "Home",
	Work: "Work",
}

func (cc ContactContext) String() string {
	return ccNames[cc]
}

package contact

// Context indicates the type of context for this Contact.
type Context uint8

// Context constants
const (
	Home Context = iota
	Work
)

var cNames = map[Context]string{
	Home: "Home",
	Work: "Work",
}

func (c Context) String() string {
	return cNames[c]
}

package presence

// Status provides a Person's presence information
type Status struct {
	Availability Availability `firebase:",omitempty"`
	Text         string       `firebase:",omitempty"`
}

// Availability is an indication of a Person's current availability
type Availability uint8

// Availability
const (
	Unknown Availability = iota
	Available
	Busy
	Away
	ExtendedAway
)

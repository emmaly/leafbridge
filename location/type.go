package location

// Type indicates what type of location this Location is.
type Type uint8

// Type constants
const (
	General Type = iota
	Custom
)

var tNames = map[Type]string{
	General: "General",
	Custom:  "Custom",
}

func (t Type) String() string {
	return tNames[t]
}

package note

// Type indicates what type of note this Note is.
type Type uint8

// Type constants
const (
	General Type = iota
	Custom
	Purchase
	Install
	Disposition
	Return
	Repair
	Transfer
)

var tNames = map[Type]string{
	General:     "General",
	Custom:      "Custom",
	Purchase:    "Purchase",
	Install:     "Install",
	Disposition: "Disposition",
	Return:      "Return",
	Repair:      "Repair",
	Transfer:    "Transfer",
}

func (t Type) String() string {
	return tNames[t]
}

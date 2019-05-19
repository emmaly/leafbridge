package note

import (
	"time"

	"github.com/emmaly/leafbridge/id"
)

// Note is a note
type Note struct {
	ID         id.Note
	Title      string
	Body       string
	Created    time.Time
	CreatedBy  id.Person
	Modified   time.Time
	ModifiedBy id.Person
}

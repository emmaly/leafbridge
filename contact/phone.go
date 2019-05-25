package contact

// Phone is an pmail type of Contact
type Phone struct {
	Common
	Number string `firestore:",omitempty"`
}

// Type returns the type of contact
func (p Phone) Type() Type {
	return p.Common.Type
}

// NewVoicePhone returns a new and initialized Phone contact type
func NewVoicePhone() Phone {
	var p Phone
	p.Common.New()
	p.Common.Type = TypeVoicePhone
	return p
}

// NewMobilePhone returns a new and initialized Phone contact type
func NewMobilePhone() Phone {
	var p Phone
	p.Common.New()
	p.Common.Type = TypeMobilePhone
	return p
}

// NewVoiceMessage returns a new and initialized Phone contact type
func NewVoiceMessage() Phone {
	var p Phone
	p.Common.New()
	p.Common.Type = TypeVoiceMessage
	return p
}

// NewFax returns a new and initialized Phone contact type
func NewFax() Phone {
	var p Phone
	p.Common.New()
	p.Common.Type = TypeFax
	return p
}

// NewPager returns a new and initialized Phone contact type
func NewPager() Phone {
	var p Phone
	p.Common.New()
	p.Common.Type = TypePager
	return p
}

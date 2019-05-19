package person

// ContactType indicates what type of contact this Contact is.
type ContactType uint8

// ContactType constants
const (
	Custom           ContactType = iota // CustomContact
	Email                               // EmailContact
	PostalMail                          // LocationContact
	PhysicalLocation                    // LocationContact
	VoicePhone                          // PhoneContact
	MobilePhone                         // PhoneContact
	VoiceMessage                        // PhoneContact
	Fax                                 // PhoneContact
	Pager                               // PhoneContact
	XMPP                                // XMPPContact
	GitHub                              // UsernameContact
	Twitter                             // UsernameContact
	Tumblr                              // UsernameContact
	Diaspora                            // UsernameContact
	Mastodon                            // UsernameContact
	Facebook                            // UsernameContact
)

var ctNames = map[ContactType]string{
	Custom:           "Custom",
	Email:            "Email",
	PostalMail:       "Postal Mail",
	PhysicalLocation: "Physical Location",
	VoicePhone:       "Voice Phone",
	MobilePhone:      "Mobile Phone",
	VoiceMessage:     "Voice Message",
	Fax:              "Fax",
	Pager:            "Pager",
	XMPP:             "XMPP",
	GitHub:           "GitHub",
	Twitter:          "Twitter",
	Tumblr:           "Tumblr",
	Diaspora:         "Diaspora",
	Mastodon:         "Mastodon",
	Facebook:         "Facebook",
}

func (ct ContactType) String() string {
	return ctNames[ct]
}

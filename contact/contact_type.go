package contact

// Type indicates what type of contact this Contact is.
type Type uint8

// Type constants
const (
	TypeCustom           Type = iota // Custom
	TypeEmail                        // Email
	TypePostalMail                   // Location
	TypePhysicalLocation             // Location
	TypeVoicePhone                   // Phone
	TypeMobilePhone                  // Phone
	TypeVoiceMessage                 // Phone
	TypeFax                          // Phone
	TypePager                        // Phone
	TypeXMPP                         // XMPP
	TypeGitHub                       // Username
	TypeTwitter                      // Username
	TypeTumblr                       // Username
	TypeDiaspora                     // Username
	TypeMastodon                     // Username
	TypeFacebook                     // Username
)

var tNames = map[Type]string{
	TypeCustom:           "Custom",
	TypeEmail:            "Email",
	TypePostalMail:       "Postal Mail",
	TypePhysicalLocation: "Physical Location",
	TypeVoicePhone:       "Voice Phone",
	TypeMobilePhone:      "Mobile Phone",
	TypeVoiceMessage:     "Voice Message",
	TypeFax:              "Fax",
	TypePager:            "Pager",
	TypeXMPP:             "XMPP",
	TypeGitHub:           "GitHub",
	TypeTwitter:          "Twitter",
	TypeTumblr:           "Tumblr",
	TypeDiaspora:         "Diaspora",
	TypeMastodon:         "Mastodon",
	TypeFacebook:         "Facebook",
}

func (t Type) String() string {
	return tNames[t]
}

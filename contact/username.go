package contact

// Username is a username type of Contact
type Username struct {
	Common
	Username string `firestore:",omitempty"`
}

// Type returns the type of contact
func (u Username) Type() Type {
	return u.Common.Type
}

// URL returns the URL to the account, if there is one available
func (u Username) URL() string {
	switch u.Common.Type {
	case TypeGitHub:
		return "https://github.com/" + u.Username
	case TypeTwitter:
		return "https://twitter.com/" + u.Username
	case TypeTumblr:
		return "https://" + u.Username + "tumblr.com/"
	case TypeFacebook:
		return "https://www.facebook.com/" + u.Username
	default:
		return ""
	}
}

// NewGitHub returns a new and initialized Username contact type
func NewGitHub() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeGitHub
	return u
}

// NewTwitter returns a new and initialized Username contact type
func NewTwitter() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeTwitter
	return u
}

// NewTumblr returns a new and initialized Username contact type
func NewTumblr() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeTumblr
	return u
}

// NewDiaspora returns a new and initialized Username contact type
func NewDiaspora() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeDiaspora
	return u
}

// NewMastodon returns a new and initialized Username contact type
func NewMastodon() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeMastodon
	return u
}

// NewFacebook returns a new and initialized Username contact type
func NewFacebook() Username {
	var u Username
	u.Common.New()
	u.Common.Type = TypeFacebook
	return u
}

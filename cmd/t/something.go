package main

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"github.com/emmaly/leafbridge/contact"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
	"github.com/emmaly/leafbridge/person"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("creds.json"))
	if err != nil {
		log.Fatal(err)
	}
	fs, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	createdTime := time.Now()
	creatorPersonID := id.NewPerson()

	pID := id.NewPerson()
	p := person.Person{
		ID: pID,
		Name: person.Name{
			Prefix:  "Mr.",
			Family:  "Flintstone",
			Given:   "Frederick",
			Ordinal: "III",
			Suffix:  "PE",
			Format:  person.WesternOrder,
			// Format: person.EasternOrder,
		},
	}

	email := contact.Email{
		ID:           id.NewContact(),
		Type:         contact.TypeEmail,
		Context:      contact.Work,
		EmailAddress: "fred.flintstone@bedrock.quarry",
	}
	p.Contacts = append(p.Contacts, email.AsContact())

	mail := contact.Location{
		ID:      id.NewContact(),
		Type:    contact.TypePostalMail,
		Context: contact.Work,
		Address: []string{
			"Mail Stop 132",
			"123 Main St",
		},
		City:        "Onett",
		Region:      "EB",
		PostalCode:  "55443",
		CountryCode: "US",
	}
	p.Contacts = append(p.Contacts, mail.AsContact())

	chat := contact.Contact{
		ID:          id.NewContact(),
		Type:        contact.TypeCustom,
		Context:     contact.Work,
		Description: "Favorite Chat Server",
		URL:         "https://chat.example.com/flanges",
	}
	p.Contacts = append(p.Contacts, chat.AsContact())

	cell := contact.Phone{
		ID:      id.NewContact(),
		Type:    contact.TypeMobilePhone,
		Context: contact.Work,
		Number:  "+1 555-555-1212",
	}
	p.Contacts = append(p.Contacts, cell.AsContact())

	diaspora := contact.Username{
		ID:       id.NewContact(),
		Type:     contact.TypeDiaspora,
		Context:  contact.Home,
		Username: "fred.flintstone@bedrock.isp",
	}
	p.Contacts = append(p.Contacts, diaspora.AsContact())

	github := contact.Username{
		ID:       id.NewContact(),
		Type:     contact.TypeGitHub,
		Context:  contact.Work,
		Username: "FredFlintstone",
	}
	p.Contacts = append(p.Contacts, github.AsContact())

	twitter := contact.Username{
		ID:       id.NewContact(),
		Type:     contact.TypeTwitter,
		Context:  contact.Home,
		Username: "FredFlintstone",
	}
	p.Contacts = append(p.Contacts, twitter.AsContact())

	facebook := contact.Username{
		ID:       id.NewContact(),
		Type:     contact.TypeFacebook,
		Context:  contact.Home,
		Username: "FredFlintstone",
	}
	p.Contacts = append(p.Contacts, facebook.AsContact())

	musicNote := note.Note{
		ID:         id.NewNote(),
		Title:      "Likes Rock and Roll",
		Body:       "Seems to like music that gets his stone wheels spinning.",
		Created:    createdTime,
		CreatedBy:  creatorPersonID,
		Modified:   createdTime,
		ModifiedBy: creatorPersonID,
	}
	p.Notes = append(p.Notes, musicNote)

	err = p.Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%+v\n\n", p)

	for _, v := range p.Contacts {
		switch v.Type {
		case contact.TypeEmail:
			fmt.Printf("\n%+v\n\n", v.AsEmail())
		case contact.TypePostalMail,
			contact.TypePhysicalLocation:
			fmt.Printf("\n%+v\n\n", v.AsLocation())
		case contact.TypeMobilePhone,
			contact.TypeVoicePhone,
			contact.TypeVoiceMessage,
			contact.TypeFax,
			contact.TypePager:
			fmt.Printf("\n%+v\n\n", v.AsPhone())
		case contact.TypeGitHub,
			contact.TypeTwitter,
			contact.TypeTumblr,
			contact.TypeDiaspora,
			contact.TypeMastodon,
			contact.TypeFacebook:
			fmt.Printf("\n%+v\n\t\t%s\n\n", v.AsUsername(), v.AsUsername().URL())
		default:
			fmt.Printf("\nDon't know how to show the contact type called %s.\n\n", v.Type)
		}
	}
}

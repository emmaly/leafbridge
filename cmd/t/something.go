package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/emmaly/leafbridge"
	"github.com/emmaly/leafbridge/contact"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/location"
	"github.com/emmaly/leafbridge/person"
	"github.com/emmaly/leafbridge/presence"
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
		Title:          "Geological Engineer",
		PresenceStatus: presence.Status{},
	}

	office := location.NewLocation()
	office.Name = "Slate HQ"
	p.PrimaryLocation = office.ID

	email := contact.NewEmail()
	email.Context = contact.Work
	email.EmailAddress = "fred.flintstone@bedrock.quarry"
	p.Contacts = append(p.Contacts, email)

	mail := contact.NewPostalMail()
	mail.Context = contact.Work
	mail.Address = []string{
		"Mail Stop 132",
		"123 Main St",
	}
	mail.City = "Onett"
	mail.Region = "EB"
	mail.PostalCode = "55443"
	mail.CountryCode = "US"
	p.Contacts = append(p.Contacts, mail)

	// chat := contact.Contact{
	// 	ID:          id.NewContact(),
	// 	Type:        contact.TypeCustom,
	// 	Context:     contact.Work,
	// 	Description: "Favorite Chat Server",
	// 	URL:         "https://chat.example.com/flanges",
	// 	Common: leafbridge.Common{
	// 		Created:    createdTime,
	// 		CreatedBy:  creatorPersonID,
	// 		Modified:   createdTime,
	// 		ModifiedBy: creatorPersonID,
	// 	},
	// }
	// p.Contacts = append(p.Contacts, chat.AsContact())

	cell := contact.NewMobilePhone()
	cell.Context = contact.Work
	cell.Number = "+1 555-555-1212"
	p.Contacts = append(p.Contacts, cell)

	diaspora := contact.NewDiaspora()
	diaspora.Context = contact.Home
	diaspora.Username = "fred.flintstone@bedrock.isp"
	p.Contacts = append(p.Contacts, diaspora)

	github := contact.NewGitHub()
	github.Context = contact.Work
	github.Username = "FredFlintstone"
	p.Contacts = append(p.Contacts, github)

	twitter := contact.NewTwitter()
	twitter.Context = contact.Home
	twitter.Username = "FredFlintstone"
	p.Contacts = append(p.Contacts, twitter)

	facebook := contact.NewFacebook()
	facebook.Context = contact.Home
	facebook.Username = "FredFlintstone"
	p.Contacts = append(p.Contacts, facebook)

	musicNote := leafbridge.NewNote()
	musicNote.Title = "Likes Rock and Roll"
	musicNote.Body = "Seems to like music that gets his stone wheels spinning."
	p.Notes = append(p.Notes, musicNote)

	if false {
		err = p.Save(ctx, fs)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("\n%+v\n\n", p)

	for _, v := range p.Contacts {
		switch c := v.(type) {
		case contact.Email:
			fmt.Printf("\nEmail Type:\n%+v\n\n", c)
		case contact.Location:
			fmt.Printf("\nLocation Type:\n%+v\n\n", c)
		case contact.Phone:
			fmt.Printf("\nPhone Type:\n%+v\n\n", c)
		case contact.Username:
			fmt.Printf("\nUsername Type:\n%+v\n\t\t%s\n\n", c, c.URL())
		default:
			fmt.Printf("\nUnknown Type:\n%+v\n\n", c)
		}
	}
}

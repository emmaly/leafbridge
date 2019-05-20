package main

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
	"github.com/emmaly/leafbridge/person"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile("creds.json"))
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
	err = p.Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}

	email := person.EmailContact{
		ID:           id.NewContact(),
		Type:         person.Email,
		Context:      person.Work,
		EmailAddress: "fred.flintstone@bedrock.quarry",
	}
	err = email.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, email.ID)

	mail := person.LocationContact{
		ID:      id.NewContact(),
		Type:    person.PostalMail,
		Context: person.Work,
		Address: []string{
			"Mail Stop 132",
			"123 Main St",
		},
		City:        "Onett",
		Region:      "EB",
		PostalCode:  "55443",
		CountryCode: "US",
	}
	err = mail.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, mail.ID)

	chat := person.Contact{
		ID:          id.NewContact(),
		Type:        person.Custom,
		Context:     person.Work,
		Description: "Favorite Chat Server",
		URL:         "https://chat.example.com/flanges",
	}
	err = chat.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, chat.ID)

	cell := person.PhoneContact{
		ID:      id.NewContact(),
		Type:    person.MobilePhone,
		Context: person.Work,
		Number:  "+1 555-555-1212",
	}
	err = cell.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, cell.ID)

	diaspora := person.UsernameContact{
		ID:       id.NewContact(),
		Type:     person.Diaspora,
		Context:  person.Home,
		Username: "fred.flintstone@bedrock.isp",
	}
	err = diaspora.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, diaspora.ID)

	github := person.UsernameContact{
		ID:       id.NewContact(),
		Type:     person.GitHub,
		Context:  person.Work,
		Username: "FredFlintstone",
	}
	err = github.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, github.ID)

	twitter := person.UsernameContact{
		ID:       id.NewContact(),
		Type:     person.Twitter,
		Context:  person.Home,
		Username: "FredFlintstone",
	}
	err = twitter.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, twitter.ID)

	facebook := person.UsernameContact{
		ID:       id.NewContact(),
		Type:     person.Facebook,
		Context:  person.Home,
		Username: "FredFlintstone",
	}
	err = facebook.Contact().Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Contacts = append(p.Contacts, facebook.ID)

	musicNote := note.Note{
		ID:         id.NewNote(),
		Title:      "Likes Rock and Roll",
		Body:       "Seems to like music that gets his stone wheels spinning.",
		Created:    createdTime,
		CreatedBy:  creatorPersonID,
		Modified:   createdTime,
		ModifiedBy: creatorPersonID,
	}
	err = musicNote.Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	p.Notes = append(p.Notes, musicNote.ID)

	err = p.Save(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%+v\n\n", p)

	for _, id := range p.Contacts {
		v, err := person.LoadContact(ctx, fs, id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch v.Type {
		case person.Email:
			fmt.Printf("\n%+v\n\n", v.EmailContact())
		case person.PostalMail,
			person.PhysicalLocation:
			fmt.Printf("\n%+v\n\n", v.LocationContact())
		case person.MobilePhone,
			person.VoicePhone,
			person.VoiceMessage,
			person.Fax,
			person.Pager:
			fmt.Printf("\n%+v\n\n", v.PhoneContact())
		case person.GitHub,
			person.Twitter,
			person.Tumblr,
			person.Diaspora,
			person.Mastodon,
			person.Facebook:
			fmt.Printf("\n%+v\n\t\t%s\n\n", v.UsernameContact(), v.UsernameContact().URL())
		default:
			fmt.Printf("\nDon't know how to show the contact type called %s.\n\n", v.Type)
		}
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/emmaly/leafbridge/id"
	"github.com/emmaly/leafbridge/note"
	"github.com/emmaly/leafbridge/person"
)

func main() {
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
		Contact: []person.Contact{
			person.EmailContact{
				Type:    person.Email,
				Context: person.Work,
				Notes: []note.Note{
					{
						ID:         id.NewNote(),
						Body:       "Only use to ask about the quarry.",
						Created:    createdTime,
						CreatedBy:  creatorPersonID,
						Modified:   createdTime,
						ModifiedBy: creatorPersonID,
					},
				},
				EmailAddress: "fred.flintstone@bedrock.quarry",
			}.Contact(),
			person.LocationContact{
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
			}.Contact(),
			person.Contact{
				Type:        person.Custom,
				Context:     person.Work,
				Description: "Favorite Chat Server",
				URL:         "https://chat.example.com/flanges",
			},
			person.PhoneContact{
				Type:    person.MobilePhone,
				Context: person.Work,
				Number:  "+1 555-555-1212",
			}.Contact(),
			person.UsernameContact{
				Type:     person.Diaspora,
				Context:  person.Home,
				Username: "fred.flintstone@bedrock.isp",
			}.Contact(),
			person.UsernameContact{
				Type:     person.GitHub,
				Context:  person.Work,
				Username: "FredFlintstone",
			}.Contact(),
			person.UsernameContact{
				Type:     person.Twitter,
				Context:  person.Home,
				Username: "FredFlintstone",
			}.Contact(),
			person.UsernameContact{
				Type:     person.Facebook,
				Context:  person.Home,
				Username: "FredFlintstone",
				Notes: []note.Note{
					{
						ID:         id.NewNote(),
						Title:      "Use sparingly",
						Body:       "They don't want to be spammed via Facebook.",
						Created:    createdTime,
						CreatedBy:  creatorPersonID,
						Modified:   createdTime,
						ModifiedBy: creatorPersonID,
					},
				},
			}.Contact(),
		},
		Notes: []note.Note{
			{
				ID:         id.NewNote(),
				Title:      "Likes Rock and Roll",
				Body:       "Seems to like music that gets his stone wheels spinning.",
				Created:    createdTime,
				CreatedBy:  creatorPersonID,
				Modified:   createdTime,
				ModifiedBy: creatorPersonID,
			},
		},
	}

	fmt.Printf("\n%+v\n\n", p)

	for _, v := range p.Contact {
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

package email

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
)

func GetMessageBody(msg *imap.Message, section imap.BodySectionName) *mail.Part {
	log.Println("Getting message body...")

	// Get the whole message body
	r := msg.GetBody(&section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}
	if subject, err := header.Subject(); err == nil {
		log.Println("Subject:", subject)
	}

	// Get HTML part of body
	mr.NextPart()
	p, err := mr.NextPart()

	if err != nil {
		log.Fatal(err)
	}

	return p
}

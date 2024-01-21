package email

import (
	"io"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
)

func GetMessageBody(msg *imap.Message, section imap.BodySectionName) {
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

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := io.ReadAll(p.Body)
			log.Printf("Got text: %v", string(b))
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Printf("Got attachment: %v", filename)
		}
	}
}

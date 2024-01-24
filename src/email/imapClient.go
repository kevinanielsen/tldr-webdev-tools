package email

import (
	"log"

	"github.com/emersion/go-imap/client"
	"github.com/kevinanielsen/tldr-webdev-tools/src/initializers"
)

var Client *client.Client

func ConnectToImap() {
	log.Println("Connecting to IMAP server...")

	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Login
	log.Println("Logging in...")
	if err := c.Login(initializers.GMAIL_USERNAME, initializers.GMAIL_APP_PASSWORD); err != nil {
		log.Fatal(err)
	}

	Client = c
}

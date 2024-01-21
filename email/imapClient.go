package email

import (
	"log"

	"github.com/emersion/go-imap/client"
	"github.com/kevinanielsen/tldr-webdev-tools/initializers"
)

var Client *client.Client

func ConnectToImap() {
	log.Println("Connecting to IMAP server...")

	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected!")

	defer c.Logout()

	// Login
	if err := c.Login("keveran@gmail.com", initializers.GMAIL_APP_PASSWORD); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in!")

	Client = c

}

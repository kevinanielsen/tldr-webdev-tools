package main

import (
	"github.com/kevinanielsen/tldr-webdev-tools/email"
	"github.com/kevinanielsen/tldr-webdev-tools/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	email.ConnectToImap()
}

func main() {
	msg, section := email.GetMessage(email.Client)
	email.GetMessageBody(msg, section)

	defer email.Client.Logout()
}

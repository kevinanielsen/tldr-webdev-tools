package main

import (
	"log"

	"github.com/kevinanielsen/tldr-webdev-tools/email"
	"github.com/kevinanielsen/tldr-webdev-tools/email/utils"
	"github.com/kevinanielsen/tldr-webdev-tools/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	email.ConnectToImap()
}

func main() {
	msg, section := email.GetMessage(email.Client)
	htmlPart := email.GetMessageBody(msg, section)
	toolsSection := utils.FindSection(htmlPart)

	log.Println(toolsSection)

	defer email.Client.Logout()
}

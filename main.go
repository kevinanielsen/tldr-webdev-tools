package main

import (
	"fmt"

	"github.com/kevinanielsen/tldr-webdev-tools/src/database"
	"github.com/kevinanielsen/tldr-webdev-tools/src/email"
	"github.com/kevinanielsen/tldr-webdev-tools/src/email/utils"
	"github.com/kevinanielsen/tldr-webdev-tools/src/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	database.ConnectToDatabase()
	email.ConnectToImap()
}

func main() {
	msg, section := email.GetMessage(email.Client)
	htmlPart := email.GetMessageBody(msg, section)
	tools := utils.FindSection(htmlPart)
	for _, v := range tools {
		fmt.Println(v.Name)
		fmt.Println(v.URL)
		fmt.Println(v.Description)
		fmt.Println("————————————————————————————————————————————————————")

	}

	defer email.Client.Logout()
}

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

}

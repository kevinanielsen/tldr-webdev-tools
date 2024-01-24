package utils

import (
	"fmt"
	"io"
	"strings"

	"github.com/emersion/go-message/mail"
	"github.com/k3a/html2text"
)

type Tool struct {
	Name        string
	URL         string
	Description string
}

// getToolsStr extracts the "Launches & Tools" section from the
// provided HTML string by converting to plain text, splitting on
// "Launches & Tools", and returning the text up to "Miscellaneous".
func getToolsStr(str string) string {
	htmlPlain := html2text.HTML2Text(str)
	launchesToEnd := strings.Split(htmlPlain, "Launches & Tools")[1]
	return strings.Split(launchesToEnd, "Miscellaneous")[0]
}

func FindSection(p *mail.Part) []Tool {
	b, _ := io.ReadAll(p.Body)
	emailBody := string(b)
	toolsStr := getToolsStr(emailBody)

	var toolsArr []Tool
	for i, v := range strings.Split(toolsStr, "https://") {
		if i == 0 {
			continue
		}
		current := fmt.Sprintf("https://%s", v)

		currentArr := strings.Split(current, "\n")

		toolsArr = append(toolsArr, Tool{
			Name:        strings.TrimSpace(strings.Split(currentArr[2], " ")[1]),
			URL:         strings.TrimSpace(currentArr[0]),
			Description: strings.TrimSpace(currentArr[2]),
		})
	}

	return toolsArr
}

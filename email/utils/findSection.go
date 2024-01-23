package utils

import (
	"io"
	"strings"

	"github.com/emersion/go-message/mail"
)

func FindSection(p *mail.Part) string {
	b, _ := io.ReadAll(p.Body)
	emailBody := string(b)
	htmlPart := strings.Split(emailBody, "Launches &amp; Tools")[1]
	toolsToEnd := strings.Split(htmlPart, `<table style="table-layout: fixed; width: 100%;" width="100%"><tbody>`)
	// toolsSection := strings.Split(toolsToEnd[1], `</td></tr></tbody></table>`)[0]

	return toolsToEnd[1]
}

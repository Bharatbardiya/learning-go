package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string // demo a loop
	UnreadCount   int
}

func main() {
	emailTemplate := `
	Subject: {{.Subject}}
	Body: {{.Body}}
	
	{{if .Items}}
		Related Items:
	{{range .Items}}
			{{.}}
	{{end}}
	{{end}}

	{{if gt .UnreadCount 0}}
	You have {{.UnreadCount}} unreads.
	{{else}}
	you have no messages
	{{end}}`

	emailData := EmailData{
		RecipientName: "bharat@google.com",
		SenderName:    "Bharat's auto sender",
		Subject:       "Weekly update",
		Body:          "here is your weekly updates.",
		Items:         []string{"Report A", "Summary B", "Document C"},
		UnreadCount:   3,
	}

	templ, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var output strings.Builder

	err = templ.Execute(&output, emailData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Email template successfully executed")
	fmt.Println(strings.ToUpper(output.String()))
}

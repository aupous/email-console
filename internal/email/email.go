package email

import (
	"email-console/internal/customer"
	"strings"
	"time"
)

type EmailTemplate struct {
	From     string
	Subject  string
	MimeType string
	Body     string
}

const timeFormat = "02 Feb 2006"

type EmailData struct {
	EmailTemplate
	To string
}

func (et EmailTemplate) CustomerEmail(cus customer.Customer) EmailData {
	emailData := EmailData{
		EmailTemplate: et,
		To:            cus.Email,
	}
	emailData.Body = strings.Replace(emailData.Body, "{{TITLE}}", cus.Title, 1)
	emailData.Body = strings.Replace(emailData.Body, "{{FIRST_NAME}}", cus.FirstName, 1)
	emailData.Body = strings.Replace(emailData.Body, "{{LAST_NAME}}", cus.LastName, 1)
	emailData.Body = strings.Replace(emailData.Body, "{{TODAY}}", time.Now().Format(timeFormat), 1)
	return emailData
}

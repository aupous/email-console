package email

import (
	"email-console/internal/customer"
	"fmt"
	"testing"
	"time"
)

var testCustomer = customer.Customer{
	Title:     "Mr",
	FirstName: "John",
	LastName:  "Smith",
	Email:     "john.smith@example.com",
}
var secondTestCustomer = customer.Customer{
	Title:     "Mrs",
	FirstName: "Michelle",
	LastName:  "Smith",
	Email:     "michelle.smith@example.com",
}

func TestEmailTemplate_CustomerEmail(t *testing.T) {
	tests := []struct {
		name          string
		emailTemplate EmailTemplate
		expected      EmailData
	}{
		{
			name: "customer email as expected",
			emailTemplate: EmailTemplate{
				From:     "The Marketing Team<marketing@example.com>",
				Subject:  "A new product is being launched soon...",
				MimeType: "text/plain",
				Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}},\nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			},
			expected: EmailData{
				To: testCustomer.Email,
				EmailTemplate: EmailTemplate{
					From:     "The Marketing Team<marketing@example.com>",
					Subject:  "A new product is being launched soon...",
					MimeType: "text/plain",
					Body:     fmt.Sprintf("Hi %s %s %s,\nToday, %s, we would like to tell you that... Sincerely,\nThe Marketing Team", testCustomer.Title, testCustomer.FirstName, testCustomer.LastName, time.Now().Format("25 Feb 2015")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.emailTemplate.CustomerEmail(testCustomer)
			if actual.From != tt.expected.From ||
				actual.To != tt.expected.To ||
				actual.Subject != tt.expected.Subject ||
				actual.MimeType != tt.expected.MimeType ||
				actual.Body != tt.expected.Body {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}

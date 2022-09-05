package email

import (
	"email-console/internal/customer"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	testTemplateJSON = `{
	  "from": "The Marketing Team<marketing@example.com>",
	  "subject": "A new product is being launched soon...",
	  "mimeType": "text/plain",
	  "body": "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}},\nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team"
	}`
	errTemplateJSON       = `example`
	templatePath          = "email_template.json"
	errTemplatePath       = "err_email_template.json"
	notExistsTemplatePath = "not_exists_email_template.json"
	emailsPath            = "emails"
)

func TestNewEmailService(t *testing.T) {
	// setup data to test
	if err := ioutil.WriteFile(templatePath, []byte(testTemplateJSON), 0644); err != nil {
		panic("error when write data to email template file")
	}
	defer os.Remove(templatePath)
	if err := ioutil.WriteFile(errTemplatePath, []byte(errTemplateJSON), 0644); err != nil {
		panic("error when write data to error email template file")
	}
	defer os.Remove(errTemplatePath)

	tests := []struct {
		name       string
		input      []string
		wantErr    bool
		errMessage string
	}{
		{
			name:    "success",
			input:   []string{templatePath, emailsPath},
			wantErr: false,
		},
		{
			name:       "error when data in email template file is not valid json",
			input:      []string{errTemplatePath, emailsPath},
			wantErr:    true,
			errMessage: errUnmarshalEmailTemplate,
		},
		{
			name:       "error when email template path not exists",
			input:      []string{notExistsTemplatePath, emailsPath},
			wantErr:    true,
			errMessage: errReadEmailTemplate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewEmailService(tt.input[0], tt.input[1])
			if !tt.wantErr && actual == nil {
				t.Error("expect success and return an email service but got nil")
			} else if tt.wantErr && !strings.Contains(err.Error(), tt.errMessage) {
				t.Errorf("expect error with message %s but got error %s", tt.errMessage, err.Error())
			}
		})
	}
}

func TestEmailService_SendMail(t *testing.T) {
	if err := ioutil.WriteFile(templatePath, []byte(testTemplateJSON), 0644); err != nil {
		panic("error when write data to email template file")
	}
	defer os.Remove(templatePath)
	if err := os.Mkdir(emailsPath, 0755); err != nil {
		panic("error when create emails folder")
	}
	defer os.RemoveAll(emailsPath)
	emailService, err := NewEmailService(templatePath, emailsPath)
	if err != nil {
		t.Error(errors.Wrap(err, "error when create email service"))
		return
	}

	tests := []struct {
		name    string
		input   customer.Customer
		wantErr bool
	}{
		{
			name:  "success",
			input: testCustomer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = emailService.SendMail(tt.input)
			if err != nil {
				t.Errorf("expect success but got error %s", err.Error())
			} else {
				files, err := ioutil.ReadDir(emailsPath)
				if err != nil {
					t.Error("error when read files in emails folder")
					return
				}
				exists := false
				for _, file := range files {
					if strings.Contains(file.Name(), tt.input.Email) {
						exists = true
					}
				}
				if !exists {
					t.Error("expect output emails folder to contain file named with customer email but not")
				}
			}
		})
	}
}

func TestEmailService_BulkSendMails(t *testing.T) {
	if err := ioutil.WriteFile(templatePath, []byte(testTemplateJSON), 0644); err != nil {
		panic("error when write data to email template file")
	}
	defer os.Remove(templatePath)
	if err := os.Mkdir(emailsPath, 0755); err != nil {
		panic("error when create emails folder")
	}
	defer os.RemoveAll(emailsPath)
	emailService, err := NewEmailService(templatePath, emailsPath)
	if err != nil {
		t.Error(errors.Wrap(err, "error when create email service"))
		return
	}

	tests := []struct {
		name    string
		input   []customer.Customer
		wantErr bool
	}{
		{
			name:  "success",
			input: []customer.Customer{testCustomer, secondTestCustomer},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emailService.BulkSendMails(tt.input)
			files, err := ioutil.ReadDir(emailsPath)
			if err != nil {
				t.Error("error when read files in emails folder")
				return
			}
			for _, cus := range tt.input {
				exists := false
				for _, file := range files {
					if strings.Contains(file.Name(), cus.Email) {
						exists = true
					}
				}
				if !exists {
					t.Errorf("expect output emails folder to contain file named with customer email %s but not", cus.Email)
				}
			}
		})
	}
}

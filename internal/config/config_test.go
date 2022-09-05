package config

import (
	"os"
	"strings"
	"testing"
)

const (
	testEmailTemplatePath    = "email_template.json"
	testErrEmailTemplatePath = "err_email_template.json"
	testCustomersPath        = "customers.csv"
	testErrCustomersPath     = "err_customers.csv"
	testEmailsPath           = "emails"
	testErrEmailsPath        = "err_emails"
	testErrorsPath           = "errors.csv"
)

func TestConfig_Verify(t *testing.T) {
	if _, err := os.Create(testEmailTemplatePath); err != nil {
		panic("error when create test email template file")
	}
	defer os.Remove(testEmailTemplatePath)
	if _, err := os.Create(testCustomersPath); err != nil {
		panic("error when create test customers file")
	}
	defer os.Remove(testCustomersPath)
	if err := os.Mkdir(testEmailsPath, 0755); err != nil {
		panic("error when create test emails folder")
	}
	defer os.Remove(testEmailsPath)
	tests := []struct {
		name         string
		config       Config
		wantErr      bool
		errorMessage string
	}{
		{
			name: "config is valid",
			config: Config{
				EmailTemplatePath:  testEmailTemplatePath,
				CustomersPath:      testCustomersPath,
				OutputEmailsPath:   testEmailsPath,
				ErrorCustomersPath: testErrorsPath,
			},
			wantErr: false,
		},
		{
			name: "email template path is not exists",
			config: Config{
				EmailTemplatePath:  testErrEmailTemplatePath,
				CustomersPath:      testCustomersPath,
				OutputEmailsPath:   testEmailsPath,
				ErrorCustomersPath: testErrorsPath,
			},
			wantErr:      true,
			errorMessage: errorEmailTemplatePath,
		},
		{
			name: "customers path is not exists",
			config: Config{
				EmailTemplatePath:  testEmailTemplatePath,
				CustomersPath:      testErrCustomersPath,
				OutputEmailsPath:   testEmailsPath,
				ErrorCustomersPath: testErrorsPath,
			},
			wantErr:      true,
			errorMessage: errorCustomersPath,
		},
		{
			name: "output emails path is not exists",
			config: Config{
				EmailTemplatePath:  testEmailTemplatePath,
				CustomersPath:      testCustomersPath,
				OutputEmailsPath:   testErrEmailsPath,
				ErrorCustomersPath: testErrorsPath,
			},
			wantErr:      true,
			errorMessage: errorOutputEmailsPath,
		},
		{
			name: "output emails path is not a directory",
			config: Config{
				EmailTemplatePath:  testEmailTemplatePath,
				CustomersPath:      testCustomersPath,
				OutputEmailsPath:   testEmailTemplatePath,
				ErrorCustomersPath: testErrorsPath,
			},
			wantErr:      true,
			errorMessage: outputEmailsPathIsNotDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Verify()
			if !tt.wantErr && err != nil {
				t.Errorf("expect verify success but got error %s", err.Error())
			} else if tt.wantErr {
				if err == nil {
					t.Error("expect verify error but got nil")
				}
				if !strings.Contains(err.Error(), tt.errorMessage) {
					t.Errorf("expect verify error to contain %s but got error %s", tt.errorMessage, err.Error())
				}
			}
		})
	}
}

package config

import (
	"path"
	"strings"
	"testing"
)

func TestConfig_Verify(t *testing.T) {
	directoryPath := "../../"
	emailTemplatePath := path.Join(directoryPath, "email_template.json")
	errEmailTemplatePath := path.Join(directoryPath, "err_email_template.json")
	customersPath := path.Join(directoryPath, "customers.csv")
	errCustomersPath := path.Join(directoryPath, "err_customers.csv")
	emailsPath := path.Join(directoryPath, "emails")
	errEmailsPath := path.Join(directoryPath, "err_emails")
	errorsPath := path.Join(directoryPath, "errors.csv")
	tests := []struct {
		name         string
		config       Config
		wantErr      bool
		errorMessage string
	}{
		{
			name: "config is valid",
			config: Config{
				EmailTemplatePath:  emailTemplatePath,
				CustomersPath:      customersPath,
				OutputEmailsPath:   emailsPath,
				ErrorCustomersPath: errorsPath,
			},
			wantErr: false,
		},
		{
			name: "email template path is not exists",
			config: Config{
				EmailTemplatePath:  errEmailTemplatePath,
				CustomersPath:      customersPath,
				OutputEmailsPath:   emailsPath,
				ErrorCustomersPath: errorsPath,
			},
			wantErr:      true,
			errorMessage: errorEmailTemplatePath,
		},
		{
			name: "customers path is not exists",
			config: Config{
				EmailTemplatePath:  emailTemplatePath,
				CustomersPath:      errCustomersPath,
				OutputEmailsPath:   emailsPath,
				ErrorCustomersPath: errorsPath,
			},
			wantErr:      true,
			errorMessage: errorCustomersPath,
		},
		{
			name: "output emails path is not exists",
			config: Config{
				EmailTemplatePath:  emailTemplatePath,
				CustomersPath:      customersPath,
				OutputEmailsPath:   errEmailsPath,
				ErrorCustomersPath: errorsPath,
			},
			wantErr:      true,
			errorMessage: errorOutputEmailsPath,
		},
		{
			name: "output emails path is not a directory",
			config: Config{
				EmailTemplatePath:  emailTemplatePath,
				CustomersPath:      customersPath,
				OutputEmailsPath:   emailTemplatePath,
				ErrorCustomersPath: errorsPath,
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

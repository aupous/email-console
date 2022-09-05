package console

import (
	"email-console/internal/config"
	"email-console/internal/customer"
	"email-console/internal/email"
	"github.com/pkg/errors"
	"os"
)

type Console struct{}

func (c Console) Start() {
	args := os.Args
	if len(args) < 5 {
		panic("not enough arguments")
	}
	cfg := config.Config{
		EmailTemplatePath:  args[1],
		CustomersPath:      args[2],
		OutputEmailsPath:   args[3],
		ErrorCustomersPath: args[4],
	}
	if err := cfg.Verify(); err != nil {
		panic(errors.Wrap(err, "arguments are not valid"))
	}

	emailService, err := email.NewEmailService(cfg.EmailTemplatePath, cfg.OutputEmailsPath)
	if err != nil {
		panic(errors.Wrap(err, "error when setup email service"))
	}

	customers, err := customer.LoadFromCsv(cfg.CustomersPath)
	if err != nil {
		panic(errors.Wrap(err, "error when read customers from csv"))
	}
	errorCustomers := make([]customer.Customer, 0)
	validCustomers := make([]customer.Customer, 0)
	for _, cus := range customers {
		if cus.EmailValid() {
			validCustomers = append(validCustomers, cus)
		} else {
			errorCustomers = append(errorCustomers, cus)
		}
	}

	// Send emails to valid customers
	emailService.BulkSendMails(validCustomers)

	// Write error customers to error file
	if err = customer.WriteCustomersToCsv(cfg.ErrorCustomersPath, errorCustomers); err != nil {
		panic(errors.Wrap(err, "error when write error customers to file"))
	}
}

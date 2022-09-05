package console

import (
	"email-console/internal/config"
	"email-console/internal/customer"
	"email-console/internal/email"
	"github.com/pkg/errors"
	"os"
)

type emailService interface {
	SendMail(customer.Customer) error
	BulkSendMails([]customer.Customer)
}

type Console struct {
	config       *config.Config
	emailService emailService
}

func (c Console) Start() {
	c.setup()
	customers, err := customer.LoadFromCsv(c.config.CustomersPath)
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
	c.emailService.BulkSendMails(validCustomers)

	// Write error customers to error file
	if err = customer.WriteCustomersToCsv(c.config.ErrorCustomersPath, errorCustomers); err != nil {
		panic(errors.Wrap(err, "error when write error customers to file"))
	}
}

func (c Console) setup() {
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

	// if you want to send email via SMTP or REST API, write another email service and replace it here
	emService, err := email.NewEmailService(cfg.EmailTemplatePath, cfg.OutputEmailsPath)
	if err != nil {
		panic(errors.Wrap(err, "error when setup email service"))
	}
	c.config = &cfg
	c.emailService = emService
}

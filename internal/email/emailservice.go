package email

import (
	"email-console/internal/customer"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"path"
	"sync"
	"time"
)

type EmailService struct {
	emailTemplate    EmailData
	outputEmailsPath string
}

const errReadEmailTemplate = "error when read email template file"
const errUnmarshalEmailTemplate = "error during unmarshal email template data"

func NewEmailService(emailTemplatePath, outputEmailsPath string) (*EmailService, error) {
	// Read email template file to load email template dât
	templateData, err := ioutil.ReadFile(emailTemplatePath)
	if err != nil {
		return nil, errors.Wrap(err, errReadEmailTemplate)
	}

	// Now let's unmarshall the data into `payload`
	var emailTemplate EmailData
	err = json.Unmarshal(templateData, &emailTemplate)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalEmailTemplate)
	}
	return &EmailService{
		emailTemplate:    emailTemplate,
		outputEmailsPath: outputEmailsPath,
	}, nil
}

func (em EmailService) SendMail(cus customer.Customer) error {
	customerEmail := em.emailTemplate.CustomerEmail(cus)
	emailJson, err := json.Marshal(customerEmail)
	if err != nil {
		return errors.Wrap(err, "error when marshal email data")
	}
	emailFilePath := path.Join(em.outputEmailsPath, fmt.Sprintf("%s-%d.json", cus.Email, time.Now().Unix()))
	if err = ioutil.WriteFile(emailFilePath, emailJson, 0644); err != nil {
		return errors.Wrapf(err, "error when write email to file %s", emailFilePath)
	}
	return nil
}

func (em EmailService) BulkSendMails(customers []customer.Customer) {
	wg := sync.WaitGroup{}
	for _, cus := range customers {
		wg.Add(1)
		go func(cus customer.Customer) {
			defer wg.Done()
			if err := em.SendMail(cus); err != nil {
				fmt.Printf("error when send mail to customer %s: %s\n", cus.Email, err.Error())
			}
		}(cus)
	}
	wg.Wait()
}

package customer

import (
	"email-console/pkg/array"
	"email-console/pkg/csv"
	"github.com/pkg/errors"
	"os"
)

type Customer struct {
	Title     string
	FirstName string
	LastName  string
	Email     string
}

func (c Customer) EmailValid() bool {
	return c.Email != ""
}

func (c Customer) ToStrings() []string {
	return []string{c.Title, c.FirstName, c.LastName, c.Email}
}

func LoadFromCsv(path string) ([]Customer, error) {
	records, err := csv.ReadFile(path)
	if err != nil {
		return nil, err
	}
	customers := make([]Customer, 0)
	for ind, record := range records {
		if ind == 0 {
			continue
		}
		customers = append(customers, Customer{
			Title:     array.SafeIndexAt(record, 0),
			FirstName: array.SafeIndexAt(record, 1),
			LastName:  array.SafeIndexAt(record, 2),
			Email:     array.SafeIndexAt(record, 3),
		})
	}
	return customers, nil
}

func WriteCustomersToCsv(path string, customers []Customer) error {
	lines := make([][]string, 0)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		lines = append(lines, []string{"TITLE", "FIRST_NAME", "LAST_NAME", "EMAIL"})
	} else if err != nil {
		return errors.Wrap(err, "cannot check output emails path")
	}
	for _, cus := range customers {
		lines = append(lines, cus.ToStrings())
	}
	return csv.WriteToFile(path, lines)
}

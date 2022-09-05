package customer

import (
	"email-console/pkg/array"
	"os"
	"testing"
)

var (
	testCsvData   = "TITLE,FIRST_NAME,LAST_NAME,EMAIL\nMr,John,Smith,john.smith@example.com\nMrs,Michelle,Smith,michelle.smith@example.com\nMr,David,Smith,\n"
	testCustomers = []Customer{
		{"Mr", "John", "Smith", "john.smith@example.com"},
		{"Mrs", "Michelle", "Smith", "michelle.smith@example.com"},
		{"Mr", "David", "Smith", ""},
	}
	testCsvPath = "test.csv"
)

func TestCustomer_EmailValid(t *testing.T) {
	tests := []struct {
		name     string
		customer Customer
		expected bool
	}{
		{
			name: "customer email valid",
			customer: Customer{
				Email: "test@example.com",
			},
			expected: true,
		},
		{
			name: "customer email invalid",
			customer: Customer{
				Email: "",
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.customer.EmailValid()
			if actual != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}

func TestCustomer_ToStrings(t *testing.T) {
	title, firstName, lastName, email := "Mr", "John", "Smith", "john.smith@example.com"
	tests := []struct {
		name     string
		customer Customer
		expected []string
	}{
		{
			name: "output as expected",
			customer: Customer{
				Title:     title,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			},
			expected: []string{title, firstName, lastName, email},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.customer.ToStrings()
			if !array.StringArrayEqual(actual, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}

func TestLoadFromCsv(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Customer
	}{
		{
			name:     "load success",
			input:    testCsvData,
			expected: testCustomers,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.WriteFile(testCsvPath, []byte(tt.input), 0660)
			defer os.Remove(testCsvPath)
			if err != nil {
				t.Error("error when write data to csv file")
				return
			}
			actual, err := LoadFromCsv(testCsvPath)
			if err != nil {
				t.Errorf("expected success but got error %s", err.Error())
				return
			}
			equal := len(actual) == len(tt.expected)
			for ind, cus := range actual {
				expectedCus := tt.expected[ind]
				if cus.Title != expectedCus.Title ||
					cus.FirstName != expectedCus.FirstName ||
					cus.LastName != expectedCus.LastName ||
					cus.Email != expectedCus.Email {
					equal = false
				}
			}
			if !equal {
				t.Errorf("expected customers to be %v but got %v", tt.expected, actual)
			}
		})
	}
}

func TestWriteCustomersToCsv(t *testing.T) {
	tests := []struct {
		name     string
		input    []Customer
		expected string
	}{
		{
			name:     "load success",
			input:    testCustomers,
			expected: testCsvData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteCustomersToCsv(testCsvPath, tt.input)
			defer os.Remove(testCsvPath)
			if err != nil {
				t.Errorf("expected success but got error %s", err.Error())
			}
			csvData, err := os.ReadFile(testCsvPath)
			if err != nil {
				t.Errorf("error when read csv file: %s", err.Error())
				return
			}
			actual := string(csvData)
			if actual != tt.expected {
				t.Errorf("expected data in csv file to be %s but got %s", tt.expected, actual)
			}
		})
	}
}

package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func ReadFile(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read file %s", path)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("error when close file: %s", err.Error())
		}
	}()

	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse file as CSV for %s", path)
	}
	return lines, nil
}

func WriteToFile(path string, lines [][]string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return errors.Wrap(err, "error when open file")
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("error when close file: %s", err.Error())
		}
	}()
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		return errors.Wrap(err, "error when write lines to file")
	}
	return nil
}

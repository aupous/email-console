package config

import (
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	EmailTemplatePath  string
	CustomersPath      string
	OutputEmailsPath   string
	ErrorCustomersPath string
}

const errorEmailTemplatePath = "error when check email template path"
const errorCustomersPath = "error when check customers path"
const errorOutputEmailsPath = "error when check output emails path"
const outputEmailsPathIsNotDir = "output emails path is not a directory"

func (cfg Config) Verify() error {
	if _, err := os.Stat(cfg.EmailTemplatePath); err != nil {
		return errors.Wrap(err, errorEmailTemplatePath)
	}
	// Check if customer path is exists
	if _, err := os.Stat(cfg.CustomersPath); err != nil {
		return errors.Wrap(err, errorCustomersPath)
	}
	// Check if output emails path is exists
	if stat, err := os.Stat(cfg.OutputEmailsPath); err != nil {
		return errors.Wrap(err, errorOutputEmailsPath)
	} else if !stat.IsDir() {
		return errors.New(outputEmailsPathIsNotDir)
	}
	return nil
}

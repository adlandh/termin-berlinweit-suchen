package config

import (
	"os"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/model"
	"github.com/joho/godotenv"
)

type DotEnvloader struct {
	model.Config
}

func (e *DotEnvloader) LoadConfig() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	e.Config.MainUrl = os.Getenv("MAIN_IRL")
	if os.Getenv("VERBOSE") == "true" {
		e.Config.Verbose = true
	}

	return nil
}

func NewDotEnvLoader() *DotEnvloader {
	return &DotEnvloader{}
}

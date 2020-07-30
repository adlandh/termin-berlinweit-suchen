package config

import (
	"os"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/model"
	_ "github.com/joho/godotenv/autoload"
)

type Envloader struct {
	model.Config
}

func (e *Envloader) LoadConfig() error {
	e.Config.MainUrl = os.Getenv("TERMIN_SUCHEN_MAIN_IRL")
	if os.Getenv("TERMIN_SUCHEN_VERBOSE") == "true" {
		e.Config.Verbose = true
	}

	return nil
}

func NewEnvLoader() *Envloader {
	return &Envloader{}
}

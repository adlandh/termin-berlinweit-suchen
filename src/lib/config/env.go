package config

import (
	"os"
	"strconv"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/model"
	_ "github.com/joho/godotenv/autoload"
)

type Envloader struct {
	model.Config
}

func (e *Envloader) LoadConfig() {
	e.Config.MainUrl = os.Getenv("TERMIN_SUCHEN_MAIN_IRL")
	if os.Getenv("TERMIN_SUCHEN_VERBOSE") == "true" {
		e.Config.Verbose = true
	}

	e.CheckPeriod, _ = strconv.Atoi(os.Getenv("TERMIN_SUCHEN_CHECK_PERIOD"))

	if e.CheckPeriod == 0 {
		e.CheckPeriod = 30
	}
}

func NewEnvLoader() *Envloader {
	return &Envloader{}
}

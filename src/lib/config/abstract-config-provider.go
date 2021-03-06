package config

import "github.com/adlandh/termin-berlinweit-suchen/src/lib/model"

type AbstractConfigProvider interface {
	LoadConfig()
	GetConfig() model.Config
	GetMainURL() string
	GetVerbose() bool
	GetCheckPeriod() int
}

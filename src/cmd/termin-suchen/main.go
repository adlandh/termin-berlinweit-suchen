package main

import (
	"log"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/app"
	"github.com/adlandh/termin-berlinweit-suchen/src/lib/crawler"
)
import "github.com/adlandh/termin-berlinweit-suchen/src/lib/config"

func main() {
	configProvider := config.NewEnvLoader()
	if err := configProvider.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	crawlerProvider := crawler.NewCollyCrawler(configProvider.GetVerbose())

	mainApp := app.NewApp(configProvider, crawlerProvider)

	if err := mainApp.Run(); err != nil {
		log.Fatal(err)
	}

}

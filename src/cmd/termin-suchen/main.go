package main

import (
	"log"
	"os"
	"time"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/app"
	"github.com/adlandh/termin-berlinweit-suchen/src/lib/config"
	"github.com/adlandh/termin-berlinweit-suchen/src/lib/crawler"
)

func main() {
	configProvider := config.NewEnvLoader()
	configProvider.LoadConfig()

	ticker := time.NewTicker(time.Duration(configProvider.GetCheckPeriod()) * time.Second)
	errCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	for {
		select {
		case <-ticker.C:
			crawlerProvider := crawler.NewCollyCrawler(configProvider.GetVerbose(), errCh)

			mainApp := app.NewApp(configProvider, crawlerProvider, errCh, doneCh)
			mainApp.Run()
		case <-doneCh:
			os.Exit(0)
		case err := <-errCh:
			log.Fatal(err)
		}
	}
}

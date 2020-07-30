package main

import "log"
import "github.com/adlandh/termin-berlinweit-suchen/src/lib/config"

func main() {
	configProvider := config.NewDotEnvLoader()
	if err := configProvider.LoadConfig(); err != nil {
		log.Fatal(err)
	}

}

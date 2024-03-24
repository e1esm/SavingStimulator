package main

import (
	"github.com/e1esm/SavingStimulator/src/cmd/config"
	"github.com/e1esm/SavingStimulator/src/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()
	if cfg == nil {
		log.Fatal("config initialization error")
	}

	h := handlers.New()
	if err := handlers.Start(*h, cfg.WebPort); err != nil {
		log.WithField("err", err).Fatal("handler initialization error")
	}
}

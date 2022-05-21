package main

import (
	"log"

	"github.com/Jeanhwea/baliqiao/config"
	"github.com/Jeanhwea/baliqiao/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}

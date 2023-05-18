package main

import (
	"log"
	"os"
	"stash.transparent.com/gol/notifyCardActivity/models/config"
	"stash.transparent.com/gol/notifyCardActivity/router"
	"stash.transparent.com/gol/notifyCardActivity/services/worker"
)

func main() {
	// Read from config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Panic(err)
	}
	os.Setenv("db_address", cfg.DBAddress)

	eventIdChan, err := worker.StartNewProcess()

	if err != nil {
		log.Panic(err)
	}

	router.Init(cfg.Port, eventIdChan)
}

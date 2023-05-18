package main

import (
	"log"
	"stash.transparent.com/gol/notifyCardActivity/router"
	"stash.transparent.com/gol/notifyCardActivity/services/notify"
)

func main() {
	// TODO Read from config
	port := 3000

	eventIdChan, err := notify.StartNewProcess()

	if err != nil {
		log.Panic(err)
	}

	router.Init(port, eventIdChan)
}

package notify

import (
	"log"
	"stash.transparent.com/gol/notifyCardActivity/models/card"
	"stash.transparent.com/gol/notifyCardActivity/services/mysql"
	"time"
)

func StartNewProcess() (*chan int64, error) {
	db, err := mysql.NewDB()
	if err != nil {
		return nil, err
	}

	// Create a channel to receive user data
	idChan := make(chan int64)

	// Start the worker to process user data
	go worker(idChan)

	ids, err := db.QueryIds("SELECT id FROM cardActivityEvent WHERE status = ?", "pending")
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		idChan <- id
	}

	return &idChan, nil
}

// worker processes user data and sends notifications
func worker(idsChan <-chan int64) {

	for id := range idsChan {
		// Simulate processing time
		time.Sleep(1 * time.Second)

		event := card.ActivityEvent{}

		err := event.Load(id)
		if err != nil {
			log.Printf("Failed to update user notification status %s: %v", event, err)
		}

		err = event.Notify(id)
		if err != nil {
			log.Printf("Failed to update user notification status %s: %v", event, err)
		}

	}
}

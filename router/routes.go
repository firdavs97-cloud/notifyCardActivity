package router

import (
	"context"
	"log"
	"net/http"
	"stash.transparent.com/gol/notifyCardActivity/handlers"
	"strconv"
)

func Init(port int, idsChan *chan int64) {
	ctx := context.WithValue(context.Background(), "idsChan", idsChan)
	handlers.Init(ctx)

	// Start the server and listen on port
	log.Println("Server started on http://localhost:"+strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

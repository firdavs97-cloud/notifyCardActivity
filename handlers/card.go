package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	card2 "stash.transparent.com/gol/notifyCardActivity/models/card"
	"stash.transparent.com/gol/notifyCardActivity/utils"
)

const (
	Events = "/events"
)

func(h Handler) initCardApi(path string) {
	// Register the handler function with the default HTTP server
	http.HandleFunc(path+Events+"/card", h.CardEvent)
}


func(h Handler) CardEvent(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	cardEvent := card2.ActivityEventInput{}
	err := json.NewDecoder(r.Body).Decode(&cardEvent)
	if err != nil {
		utils.SendErrorMessage(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
		return
	}
	cardActivityEvent := card2.NewActivityEvent(h.ctx, cardEvent)
	id, err := cardActivityEvent.Save()
	if err != nil {
		utils.SendErrorMessage(w, http.StatusInternalServerError, fmt.Sprintf("Failed to save card activity event data: %v", err))
		return
	}
	if d, ok := h.ctx.Value("idsChan").(*chan int64); ok {
		*d <- id
	}
	w.WriteHeader(http.StatusOK)
}

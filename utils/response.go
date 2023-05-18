package utils

import (
	"encoding/json"
	"net/http"
)

func SendErrorMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	response := map[string]interface{}{
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
